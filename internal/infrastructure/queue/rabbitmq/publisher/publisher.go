package publisher

import (
	"context"
	"fmt"

	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/queue"
)

type Publisher struct{}

func New() Publisher {
	return Publisher{}
}

func (p Publisher) Publish(ctx context.Context, message queue.Message) error {
	fmt.Printf("publishing message: %s", message)

	return nil
}
