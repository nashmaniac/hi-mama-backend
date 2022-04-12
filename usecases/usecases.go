package usecases

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/config"
)

type usecases struct {
	PeristenceStore adapters.PeristenceStore
	Configuration   *config.Config
}

func NewUsecases(
	ctx context.Context,
	p adapters.PeristenceStore,
	configuration *config.Config,
) (adapters.Usecases, error) {
	return &usecases{
		Configuration:   configuration,
		PeristenceStore: p,
	}, nil
}
