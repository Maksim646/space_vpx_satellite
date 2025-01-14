package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type Usecase struct {
	adminRepository model.IAdminRepository
}

func New(adminRepository model.IAdminRepository) model.IAdminUsecase {
	return &Usecase{
		adminRepository: adminRepository,
	}
}

func (u *Usecase) CreateAdmin(ctx context.Context, email string, name string, password string) (string, error) {
	return u.adminRepository.CreateAdmin(ctx, email, name, password)
}

func (u *Usecase) GetAdminByEmail(ctx context.Context, email string) (model.Admin, error) {
	return u.adminRepository.GetAdminByEmail(ctx, email)
}

func (u *Usecase) GetAdminByID(ctx context.Context, userID string) (model.Admin, error) {
	return u.adminRepository.GetAdminByID(ctx, userID)
}
