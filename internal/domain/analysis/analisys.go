package analysis

import (
	"github.com/diegodesousas/go-boilerplate-example/domain/order"
	"github.com/diegodesousas/go-boilerplate-example/domain/provider"
	"github.com/google/uuid"
)

type Type string

const (
	Sync  Type = "sync"
	Async Type = "async"
)

func New() Analysis {
	return Analysis{
		ID: uuid.New().String(),
	}
}

type Analysis struct {
	ID       string        `json:"id"`
	Order    order.Order   `json:"order"`
	Provider provider.Data `json:"provider"`
}

type Result struct {
	Analysis Analysis        `json:"analysis"`
	Result   provider.Result `json:"result"`
}
