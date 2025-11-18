package event

import "context"

type EventProducerInterface interface {
	Publish(ctx context.Context, topic string, key string, value []byte) error
	PublishBatch(ctx context.Context, message []Message) error
}

type EventConsumerInterface interface {
	Subcribe(topics []string) error
	Consume(ctx context.Context, handler MessageHandler) error
	Close() error
}

type Message struct {
	Topic string
	Key   string
	Value []byte
}

type ConsumeMessage struct {
	Topic     string
	Partition int32
	Offset    int64
	Key       []byte
	Value     []byte
	Timestamp int64
}

type MessageHandler interface {
	HandleMessage(ctx context.Context, topic string, value []byte) error
}
