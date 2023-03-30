package services

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
)

type UserService interface {
	RegisterUser(ctx context.Context, data *models.User) error
	GetUserDetailByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserDetailByUUID(ctx context.Context, UUID string) (*models.User, error)
	DeleteUser(ctx context.Context, UUID string) (int64, error)
}
