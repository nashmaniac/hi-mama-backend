package usecases

import (
	"context"
	"fmt"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (u *usecases) FindOngoingTime(
	ctx context.Context,
	username string,
) (*models.Entry, error) {
	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database")
	}

	entry, err := u.PeristenceStore.FindOngoingTime(ctx, *user)
	if err != nil {
		return nil, fmt.Errorf("error fetching ongoing time")
	}

	return entry, nil
}
