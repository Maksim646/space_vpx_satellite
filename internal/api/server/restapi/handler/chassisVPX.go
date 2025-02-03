package handler

import (
	"context"
	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateChassisVPXHandler(req api.CreateChassisVPXParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create chassis request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	chassis := model.Chassis{
		Axes:                            *req.CreateChassisVPXBody.Axes,
		Height:                          *req.CreateChassisVPXBody.Height,
		Length:                          *req.CreateChassisVPXBody.Length,
		MaxNonOperatingTemperature:      *req.CreateChassisVPXBody.MaxNonOperatingTemperature,
		MaxOperatingTemperature:         *req.CreateChassisVPXBody.MaxOperatingTemperature,
		MaxVibrationRandom:              *req.CreateChassisVPXBody.MaxVibrationRandom,
		MaxVibrationSine:                *req.CreateChassisVPXBody.MaxVibrationSine,
		MinNonOperatingTemperature:      *req.CreateChassisVPXBody.MinNonOperatingTemperature,
		MinOperatingTemperature:         *req.CreateChassisVPXBody.MinOperatingTemperature,
		MinVibrationRandom:              *req.CreateChassisVPXBody.MinVibrationRandom,
		MinVibrationSine:                *req.CreateChassisVPXBody.MinVibrationSine,
		Name:                            *req.CreateChassisVPXBody.Name,
		Overload:                        *req.CreateChassisVPXBody.Overload,
		PeakFrequencySpectrum1:          *req.CreateChassisVPXBody.PeakFrequencySpectrum1,
		PeakFrequencySpectrum2:          *req.CreateChassisVPXBody.PeakFrequencySpectrum2,
		PeakOverloadSpectrum1:           *req.CreateChassisVPXBody.PeakOverloadSpectrum1,
		PeakOverloadSpectrum2:           *req.CreateChassisVPXBody.PeakOverloadSpectrum2,
		PowerHandlingCapabilityPerBoard: *req.CreateChassisVPXBody.PowerHandlingCapabilityPerBoard,
		ShockResponseSpectrum:           *req.CreateChassisVPXBody.ShockResponseSpectrum,
		Size:                            *req.CreateChassisVPXBody.Size,
		Slots:                           *req.CreateChassisVPXBody.Slots,
		TemperaturePerBoard:             *req.CreateChassisVPXBody.TemperaturePerBoard,
		Weight:                          *req.CreateChassisVPXBody.Weight,
		Width:                           *req.CreateChassisVPXBody.Width,
	}

	newChassisID, err := h.chassisUsecase.CreateChassisByID(ctx, chassis)
	if err != nil {
		zap.L().Error("failed to create chassis", zap.Error(err))
		return api.NewCreateChassisInternalServerError()
	}

	return api.NewCreateChassisOK().WithPayload(&definition.CreateChassisResponse{
		ID: useful.Int64Ptr(newChassisID),
	})
}

func (h *Handler) UpdateChassisVPXHandler(req api.UpdateChassisVPXParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update chassis request, id:" + string(req.ID))
	ctx := req.HTTPRequest.Context()

	// Проверка роли пользователя
	if principal.Role == 0 {
		return api.NewCreateChassisVPXForbidden()
	}

	chassis, err := h.chassisUsecase.GetChassisByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateChassisVPXBadRequest().WithPayload(&definition.Error{
			Message: &model.ChassisNotFound,
		})
	}

	if *req.CreateChassisVPXBody.Axes != 0 {
		chassis.Axes = *req.CreateChassisVPXBody.Axes
	}
	if *req.CreateChassisVPXBody.Height != 0 {
		chassis.Height = *req.CreateChassisVPXBody.Height
	}
	if *req.CreateChassisVPXBody.Length != 0 {
		chassis.Length = *req.CreateChassisVPXBody.Length
	}
	if *req.CreateChassisVPXBody.MaxNonOperatingTemperature != 0 {
		chassis.MaxNonOperatingTemperature = *req.CreateChassisVPXBody.MaxNonOperatingTemperature
	}
	if *req.CreateChassisVPXBody.MaxOperatingTemperature != 0 {
		chassis.MaxOperatingTemperature = *req.CreateChassisVPXBody.MaxOperatingTemperature
	}
	if *req.CreateChassisVPXBody.MaxVibrationRandom != 0 {
		chassis.MaxVibrationRandom = *req.CreateChassisVPXBody.MaxVibrationRandom
	}
	if *req.CreateChassisVPXBody.MaxVibrationSine != 0 {
		chassis.MaxVibrationSine = *req.CreateChassisVPXBody.MaxVibrationSine
	}
	if *req.CreateChassisVPXBody.MinNonOperatingTemperature != 0 {
		chassis.MinNonOperatingTemperature = *req.CreateChassisVPXBody.MinNonOperatingTemperature
	}
	if *req.CreateChassisVPXBody.MinOperatingTemperature != 0 {
		chassis.MinOperatingTemperature = *req.CreateChassisVPXBody.MinOperatingTemperature
	}
	if *req.CreateChassisVPXBody.MinVibrationRandom != 0 {
		chassis.MinVibrationRandom = *req.CreateChassisVPXBody.MinVibrationRandom
	}
	if *req.CreateChassisVPXBody.MinVibrationSine != 0 {
		chassis.MinVibrationSine = *req.CreateChassisVPXBody.MinVibrationSine
	}
	if *req.CreateChassisVPXBody.Name != "" {
		chassis.Name = *req.CreateChassisVPXBody.Name
	}
	if *req.CreateChassisVPXBody.Overload != 0 {
		chassis.Overload = *req.CreateChassisVPXBody.Overload
	}
	if *req.CreateChassisVPXBody.PeakFrequencySpectrum1 != 0 {
		chassis.PeakFrequencySpectrum1 = *req.CreateChassisVPXBody.PeakFrequencySpectrum1
	}
	if *req.CreateChassisVPXBody.PeakFrequencySpectrum2 != 0 {
		chassis.PeakFrequencySpectrum2 = *req.CreateChassisVPXBody.PeakFrequencySpectrum2
	}
	if *req.CreateChassisVPXBody.PeakOverloadSpectrum1 != 0 {
		chassis.PeakOverloadSpectrum1 = *req.CreateChassisVPXBody.PeakOverloadSpectrum1
	}
	if *req.CreateChassisVPXBody.PeakOverloadSpectrum2 != 0 {
		chassis.PeakOverloadSpectrum2 = *req.CreateChassisVPXBody.PeakOverloadSpectrum2
	}
	if *req.CreateChassisVPXBody.PowerHandlingCapabilityPerBoard != 0 {
		chassis.PowerHandlingCapabilityPerBoard = *req.CreateChassisVPXBody.PowerHandlingCapabilityPerBoard
	}
	if *req.CreateChassisVPXBody.ShockResponseSpectrum != 0 {
		chassis.ShockResponseSpectrum = *req.CreateChassisVPXBody.ShockResponseSpectrum
	}
	if *req.CreateChassisVPXBody.Size != "" {
		chassis.Size = *req.CreateChassisVPXBody.Size
	}
	if *req.CreateChassisVPXBody.Slots != 0 {
		chassis.Slots = *req.CreateChassisVPXBody.Slots
	}
	if *req.CreateChassisVPXBody.TemperaturePerBoard != 0 {
		chassis.TemperaturePerBoard = *req.CreateChassisVPXBody.TemperaturePerBoard
	}
	if *req.CreateChassisVPXBody.Weight != 0 {
		chassis.Weight = *req.CreateChassisVPXBody.Weight
	}
	if *req.CreateChassisVPXBody.Width != 0 {
		chassis.Width = *req.CreateChassisVPXBody.Width
	}

	err = h.chassisUsecase.UpdateChassisByID(ctx, chassis)
	if err != nil {
		zap.L().Error("failed to update chassis", zap.Error(err))
		return api.NewUpdateChassisInternalServerError()
	}

	return api.NewUpdateChassisVPXOK().WithPayload(&definition.ChassisVPX{
		ID:                              chassis.ID,
		Axes:                            &chassis.Axes,
		Height:                          &chassis.Height,
		Length:                          &chassis.Length,
		MaxNonOperatingTemperature:      &chassis.MaxNonOperatingTemperature,
		MaxOperatingTemperature:         &chassis.MaxOperatingTemperature,
		MaxVibrationRandom:              &chassis.MaxVibrationRandom,
		MaxVibrationSine:                &chassis.MaxVibrationSine,
		MinNonOperatingTemperature:      &chassis.MinNonOperatingTemperature,
		MinOperatingTemperature:         &chassis.MinOperatingTemperature,
		MinVibrationRandom:              &chassis.MinVibrationRandom,
		MinVibrationSine:                &chassis.MinVibrationSine,
		Name:                            &chassis.Name,
		Overload:                        &chassis.Overload,
		PeakFrequencySpectrum1:          &chassis.PeakFrequencySpectrum1,
		PeakFrequencySpectrum2:          &chassis.PeakFrequencySpectrum2,
		PeakOverloadSpectrum1:           &chassis.PeakOverloadSpectrum1,
		PeakOverloadSpectrum2:           &chassis.PeakOverloadSpectrum2,
		PowerHandlingCapabilityPerBoard: &chassis.PowerHandlingCapabilityPerBoard,
		ShockResponseSpectrum:           &chassis.ShockResponseSpectrum,
		Size:                            &chassis.Size,
		Slots:                           &chassis.Slots,
		TemperaturePerBoard:             &chassis.TemperaturePerBoard,
		Weight:                          &chassis.Weight,
		Width:                           &chassis.Width,
		CreatedAt:                       chassis.CreatedAt.Unix(),
		UpdatedAt:                       time.Now().Unix(),
	})
}

func (h *Handler) DeleteChassisVPXHandler(req api.DeleteChassisVPXParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update chassis request, id:" + string(req.ID))
	ctx := req.HTTPRequest.Context()

	// Проверка роли пользователя
	if principal.Role == 0 {
		return api.NewDeleteChassisVPXForbidden()
	}

	chassis, err := h.chassisUsecase.GetChassisByID(ctx, req.ID)
	if err != nil {
		return api.NewDeleteChassisVPXNotFound().WithPayload(&definition.Error{
			Message: &model.ChassisNotFound,
		})
	}

	err = h.chassisUsecase.DeleteChassisByID(ctx, chassis.ID)
	if err != nil {
		zap.L().Error("failed to delete chassis", zap.Error(err))
		return api.NewDeleteChassisInternalServerError()
	}

	return api.NewDeleteChassisOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("Chassis deleted successfully"),
	})
}

func (h *Handler) GetChassisVPXHandler(req api.GetChassisVPXByIDParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update chassis request, id:" + string(req.ID))
	ctx := req.HTTPRequest.Context()

	// Проверка роли пользователя
	if principal.Role == 0 {
		return api.NewDeleteChassisVPXForbidden()
	}

	chassis, err := h.chassisUsecase.GetChassisByID(ctx, req.ID)
	if err != nil {
		return api.NewDeleteChassisVPXNotFound().WithPayload(&definition.Error{
			Message: &model.ChassisNotFound,
		})
	}

	return api.NewGetChassisVPXByIDOK().WithPayload(&definition.ChassisVPX{
		ID:                              chassis.ID,
		Axes:                            &chassis.Axes,
		Height:                          &chassis.Height,
		Length:                          &chassis.Length,
		MaxNonOperatingTemperature:      &chassis.MaxNonOperatingTemperature,
		MaxOperatingTemperature:         &chassis.MaxOperatingTemperature,
		MaxVibrationRandom:              &chassis.MaxVibrationRandom,
		MaxVibrationSine:                &chassis.MaxVibrationSine,
		MinNonOperatingTemperature:      &chassis.MinNonOperatingTemperature,
		MinOperatingTemperature:         &chassis.MinOperatingTemperature,
		MinVibrationRandom:              &chassis.MinVibrationRandom,
		MinVibrationSine:                &chassis.MinVibrationSine,
		Name:                            &chassis.Name,
		Overload:                        &chassis.Overload,
		PeakFrequencySpectrum1:          &chassis.PeakFrequencySpectrum1,
		PeakFrequencySpectrum2:          &chassis.PeakFrequencySpectrum2,
		PeakOverloadSpectrum1:           &chassis.PeakOverloadSpectrum1,
		PeakOverloadSpectrum2:           &chassis.PeakOverloadSpectrum2,
		PowerHandlingCapabilityPerBoard: &chassis.PowerHandlingCapabilityPerBoard,
		ShockResponseSpectrum:           &chassis.ShockResponseSpectrum,
		Size:                            &chassis.Size,
		Slots:                           &chassis.Slots,
		TemperaturePerBoard:             &chassis.TemperaturePerBoard,
		Weight:                          &chassis.Weight,
		Width:                           &chassis.Width,
		CreatedAt:                       chassis.CreatedAt.Unix(),
		UpdatedAt:                       time.Now().Unix(),
	})
}

func (h *Handler) GetAvailableChassisVPX(req api.GetAvailableChassisVPXParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get available chassis request, id:" + string(principal.ID))
	ctx := req.HTTPRequest.Context()

	filters := make(map[string]interface{})

	if req.FilterChassisVPXByMaxLengthFrom != nil {
		filters[model.FilterChassisByMaxLength] = *req.FilterChassisVPXByMaxLengthFrom
	}
	if req.FilterChassisVPXByMinLengthTo != nil {
		filters[model.FilterChassisByMinLength] = *req.FilterChassisVPXByMinLengthTo
	}
	if req.FilterChassisVPXByMaxWidthFrom != nil {
		filters[model.FilterChassisByMaxWidth] = *req.FilterChassisVPXByMaxWidthFrom
	}
	if req.FilterChassisVPXByMinWidthTo != nil {
		filters[model.FilterChassisByMinWidth] = *req.FilterChassisVPXByMinWidthTo
	}
	if req.FilterChassisVPXByMaxHeightFrom != nil {
		filters[model.FilterChassisByMaxHeight] = *req.FilterChassisVPXByMaxHeightFrom
	}
	if req.FilterChassisVPXByMinHeightTo != nil {
		filters[model.FilterChassisByMinHeight] = *req.FilterChassisVPXByMinHeightTo
	}
	if req.FilterChassisVPXByMaxWeightFrom != nil {
		filters[model.FilterChassisByMaxWeight] = *req.FilterChassisVPXByMaxWeightFrom
	}
	if req.FilterChassisVPXByMinWeightTo != nil {
		filters[model.FilterChassisByMinWeight] = *req.FilterChassisVPXByMinWeightTo
	}
	if req.FilterChassisVPXByMaxPowerHandlingCapabilityPerBoardFrom != nil {
		filters[model.FilterChassisByMaxPowerHandlingCapabilityPerBoard] = *req.FilterChassisVPXByMaxPowerHandlingCapabilityPerBoardFrom
	}
	if req.FilterChassisVPXByMinPowerHandlingCapabilityPerBoardTo != nil {
		filters[model.FilterChassisByMinPowerHandlingCapabilityPerBoard] = *req.FilterChassisVPXByMinPowerHandlingCapabilityPerBoardTo
	}
	if req.FilterChassisVPXByMaxTemperaturePerBoardFrom != nil {
		filters[model.FilterChassisByMaxTemperaturePerBoard] = *req.FilterChassisVPXByMaxTemperaturePerBoardFrom
	}
	if req.FilterChassisVPXByMinTemperaturePerBoardTo != nil {
		filters[model.FilterChassisByMinTemperaturePerBoard] = *req.FilterChassisVPXByMinTemperaturePerBoardTo
	}

	if req.SortField == nil {
		*req.SortField = model.DefaultChassisSort
	}

	chassises, err := h.chassisUsecase.GetChassisByFilters(ctx, req.Offset, req.Limit, *req.SortField, filters)
	if err != nil {
		zap.L().Error("error fetch chassis", zap.Error(err))
		return api.NewGetAvailableChassisBadRequest()
	}

	chassisesResult := h.ChassisesVPXToDefinition(ctx, chassises)

	return api.NewGetAvailableChassisVPXOK().WithPayload(&definition.ChassisesVPX{
		Count:     useful.Int64Ptr(int64(len(chassisesResult))),
		Chassises: chassisesResult,
	})
}

func (h *Handler) ChassisesVPXToDefinition(ctx context.Context, chassises []model.Chassis) []*definition.ChassisVPX {
	chassisesData := make([]*definition.ChassisVPX, len(chassises))

	for i, _ := range chassises {
		chassisesData[i] = &definition.ChassisVPX{
			ID:                              chassises[i].ID,
			Axes:                            &chassises[i].Axes,
			Height:                          &chassises[i].Height,
			Length:                          &chassises[i].Length,
			MaxNonOperatingTemperature:      &chassises[i].MaxNonOperatingTemperature,
			MaxOperatingTemperature:         &chassises[i].MaxOperatingTemperature,
			MaxVibrationRandom:              &chassises[i].MaxVibrationRandom,
			MaxVibrationSine:                &chassises[i].MaxVibrationSine,
			MinNonOperatingTemperature:      &chassises[i].MinNonOperatingTemperature,
			MinOperatingTemperature:         &chassises[i].MinOperatingTemperature,
			MinVibrationRandom:              &chassises[i].MinVibrationRandom,
			MinVibrationSine:                &chassises[i].MinVibrationSine,
			Name:                            &chassises[i].Name,
			Overload:                        &chassises[i].Overload,
			PeakFrequencySpectrum1:          &chassises[i].PeakFrequencySpectrum1,
			PeakFrequencySpectrum2:          &chassises[i].PeakFrequencySpectrum2,
			PeakOverloadSpectrum1:           &chassises[i].PeakOverloadSpectrum1,
			PeakOverloadSpectrum2:           &chassises[i].PeakOverloadSpectrum2,
			PowerHandlingCapabilityPerBoard: &chassises[i].PowerHandlingCapabilityPerBoard,
			ShockResponseSpectrum:           &chassises[i].ShockResponseSpectrum,
			Size:                            &chassises[i].Size,
			Slots:                           &chassises[i].Slots,
			TemperaturePerBoard:             &chassises[i].TemperaturePerBoard,
			Weight:                          &chassises[i].Weight,
			Width:                           &chassises[i].Width,
			CreatedAt:                       chassises[i].CreatedAt.Unix(),
			UpdatedAt:                       chassises[i].UpdatedAt.Time.Unix(),
		}
	}
	return chassisesData
}
