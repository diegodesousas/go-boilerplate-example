package queue

import "context"

type Message struct {
	Content    []byte
	RoutingKey string
	Exchange   string
}

type Publisher interface {
	Publish(ctx context.Context, message Message) error
}
