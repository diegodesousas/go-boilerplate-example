package konduto

import (
	"context"

	"github.com/diegodesousas/go-boilerplate-example/domain/order"
	"github.com/diegodesousas/go-boilerplate-example/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/logger"
)

var TYPE provider.Type = "konduto"

type Provider struct{}

func New() Provider {
	return Provider{}
}

func (p Provider) Type() provider.Type {
	return TYPE
}

func (p Provider) Analyse(ctx context.Context, order order.Order) (provider.Result, error) {
	logger.FromContext(ctx).Info("analysing on trex provider")

	result := provider.Result{
		Status: "success",
	}

	return result, nil
}