package event

import (
	"context"
	"log"
	"strings"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaConsumerInfra struct {
	Client *kgo.Client
}

func NewKafkaConsumerInfra(topics []string) (*KafkaConsumerInfra, error) {
	brokers := strings.Split(config.GlobalEnv.KafkaHost, ",")

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(config.GlobalEnv.KafkaConsumerGroup),
		kgo.ConsumeTopics(topics...),
		kgo.DisableAutoCommit(),
		kgo.BlockRebalanceOnPoll(),
		kgo.OnPartitionsAssigned(func(ctx context.Context, cl *kgo.Client, assignments map[string][]int32) {
			log.Printf("kafka: partitions assigned: %v", assignments)
		}),
		kgo.OnPartitionsRevoked(func(ctx context.Context, cl *kgo.Client, revoked map[string][]int32) {
			log.Printf("kafka: partitions revoked: %v", revoked)
		}),
	)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumerInfra{
		Client: client,
	}, nil
}
