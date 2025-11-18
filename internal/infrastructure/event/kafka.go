package event

import (
	"context"
	"log"
	"strings"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInfra struct {
	Client  *kgo.Client
	handler MessageHandler
}

func NewKafkaInfra() (*KafkaInfra, error) {
	brokers := strings.Split(config.GlobalEnv.KafkaHost, ",")

	client, err := kgo.NewClient(
		kgo.AllowAutoTopicCreation(),
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(config.GlobalEnv.KafkaConsumerGroup),
		kgo.ConsumeTopics(
			constant.DefaultTopic,
		),
	)
	if err != nil {
		return nil, err
	}

	return &KafkaInfra{
		Client: client,
	}, nil

}

func (k *KafkaInfra) Shutdown(ctx context.Context) error {
	if k.Client != nil {
		k.Client.Close()
	}
	return nil
}

func (k *KafkaInfra) Subcribe(topics []string) error {
	if k.Client == nil {
		return nil
	}

	k.Client.AddConsumeTopics(topics...)
	return nil
}

func (k *KafkaInfra) Publish(ctx context.Context, topic string, key string, value []byte) error {
	if k.Client == nil {
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
		return err
	}

	return nil
}

func (k *KafkaInfra) RegisterConsumer(ctx context.Context, handler MessageHandler) error {
	if k.Client == nil {
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fetches := k.Client.PollFetches(ctx)

			if errs := fetches.Errors(); len(errs) > 0 {
				for _, err := range errs {
					log.Default().Println(err)
				}
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				err := k.handler.HandleMessage(ctx, record.Topic, record.Value)
				if err != nil {
					log.Default().Println(err)
				}

			}
		}
	}
}
