package postgres

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (r *postgresRepository) CreateEntry(
	ctx context.Context,
	entry *models.Entry,
) (*models.Entry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
