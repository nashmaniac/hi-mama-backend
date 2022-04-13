package usecases

import (
	"context"
	"fmt"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (u *usecases) GetTimeList(
	ctx context.Context,
	username string,
) ([]models.Entry, error) {

	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database")
	}

	entries, err := u.PeristenceStore.FindEntries(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error in fetching entries")
	}

	return entries, nil
}
