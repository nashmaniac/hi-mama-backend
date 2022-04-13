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

	FindOngoingTime(ctx context.Context, user models.User) (*models.Entry, error)
	CreateEntry(ctx context.Context, entry *models.Entry) (*models.Entry, error)
	SaveEntry(ctx context.Context, entry *models.Entry) (*models.Entry, error)
	FindEntries(ctx context.Context, user *models.User) ([]models.Entry, error)
	FindEntryById(ctx context.Context, id uint) (*models.Entry, error)
}
