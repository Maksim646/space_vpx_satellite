package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CubeSatFrameToDefinition(ctx context.Context, cubeSat model.CubeSatFrame) *definition.CubeSatFrame {
	cubeSatFrameDefinition := definition.CubeSatFrame{
		Height:                  &cubeSat.Height.Float64,
		ID:                      &cubeSat.ID,
		Length:                  &cubeSat.Length.Float64,
		Link:                    &cubeSat.Link.String,
		MechanicalShock:         &cubeSat.MechanicalShock.Int64,
		MechanicalVibration:     &cubeSat.MechanicalVibration.Int64,
		OperatingTemperatureMax: &cubeSat.OperatingTemperatureMax.Int64,
		OperatingTemperatureMin: &cubeSat.OperatingTemperatureMin.Int64,
		Weight:                  &cubeSat.Weight.Int64,
		Width:                   &cubeSat.Width.Float64,
	}
	return &cubeSatFrameDefinition
}

func (h *Handler) CreateCubeSatFrameHandler(req api.CreateCubeSatFrameParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	cubeSatFrame := model.CubeSatFrame{
		Length:                  sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Width, Valid: true},
		Weight:                  sql.NullInt64{Int64: *req.CreateCubeSatFrameBody.Weight, Valid: true},
		OperatingTemperatureMin: sql.NullInt64{Int64: *req.CreateCubeSatFrameBody.OperatingTemperatureMin, Valid: true},
		OperatingTemperatureMax: sql.NullInt64{Int64: *req.CreateCubeSatFrameBody.OperatingTemperatureMax, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: *req.CreateCubeSatFrameBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: *req.CreateCubeSatFrameBody.MechanicalShock, Valid: true},
		Link:                    sql.NullString{String: *req.CreateCubeSatFrameBody.Link, Valid: true},
	}

	cubeSatFrameID, err := h.cubeSatFrameUsecase.CreateCubeSatFrame(ctx, cubeSatFrame)
	if err != nil {
		zap.L().Error("error create new cube sat frame", zap.Error(err))
		return api.NewCreateCubeSatFrameInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new cube sat frame"),
		})
	}

	cubeSatFrame, err = h.cubeSatFrameUsecase.GetCubeSatFrameByID(ctx, cubeSatFrameID)
	if err != nil {
		zap.L().Error("error fetch new cube sat frame", zap.Error(err))
		return api.NewCreateCubeSatFrameInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new cube sat frame"),
		})
	}

	cubeSatFrameDefinition := h.CubeSatFrameToDefinition(ctx, cubeSatFrame)

	return api.NewCreateCubeSatFrameOK().WithPayload(cubeSatFrameDefinition)
}

func (h *Handler) GetCubeSatFrame(req api.GetCubeSatFrameParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewCreateChassisForbidden()
	}

	cubeSatFrame, err := h.cubeSatFrameUsecase.GetCubeSatFrameByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch new cube sat frame", zap.Error(err))
		return api.NewGetCubeSatFrameBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFrameNotFound,
		})
	}

	cubeSatFrameDefinition := h.CubeSatFrameToDefinition(ctx, cubeSatFrame)

	return api.NewGetCubeSatFrameOK().WithPayload(cubeSatFrameDefinition)
}

func (h *Handler) UpdateCubeSatFrameHandler(req api.UpdateCubeSatFrameParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	cubeSatFrame, err := h.cubeSatFrameUsecase.GetCubeSatFrameByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch new cube sat frame", zap.Error(err))
		return api.NewUpdateCubeSatFrameBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFrameNotFound,
		})
	}

	if req.CreateCubeSatFrameBody.Height != nil {
		cubeSatFrame.Height.Float64 = *req.CreateCubeSatFrameBody.Height
	}
	if req.CreateCubeSatFrameBody.Length != nil {
		cubeSatFrame.Length.Float64 = *req.CreateCubeSatFrameBody.Length
	}
	if req.CreateCubeSatFrameBody.Width != nil {
		cubeSatFrame.Width.Float64 = *req.CreateCubeSatFrameBody.Width
	}
	if req.CreateCubeSatFrameBody.Weight != nil {
		cubeSatFrame.Weight.Int64 = *req.CreateCubeSatFrameBody.Weight
	}
	if req.CreateCubeSatFrameBody.OperatingTemperatureMin != nil {
		cubeSatFrame.OperatingTemperatureMin.Int64 = *req.CreateCubeSatFrameBody.OperatingTemperatureMin
	}
	if req.CreateCubeSatFrameBody.OperatingTemperatureMax != nil {
		cubeSatFrame.OperatingTemperatureMax.Int64 = *req.CreateCubeSatFrameBody.OperatingTemperatureMax
	}
	if req.CreateCubeSatFrameBody.MechanicalVibration != nil {
		cubeSatFrame.MechanicalVibration.Int64 = *req.CreateCubeSatFrameBody.MechanicalVibration
	}
	if req.CreateCubeSatFrameBody.MechanicalShock != nil {
		cubeSatFrame.MechanicalShock.Int64 = *req.CreateCubeSatFrameBody.MechanicalShock
	}
	if req.CreateCubeSatFrameBody.Link != nil {
		cubeSatFrame.Link.String = *req.CreateCubeSatFrameBody.Link
	}
	cubeSatFrame.UpdatedAt.Time = time.Now()

	// Сохраняем обновленный CubeSatFrame
	err = h.cubeSatFrameUsecase.UpdateCubeSatFrame(ctx, cubeSatFrame)
	if err != nil {
		zap.L().Error("error updating cube sat frame", zap.Error(err))
		return api.NewUpdateCubeSatFrameInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error update cube sat frame"),
		})
	}

	cubeSatFrame, err = h.cubeSatFrameUsecase.GetCubeSatFrameByID(ctx, cubeSatFrame.ID)
	if err != nil {
		zap.L().Error("error fetch cube sat frame", zap.Error(err))
		return api.NewUpdateCubeSatFrameInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error fetch cube sat frame"),
		})
	}

	cubeSatFrameDefinition := h.CubeSatFrameToDefinition(ctx, cubeSatFrame)

	return api.NewUpdateCubeSatFrameOK().WithPayload(cubeSatFrameDefinition)
}
