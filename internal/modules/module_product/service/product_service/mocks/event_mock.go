package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockEventProducer is a mock implementation of EventProducerInterface
type MockEventProducer struct {
	mock.Mock
}

func (m *MockEventProducer) Publish(ctx context.Context, topic string, key string, value []byte) error {
	args := m.Called(ctx, topic, key, value)
	return args.Error(0)
}

func (m *MockEventProducer) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
