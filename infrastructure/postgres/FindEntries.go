package postgres

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (r *postgresRepository) FindEntries(
	ctx context.Context,
	user *models.User,
) ([]models.Entry, error) {
	entries := make([]models.Entry, 0)

	if err := r.db.Where("user_id = ?", user.ID).Preload("User").Order("clock_in_time desc").Find(&entries).Error; err != nil {
		return nil, err
	}

	return entries, nil
}
