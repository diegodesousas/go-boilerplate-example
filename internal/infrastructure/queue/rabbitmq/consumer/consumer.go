package consumer

import (
	"context"

	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/queue"
)

type Consumer struct {
	Queue   string
	Handler queue.Handler
}

func New(queue string, handler queue.Handler) Consumer {
	return Consumer{Queue: queue, Handler: handler}
}

func (c Consumer) Consume(ctx context.Context, message queue.Message) error {

	return c.Handler(ctx, message)
}

func Listen(consumers ...Consumer) error {
	return nil
}

func Shutdown() error {
	return nil
}
