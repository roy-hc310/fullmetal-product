package event

import (
	"context"
	"strings"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInfra struct {
	Client *kgo.Client
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
