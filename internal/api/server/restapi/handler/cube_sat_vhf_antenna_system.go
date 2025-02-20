package handler

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateVHFAntennaSystemHandler(req api.CreateVHFAntennaSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create VHF antenna system request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateVHFAntennaSystemForbidden()
	}

	vhfAntennaSystem := model.VHFAntennaSystem{
		Name:                    sql.NullString{String: req.CreateVHFAntennaSystemBody.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.Width, Valid: true},
		Height:                  sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.Height, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.Weight, Valid: true},
		Interface:               sql.NullString{String: req.CreateVHFAntennaSystemBody.Interface, Valid: true},
		Frequency:               sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.Frequency, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.MaxOperatingTemperature, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.CreateVHFAntennaSystemBody.MinOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: req.CreateVHFAntennaSystemBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: req.CreateVHFAntennaSystemBody.MechanicalShock, Valid: true},
	}

	vhfAntennaSystemID, err := h.cubeSatVhfAntennaSystemUsecase.CreateVHFAntennaSystem(ctx, vhfAntennaSystem)
	if err != nil {
		zap.L().Error("error creating new VHF antenna system", zap.Error(err))
		return api.NewCreateVHFAntennaSystemInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error creating new VHF antenna system"),
		})
	}

	vhfAntennaSystemResult, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByID(ctx, vhfAntennaSystemID)
	if err != nil {
		return api.NewCreateVHFAntennaSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFAntennaSystemNotFound,
		})
	}
	return api.NewCreateVHFAntennaSystemOK().WithPayload(&definition.VHFAntennaSystem{
		ID:                      vhfAntennaSystemResult.ID,
		Name:                    vhfAntennaSystemResult.Name.String,
		Length:                  vhfAntennaSystemResult.Length.Float64,
		Width:                   vhfAntennaSystemResult.Width.Float64,
		Height:                  vhfAntennaSystemResult.Height.Float64,
		Weight:                  vhfAntennaSystemResult.Weight.Float64,
		Interface:               vhfAntennaSystemResult.Interface.String,
		Frequency:               vhfAntennaSystemResult.Frequency.Float64,
		MaxOperatingTemperature: vhfAntennaSystemResult.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: vhfAntennaSystemResult.MinOperatingTemperature.Float64,
		MechanicalVibration:     vhfAntennaSystemResult.MechanicalVibration.Int64,
		MechanicalShock:         vhfAntennaSystemResult.MechanicalShock.Int64,
		CreatedAt:               vhfAntennaSystemResult.CreatedAt.Unix(),
		UpdatedAt:               vhfAntennaSystemResult.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) GetVHFAntennaSystemHandler(req api.GetVHFAntennaSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get VHF antenna system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetVHFAntennaSystemForbidden()
	}

	vhfAntennaSystem, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching VHF antenna system", zap.Error(err))
		return api.NewGetVHFAntennaSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFAntennaSystemNotFound,
		})
	}

	return api.NewGetVHFAntennaSystemOK().WithPayload(&definition.VHFAntennaSystem{
		ID:                      vhfAntennaSystem.ID,
		Name:                    vhfAntennaSystem.Name.String,
		Length:                  vhfAntennaSystem.Length.Float64,
		Width:                   vhfAntennaSystem.Width.Float64,
		Height:                  vhfAntennaSystem.Height.Float64,
		Weight:                  vhfAntennaSystem.Weight.Float64,
		Interface:               vhfAntennaSystem.Interface.String,
		Frequency:               vhfAntennaSystem.Frequency.Float64,
		MaxOperatingTemperature: vhfAntennaSystem.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: vhfAntennaSystem.MinOperatingTemperature.Float64,
		MechanicalVibration:     vhfAntennaSystem.MechanicalVibration.Int64,
		MechanicalShock:         vhfAntennaSystem.MechanicalShock.Int64,
		CreatedAt:               vhfAntennaSystem.CreatedAt.Unix(),
		UpdatedAt:               vhfAntennaSystem.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) UpdateVHFAntennaSystemHandler(req api.UpdateVHFAntennaSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update VHF antenna system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateVHFAntennaSystemForbidden()
	}

	vhfAntennaSystem, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching VHF antenna system", zap.Error(err))
		return api.NewUpdateVHFAntennaSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFAntennaSystemNotFound,
		})
	}

	if req.UpdateVHFAntennaSystemBody.Name != "" {
		vhfAntennaSystem.Name.String = req.UpdateVHFAntennaSystemBody.Name
		vhfAntennaSystem.Name.Valid = true
	}

	if req.UpdateVHFAntennaSystemBody.Length != 0 {
		vhfAntennaSystem.Length.Float64 = req.UpdateVHFAntennaSystemBody.Length
		vhfAntennaSystem.Length.Valid = true
	}

	if req.UpdateVHFAntennaSystemBody.Width != 0 {
		vhfAntennaSystem.Width.Float64 = req.UpdateVHFAntennaSystemBody.Width
		vhfAntennaSystem.Width.Valid = true
	}

	if req.UpdateVHFAntennaSystemBody.Height != 0 {
		vhfAntennaSystem.Height.Float64 = req.UpdateVHFAntennaSystemBody.Height
		vhfAntennaSystem.Height.Valid = true
	}

	if req.UpdateVHFAntennaSystemBody.Weight != 0 {
		vhfAntennaSystem.Weight.Float64 = req.UpdateVHFAntennaSystemBody.Weight
		vhfAntennaSystem.Weight.Valid = true
	}

	if req.UpdateVHFAntennaSystemBody.Interface != "" {
		vhfAntennaSystem.Interface.String = req.UpdateVHFAntennaSystemBody.Interface
	}

	if req.UpdateVHFAntennaSystemBody.Frequency != 0 {
		vhfAntennaSystem.Frequency.Float64 = req.UpdateVHFAntennaSystemBody.Frequency
	}

	if req.UpdateVHFAntennaSystemBody.MaxOperatingTemperature != 0 {
		vhfAntennaSystem.MaxOperatingTemperature.Float64 = req.UpdateVHFAntennaSystemBody.MaxOperatingTemperature
	}

	if req.UpdateVHFAntennaSystemBody.MinOperatingTemperature != 0 {
		vhfAntennaSystem.MinOperatingTemperature.Float64 = req.UpdateVHFAntennaSystemBody.MinOperatingTemperature
	}

	if req.UpdateVHFAntennaSystemBody.MechanicalVibration != 0 {
		vhfAntennaSystem.MechanicalVibration.Int64 = req.UpdateVHFAntennaSystemBody.MechanicalVibration
	}

	if req.UpdateVHFAntennaSystemBody.MechanicalShock != 0 {
		vhfAntennaSystem.MechanicalShock.Int64 = req.UpdateVHFAntennaSystemBody.MechanicalShock
	}

	vhfAntennaSystem.UpdatedAt.Time = time.Now()

	err = h.cubeSatVhfAntennaSystemUsecase.UpdateVHFAntennaSystem(ctx, vhfAntennaSystem)
	if err != nil {
		zap.L().Error("error updating VHF antenna system", zap.Error(err))
		return api.NewUpdateVHFAntennaSystemInternalServerError()
	}

	newVHFAntennaSystem, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByID(ctx, vhfAntennaSystem.ID)
	if err != nil {
		zap.L().Error("error fetching updated VHF antenna system", zap.Error(err))
		return api.NewUpdateVHFAntennaSystemInternalServerError()
	}

	return api.NewUpdateVHFAntennaSystemOK().WithPayload(&definition.VHFAntennaSystem{
		ID:                      newVHFAntennaSystem.ID,
		Name:                    newVHFAntennaSystem.Name.String,
		Length:                  newVHFAntennaSystem.Length.Float64,
		Width:                   newVHFAntennaSystem.Width.Float64,
		Height:                  newVHFAntennaSystem.Height.Float64,
		Weight:                  newVHFAntennaSystem.Weight.Float64,
		Interface:               newVHFAntennaSystem.Interface.String,
		Frequency:               newVHFAntennaSystem.Frequency.Float64,
		MaxOperatingTemperature: newVHFAntennaSystem.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: newVHFAntennaSystem.MinOperatingTemperature.Float64,
		MechanicalVibration:     newVHFAntennaSystem.MechanicalVibration.Int64,
		MechanicalShock:         newVHFAntennaSystem.MechanicalShock.Int64,
		CreatedAt:               newVHFAntennaSystem.CreatedAt.Unix(),
		UpdatedAt:               newVHFAntennaSystem.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) DeleteVHFAntennaSystemHandler(req api.DeleteVHFAntennaSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete VHF antenna system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteVHFAntennaSystemForbidden()
	}

	vhfAntennaSystem, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching VHF antenna system", zap.Error(err))
		return api.NewDeleteVHFAntennaSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFAntennaSystemNotFound,
		})
	}

	err = h.cubeSatVhfAntennaSystemUsecase.DeleteVHFAntennaSystem(ctx, vhfAntennaSystem.ID)
	if err != nil {
		zap.L().Error("error deleting VHF antenna system", zap.Error(err))
		return api.NewDeleteVHFAntennaSystemInternalServerError()
	}

	return api.NewDeleteVHFAntennaSystemOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("VHF antenna system deleted successfully"),
	})
}

func (h *Handler) GetAvailableVHFAntennaSystemsHandler(req api.GetAvailableVHFAntennaSystemsParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch VHF antenna systems request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetAvailableVHFAntennaSystemsForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultVHFAntennaSystemsSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.VHFAntennaSystemsSortByCreatedAt
		default:
			sortParams = model.DefaultVHFAntennaSystemsSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterVHFAntennaSystemByMinFrequency != nil {
		filters[model.FilterVHFAntennaSystemByMinFrequency] = *req.FilterVHFAntennaSystemByMinFrequency
	}

	if req.FilterVHFAntennaSystemByMaxFrequency != nil {
		filters[model.FilterVHFAntennaSystemByMaxFrequency] = *req.FilterVHFAntennaSystemByMaxFrequency
	}

	vhfAntennaSystems, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemsByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetching VHF antenna systems", zap.Error(err))
		return api.NewGetAvailableVHFAntennaSystemsBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFAntennaSystemsNotFound,
		})
	}

	vhfAntennaSystemsDefinition := h.VHFAntennaSystemsToDefinition(ctx, vhfAntennaSystems)

	return api.NewGetAvailableVHFAntennaSystemsOK().WithPayload(&definition.VHFAntennaSystems{
		Count:            useful.Int64Ptr(int64(len(vhfAntennaSystemsDefinition))),
		VhfAntennaSystem: vhfAntennaSystemsDefinition,
	})
}

func (h *Handler) VHFAntennaSystemsToDefinition(ctx context.Context, vhfAntennaSystems []model.VHFAntennaSystem) []*definition.VHFAntennaSystem {
	vhfAntennaSystemsDefinition := make([]*definition.VHFAntennaSystem, len(vhfAntennaSystems))

	for i := range vhfAntennaSystems {
		vhfAntennaSystemsDefinition[i] = &definition.VHFAntennaSystem{
			ID:                      vhfAntennaSystems[i].ID,
			Name:                    vhfAntennaSystems[i].Name.String,
			Length:                  vhfAntennaSystems[i].Length.Float64,
			Width:                   vhfAntennaSystems[i].Width.Float64,
			Height:                  vhfAntennaSystems[i].Height.Float64,
			Weight:                  vhfAntennaSystems[i].Weight.Float64,
			Interface:               vhfAntennaSystems[i].Interface.String,
			Frequency:               vhfAntennaSystems[i].Frequency.Float64,
			MaxOperatingTemperature: vhfAntennaSystems[i].MaxOperatingTemperature.Float64,
			MinOperatingTemperature: vhfAntennaSystems[i].MinOperatingTemperature.Float64,
			MechanicalVibration:     vhfAntennaSystems[i].MechanicalVibration.Int64,
			MechanicalShock:         vhfAntennaSystems[i].MechanicalShock.Int64,
			CreatedAt:               vhfAntennaSystems[i].CreatedAt.Unix(),
			UpdatedAt:               vhfAntennaSystems[i].UpdatedAt.Time.Unix(),
		}
	}

	return vhfAntennaSystemsDefinition
}
