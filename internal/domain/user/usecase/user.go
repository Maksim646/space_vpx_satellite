package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type Usecase struct {
	userRepository model.IUserRepository
}

func New(userRepository model.IUserRepository) model.IUserUsecase {
	return &Usecase{
		userRepository: userRepository,
	}
}

func (u *Usecase) CreateUser(ctx context.Context, email string, name string, password string) (string, error) {
	return u.userRepository.CreateUser(ctx, email, name, password)
}
