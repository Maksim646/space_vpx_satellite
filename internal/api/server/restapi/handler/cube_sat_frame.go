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
		Name:                    cubeSat.Name.String,
		Length:                  &cubeSat.Length.Float64,
		Size:                    cubeSat.Size,
		Interface:               cubeSat.Interface.String,
		Link:                    &cubeSat.Link.String,
		MechanicalShock:         &cubeSat.MechanicalShock.Int64,
		MechanicalVibration:     &cubeSat.MechanicalVibration.Int64,
		OperatingTemperatureMax: &cubeSat.OperatingTemperatureMax.Int64,
		OperatingTemperatureMin: &cubeSat.OperatingTemperatureMin.Int64,
		Weight:                  &cubeSat.Weight.Float64,
		Width:                   &cubeSat.Width.Float64,
	}
	return &cubeSatFrameDefinition
}

func (h *Handler) CubeSatFramesToDefinition(ctx context.Context, cubeSats []model.CubeSatFrame) []*definition.CubeSatFrame {
	cubeSatFramesDefinition := make([]*definition.CubeSatFrame, len(cubeSats))

	for i := range cubeSats {
		cubeSatFramesDefinition[i] = &definition.CubeSatFrame{
			ID:                      &cubeSats[i].ID,
			Name:                    cubeSats[i].Name.String,
			Length:                  &cubeSats[i].Length.Float64,
			Width:                   &cubeSats[i].Width.Float64,
			Height:                  &cubeSats[i].Height.Float64,
			Weight:                  &cubeSats[i].Weight.Float64,
			Size:                    cubeSats[i].Size,
			Interface:               cubeSats[i].Interface.String,
			OperatingTemperatureMin: &cubeSats[i].OperatingTemperatureMin.Int64,
			OperatingTemperatureMax: &cubeSats[i].OperatingTemperatureMax.Int64,
			MechanicalVibration:     &cubeSats[i].MechanicalVibration.Int64,
			MechanicalShock:         &cubeSats[i].MechanicalShock.Int64,
			Link:                    &cubeSats[i].Link.String,
			UpdatedAt:               cubeSats[i].UpdatedAt.Time.Unix(),
			CreatedAt:               cubeSats[i].CreatedAt.Unix(),
		}
	}

	return cubeSatFramesDefinition
}

func (h *Handler) CreateCubeSatFrameHandler(req api.CreateCubeSatFrameParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	cubeSatFrame := model.CubeSatFrame{
		Height:                  sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Height, Valid: true},
		Name:                    sql.NullString{String: req.CreateCubeSatFrameBody.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Width, Valid: true},
		Weight:                  sql.NullFloat64{Float64: *req.CreateCubeSatFrameBody.Weight, Valid: true},
		Size:                    req.CreateCubeSatFrameBody.Size,
		Interface:               sql.NullString{String: req.CreateCubeSatFrameBody.Interface, Valid: true},
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

	if req.CreateCubeSatFrameBody.Name != "" {
		cubeSatFrame.Name.String = req.CreateCubeSatFrameBody.Name
	}

	if *req.CreateCubeSatFrameBody.Height != 0 {
		cubeSatFrame.Height.Float64 = *req.CreateCubeSatFrameBody.Height
	}
	if *req.CreateCubeSatFrameBody.Length != 0 {
		cubeSatFrame.Length.Float64 = *req.CreateCubeSatFrameBody.Length
	}
	if *req.CreateCubeSatFrameBody.Width != 0 {
		cubeSatFrame.Width.Float64 = *req.CreateCubeSatFrameBody.Width
	}
	if *req.CreateCubeSatFrameBody.Weight != 0 {
		cubeSatFrame.Weight.Float64 = *req.CreateCubeSatFrameBody.Weight
	}
	if *req.CreateCubeSatFrameBody.OperatingTemperatureMin != 0 {
		cubeSatFrame.OperatingTemperatureMin.Int64 = *req.CreateCubeSatFrameBody.OperatingTemperatureMin
	}
	if *req.CreateCubeSatFrameBody.OperatingTemperatureMax != 0 {
		cubeSatFrame.OperatingTemperatureMax.Int64 = *req.CreateCubeSatFrameBody.OperatingTemperatureMax
	}
	if *req.CreateCubeSatFrameBody.MechanicalVibration != 0 {
		cubeSatFrame.MechanicalVibration.Int64 = *req.CreateCubeSatFrameBody.MechanicalVibration
	}
	if *req.CreateCubeSatFrameBody.MechanicalShock != 0 {
		cubeSatFrame.MechanicalShock.Int64 = *req.CreateCubeSatFrameBody.MechanicalShock
	}
	if *req.CreateCubeSatFrameBody.Link != "" {
		cubeSatFrame.Link.String = *req.CreateCubeSatFrameBody.Link
	}
	cubeSatFrame.UpdatedAt.Time = time.Now()

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

func (h *Handler) DeleteCubeSatFrameHandler(req api.DeleteCubeSatFrameParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	cubeSatFrame, err := h.cubeSatFrameUsecase.GetCubeSatFrameByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch cube sat frame", zap.Error(err))
		return api.NewDeleteCubeSatFrameBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFrameNotFound,
		})
	}

	err = h.cubeSatFrameUsecase.DeleteCubeSatFrame(ctx, cubeSatFrame.ID)
	if err != nil {
		zap.L().Error("error delete cube sat frame", zap.Error(err))
		return api.NewDeleteCubeSatFrameInternalServerError()
	}

	return api.NewDeleteCubeSatFrameOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("Cube sat frame deleted successfully"),
	})
}

func (h *Handler) GetAvailableCubeSatFrames(req api.GetCubeSatFramesParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get available cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewCreateChassisForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultCubeSatFramesSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.CubeSatFramesSortByCreatedAt
		default:
			sortParams = model.DefaultCubeSatFramesSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterCubeSatFrameByMinLengthMax != nil {
		filters[model.FilterCubeSatFrameByMaxLength] = *req.FilterCubeSatFrameByMinLengthMax
	}

	if req.FilterCubeSatFrameByMinLengthMin != nil {
		filters[model.FilterCubeSatFrameByMinLength] = *req.FilterCubeSatFrameByMinLengthMin
	}

	cubeSatFrames, err := h.cubeSatFrameUsecase.GetCubeSatFramesByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error get cube sat frames by filters", zap.Error(err))
		return api.NewGetCubeSatFramesBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFramesNotFound,
		})
	}

	cubeSatFramesDefinition := h.CubeSatFramesToDefinition(ctx, cubeSatFrames)

	return api.NewGetCubeSatFramesOK().WithPayload(&definition.CubeSatFrames{
		Count:         useful.Int64Ptr(int64(len(cubeSatFramesDefinition))),
		CubeSatFrames: cubeSatFramesDefinition,
	})

}
