package adapters

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

//go:generate mockgen -destination=../mocks/persistenstore.mock.go -package=mocks --build_flags=--mod=mod github.com/nashmaniac/hi-mama/hi-mama-backend/adapters PeristenceStore
type PeristenceStore interface {
	CloseDB(ctx context.Context)
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}
