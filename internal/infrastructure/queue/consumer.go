package queue

import "context"

type Consumer interface {
	Consume(ctx context.Context, message Message) error
}

type Handler func(ctx context.Context, message Message) error
