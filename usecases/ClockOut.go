package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (u *usecases) ClockOut(
	ctx context.Context,
	username string,
) (*models.Entry, error) {
	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database")
	}

	entry, err := u.PeristenceStore.FindOngoingTime(ctx, *user)
	if err != nil {
		return nil, fmt.Errorf("error in fetching entry")
	}

	if entry == nil {
		return nil, fmt.Errorf("you are not clocked in")
	}

	currentTime := time.Now()
	entry.ClockOutTime = &currentTime
	entry, err = u.PeristenceStore.SaveEntry(ctx, entry)
	if err != nil {
		return nil, fmt.Errorf("error in creating entry")
	}
	return entry, nil
}
