package postgres

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (r *postgresRepository) FindEntryById(ctx context.Context, id uint) (*models.Entry, error) {

	var entry models.Entry
	tx := r.db.Where("id = ?", id).Preload("User").First(&entry)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &entry, nil
}
