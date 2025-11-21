package event

import (
	"context"
	"strings"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInfra struct {
	Client  *kgo.Client
	Enable  bool
	handler MessageHandler
}

func NewKafkaInfra() (*KafkaInfra, error) {
	brokers := strings.Split(config.GlobalEnv.KafkaHost, ",")

	opts := []kgo.Opt{
		kgo.AllowAutoTopicCreation(),
		kgo.SeedBrokers(brokers...),
		kgo.RecordRetries(3),
	}

	if g := strings.TrimSpace(config.GlobalEnv.KafkaConsumerGroup); g != "" {
		opts = append(opts,
			kgo.ConsumerGroup(g),
			kgo.ConsumeTopics(constant.KafkaTopics...),
			kgo.DisableAutoCommit(),
			kgo.BlockRebalanceOnPoll(),
			kgo.OnPartitionsAssigned(func(ctx context.Context, cl *kgo.Client, assignments map[string][]int32) {
				logger.Info(ctx).Interface("assignments", assignments).Msg("Kafka partitions assigned")
			}),
			kgo.OnPartitionsRevoked(func(ctx context.Context, cl *kgo.Client, revoked map[string][]int32) {
				logger.Info(ctx).Interface("revoked", revoked).Msg("Kafka partitions revoked")
			}),
		)
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	enable := true
	if err := client.Ping(context.Background()); err != nil {
		enable = false
	}

	return &KafkaInfra{
		Client: client,
		Enable: enable,
	}, nil

}

func (k *KafkaInfra) Shutdown(ctx context.Context) error {
	if !k.Enable || k.Client == nil {
		return nil
	}

	k.Client.Close()
	return nil
}

func (k *KafkaInfra) Subcribe(topics []string) error {
	if !k.Enable || k.Client == nil {
		return nil
	}

	k.Client.AddConsumeTopics(topics...)
	return nil
}

func (k *KafkaInfra) Publish(ctx context.Context, topic string, key string, value []byte) error {
	if !k.Enable || k.Client == nil {
		// log.Default().Println("Kafka: kafka client is nil")
		return nil
	}

	record := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: value,
	}

	results := k.Client.ProduceSync(ctx, record)

	if err := results.FirstErr(); err != nil {
		logger.Error(ctx).Err(err).
			Str("topic", topic).
			Str("key", key).
			Msg("Failed to publish message to Kafka")
		return err
	}

	return nil
}

func (k *KafkaInfra) RegisterConsumer(ctx context.Context, handler MessageHandler) error {
	if !k.Enable || k.Client == nil {
		return nil
	}

	k.handler = handler

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fetches := k.Client.PollFetches(ctx)
			if fetches.IsClientClosed() {
				return nil
			}

			if errs := fetches.Errors(); len(errs) > 0 {
				for _, err := range errs {
					logger.Error(ctx).Err(err.Err).
						Str("topic", err.Topic).
						Int32("partition", err.Partition).
						Msg("Kafka fetch error")
				}
			}

			var totalSuccessRecords []*kgo.Record

			fetches.EachRecord(func(record *kgo.Record) {
				if err := k.handler.HandleMessage(ctx, record.Topic, record.Value); err != nil {
					logger.Error(ctx).Err(err).
						Str("topic", record.Topic).
						Int32("partition", record.Partition).
						Int64("offset", record.Offset).
						Msg("Error handling Kafka message")
					return
				}

				totalSuccessRecords = append(totalSuccessRecords, record)
			})

			if len(totalSuccessRecords) > 0 {
				if err := k.Client.CommitRecords(ctx, totalSuccessRecords...); err != nil {
					logger.Error(ctx).Err(err).
						Int("record_count", len(totalSuccessRecords)).
						Msg("Failed to commit Kafka records")
				}
			}

			k.Client.AllowRebalance()
		}
	}
}
