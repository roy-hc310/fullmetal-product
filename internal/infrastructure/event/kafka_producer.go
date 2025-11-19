package event

import (
	"strings"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProducerInfra struct {
	Client *kgo.Client
}

func NewKafkaProducerInfra() (*KafkaProducerInfra, error) {
	brokers := strings.Split(config.GlobalEnv.KafkaHost, ",")

	opts := []kgo.Opt{
		kgo.AllowAutoTopicCreation(),
		kgo.SeedBrokers(brokers...),
		kgo.RecordRetries(3),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return &KafkaProducerInfra{
		Client: client,
	}, nil
}
