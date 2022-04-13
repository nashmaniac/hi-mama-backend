package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (u *usecases) EditEntry(
	ctx context.Context,
	id uint,
	clockIn time.Time,
	clockOut *time.Time,
) (*models.Entry, error) {
	entry, err := u.PeristenceStore.FindEntryById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in fetching entry")
	}

	if entry == nil {
		return nil, fmt.Errorf("no entry found")
	}

	entry.ClockOutTime = clockOut
	entry.ClockInTime = clockIn
	entry, err = u.PeristenceStore.SaveEntry(ctx, entry)
	if err != nil {
		return nil, fmt.Errorf("error in saving entry")
	}
	return entry, nil
}
