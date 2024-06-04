package messenger

import (
	"context"
)

type Messenger interface {
	Publish(ctx context.Context, msg []byte, queueName string, isExchange bool) error
	//Consume(ctx context.Context) (<-chan []byte, error)
	Consume(ctx context.Context, queueName string, handlerFunc func(ctx context.Context, payload []byte) error) error
}

type MessengerService struct {
	Messenger Messenger
}

func NewMessengerService(m Messenger) MessengerService {
	return MessengerService{
		Messenger: m,
	}
}
