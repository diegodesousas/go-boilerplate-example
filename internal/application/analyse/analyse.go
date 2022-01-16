package analyse

import (
	"context"

	"github.com/diegodesousas/go-boilerplate-example/internal/domain/analysis"
	"github.com/diegodesousas/go-boilerplate-example/internal/domain/provider"
)

type Analyse func(ctx context.Context, a analysis.Analysis) (provider.Result, error)

func New(factory provider.Factory) Analyse {
	return func(ctx context.Context, a analysis.Analysis) (provider.Result, error) {
		if err := Validate(ctx, a); err != nil {
			return provider.Result{}, err
		}

		p, err := factory.Get(a.Provider.Type)
		if err != nil {
			return provider.Result{}, err
		}

		result, err := p.Analyse(ctx, a.Order)
		if err != nil {
			return result, err
		}

		return result, nil
	}
}
