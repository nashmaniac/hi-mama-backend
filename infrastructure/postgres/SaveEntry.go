package postgres

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (r *postgresRepository) SaveEntry(
	ctx context.Context,
	entry *models.Entry,
) (*models.Entry, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
