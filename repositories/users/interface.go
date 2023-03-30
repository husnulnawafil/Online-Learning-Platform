package repositories

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
)

type UserRepository interface {
	GetUserDetailByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserDetailByUUID(ctx context.Context, UUID string) (*models.User, error)
	RegisterUser(ctx context.Context, data models.User) error
	DeleteUser(ctx context.Context, UUID string) (int64, error)
}
