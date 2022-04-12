package adapters

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

//go:generate mockgen -destination=../mocks/usecases.mock.go -package=mocks --build_flags=--mod=mod github.com/nashmaniac/hi-mama/hi-mama-backend/adapters Usecases
type Usecases interface {
	GetHealthz(ctx context.Context, version string) (map[string]string, error)
	CreateUser(ctx context.Context, username string, password string) (*models.User, error)
	LoginUser(ctx context.Context, username string, password string) (*string, error)
}
