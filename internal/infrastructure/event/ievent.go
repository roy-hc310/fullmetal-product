package event

import "context"

type EventProducerInterface interface {
	Publish(ctx context.Context, topic string, key string, value []byte) error
	// PublishBatch(ctx context.Context, message []Message) error
}

type MessageHandler interface {
	HandleMessage(ctx context.Context, topic string, value []byte) error
}
