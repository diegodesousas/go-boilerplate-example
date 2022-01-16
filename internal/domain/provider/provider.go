package provider

import (
	"context"

	"github.com/diegodesousas/go-boilerplate-example/internal/domain/order"
)

type Type string

type Data struct {
	Type Type `json:"type"`
}

type Result struct {
	Status string `json:"status"`
}

type Provider interface {
	Type() Type
	Analyse(ctx context.Context, order order.Order) (Result, error)
}
