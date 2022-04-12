package postgres

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (r *postgresRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
