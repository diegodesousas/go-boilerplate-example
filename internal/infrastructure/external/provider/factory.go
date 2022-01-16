package provider

import (
	"errors"

	"github.com/diegodesousas/go-boilerplate-example/internal/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/external/provider/clearsale"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/external/provider/konduto"
)

var ErrProviderNotAvailable = errors.New("provider not available")

type factory []provider.Provider

func (f factory) Get(t provider.Type) (provider.Provider, error) {
	for _, p := range f {
		if p.Type() == t {
			return p, nil
		}
	}

	return nil, ErrProviderNotAvailable
}

func NewFactory() provider.Factory {
	return factory{
		clearsale.New(),
		konduto.New(),
	}
}
