package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

// VHFAntennaSystemUsecase implements the IVHFAntennaSystemUsecase interface.
type VHFAntennaSystemUsecase struct {
	vhfAntennaSystemRepository model.IVHFAntennaSystemRepository
}

// New creates a new VHFAntennaSystemUsecase instance.
func New(vhfAntennaSystemRepository model.IVHFAntennaSystemRepository) model.IVHFAntennaSystemUsecase {
	return &VHFAntennaSystemUsecase{
		vhfAntennaSystemRepository: vhfAntennaSystemRepository,
	}
}

// CreateVHFAntennaSystem creates a new VHF antenna system.
func (u *VHFAntennaSystemUsecase) CreateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem model.VHFAntennaSystem) (int64, error) {
	return u.vhfAntennaSystemRepository.CreateVHFAntennaSystem(ctx, vhfAntennaSystem)
}

// GetVHFAntennaSystemByID retrieves a VHF antenna system by its ID.
func (u *VHFAntennaSystemUsecase) GetVHFAntennaSystemByID(ctx context.Context, vhfAntennaSystemID int64) (model.VHFAntennaSystem, error) {
	return u.vhfAntennaSystemRepository.GetVHFAntennaSystemByID(ctx, vhfAntennaSystemID)
}

// GetVHFAntennaSystemsByFilters retrieves VHF antenna systems based on the provided filters.
func (u *VHFAntennaSystemUsecase) GetVHFAntennaSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.VHFAntennaSystem, error) {
	return u.vhfAntennaSystemRepository.GetVHFAntennaSystemsByFilters(ctx, offset, limit, sortParams, filters)
}

// UpdateVHFAntennaSystem updates an existing VHF antenna system.
func (u *VHFAntennaSystemUsecase) UpdateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem model.VHFAntennaSystem) error {
	return u.vhfAntennaSystemRepository.UpdateVHFAntennaSystem(ctx, vhfAntennaSystem)
}

// DeleteVHFAntennaSystem deletes a VHF antenna system by its ID.
func (u *VHFAntennaSystemUsecase) DeleteVHFAntennaSystem(ctx context.Context, vhfAntennaSystemID int64) error {
	return u.vhfAntennaSystemRepository.DeleteVHFAntennaSystem(ctx, vhfAntennaSystemID)
}
