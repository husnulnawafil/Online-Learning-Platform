package services

import (
	"context"
	"errors"
	"time"

	"github.com/husnulnawafil/online-learning-platform/global/constants"
	"github.com/husnulnawafil/online-learning-platform/models"
	userRepositories "github.com/husnulnawafil/online-learning-platform/repositories/users"
)

type UserServiceInstance struct {
	userRepo userRepositories.UserRepository
}

func NewUserService() UserService {
	repoUser := userRepositories.NewUserRepository()
	return &UserServiceInstance{
		userRepo: repoUser,
	}
}

func (us *UserServiceInstance) RegisterUser(ctx context.Context, data *models.User) error {
	data.CreatedAt = time.Now()
	data.UdpatedAt = time.Now()
	data.Role = constants.RoleUser
	return us.userRepo.RegisterUser(ctx, *data)
}

func (us *UserServiceInstance) GetUserDetailByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := us.userRepo.GetUserDetailByEmail(ctx, email)
	if err != nil {
		return user, errors.New("oops failed to retrieve user data by email, please try again later")
	}
	return user, err
}

func (us *UserServiceInstance) GetUserDetailByUUID(ctx context.Context, UUID string) (*models.User, error) {
	user, err := us.userRepo.GetUserDetailByUUID(ctx, UUID)
	if err != nil {
		return user, errors.New("oops failed to retrieve user data by UUID, please try again later")
	}
	return user, err
}

func (us *UserServiceInstance) DeleteUser(ctx context.Context, UUID string) (int64, error) {
	return us.userRepo.DeleteUser(ctx, UUID)
}
