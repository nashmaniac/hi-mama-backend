package adapters

import (
	"context"
	"time"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

//go:generate mockgen -destination=../mocks/usecases.mock.go -package=mocks --build_flags=--mod=mod github.com/nashmaniac/hi-mama/hi-mama-backend/adapters Usecases
type Usecases interface {
	GetHealthz(ctx context.Context, version string) (map[string]string, error)
	CreateUser(ctx context.Context, username string, password string) (*models.User, error)
	LoginUser(ctx context.Context, username string, password string) (*string, error)

	ClockIn(ctx context.Context, username string) (*models.Entry, error)
	ClockOut(ctx context.Context, username string) (*models.Entry, error)
	FindOngoingTime(ctx context.Context, username string) (*models.Entry, error)
	GetTimeList(ctx context.Context, username string) ([]models.Entry, error)
	EditEntry(ctx context.Context, id uint, clockIn time.Time, clockOut *time.Time) (*models.Entry, error)
}
