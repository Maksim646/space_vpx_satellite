package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

// VHFTransceiverUsecase implements the IVHFTransceiverUsecase interface.
type VHFTransceiverUsecase struct {
	vhfTransceiverRepository model.IVHFTransceiverRepository
}

// New creates a new VHFTransceiverUsecase instance.
func New(vhfTransceiverRepository model.IVHFTransceiverRepository) model.IVHFTransceiverUsecase {
	return &VHFTransceiverUsecase{
		vhfTransceiverRepository: vhfTransceiverRepository,
	}
}

// CreateVHFTransceiver creates a new VHF transceiver.
func (u *VHFTransceiverUsecase) CreateVHFTransceiver(ctx context.Context, vhfTransceiver model.VHFTransceiver) (int64, error) {
	return u.vhfTransceiverRepository.CreateVHFTransceiver(ctx, vhfTransceiver)
}

// GetVHFTransceiverByID retrieves a VHF transceiver by its ID.
func (u *VHFTransceiverUsecase) GetVHFTransceiverByID(ctx context.Context, vhfTransceiverID int64) (model.VHFTransceiver, error) {
	return u.vhfTransceiverRepository.GetVHFTransceiverByID(ctx, vhfTransceiverID)
}

// GetVHFTransceiversByFilters retrieves VHF transceivers based on the provided filters.
func (u *VHFTransceiverUsecase) GetVHFTransceiversByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.VHFTransceiver, error) {
	return u.vhfTransceiverRepository.GetVHFTransceiversByFilters(ctx, offset, limit, sortParams, filters)
}

// UpdateVHFTransceiver updates an existing VHF transceiver.
func (u *VHFTransceiverUsecase) UpdateVHFTransceiver(ctx context.Context, vhfTransceiver model.VHFTransceiver) error {
	return u.vhfTransceiverRepository.UpdateVHFTransceiver(ctx, vhfTransceiver)
}

// DeleteVHFTransceiver deletes a VHF transceiver by its ID.
func (u *VHFTransceiverUsecase) DeleteVHFTransceiver(ctx context.Context, vhfTransceiverID int64) error {
	return u.vhfTransceiverRepository.DeleteVHFTransceiver(ctx, vhfTransceiverID)
}
