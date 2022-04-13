package postgres

import (
	"context"
	"errors"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
	"gorm.io/gorm"
)

func (r *postgresRepository) FindOngoingTime(
	ctx context.Context,
	user models.User,
) (*models.Entry, error) {

	var entry models.Entry
	tx := r.db.Where("user_id = ? and clock_out_time is null", user.ID).Preload("User").Order("clock_in_time desc").First(&entry)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entry, nil
}
