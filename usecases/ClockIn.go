package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (u *usecases) ClockIn(
	ctx context.Context,
	username string,
) (*models.Entry, error) {
	user, err := u.PeristenceStore.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user from database")
	}

	entry, err := u.PeristenceStore.FindOngoingTime(ctx, *user)
	if err == nil && entry != nil {
		return nil, fmt.Errorf("you have already clocked in")
	}

	currentTime := time.Now()
	e := &models.Entry{
		User:        *user,
		UserID:      user.ID,
		ClockInTime: currentTime,
	}
	e, err = u.PeristenceStore.CreateEntry(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("error in creating entry")
	}
	return e, nil
}
