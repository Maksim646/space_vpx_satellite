package handler

import (
	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateChassisHandler(req api.CreateChassisParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create chassis request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	chassis := model.Chassis{
		Axes:                            *req.CreateChassisBody.Axes,
		Height:                          *req.CreateChassisBody.Height,
		Length:                          *req.CreateChassisBody.Length,
		MaxNonOperatingTemperature:      *req.CreateChassisBody.MaxNonOperatingTemperature,
		MaxOperatingTemperature:         *req.CreateChassisBody.MaxOperatingTemperature,
		MaxVibrationRandom:              *req.CreateChassisBody.MaxVibrationRandom,
		MaxVibrationSine:                *req.CreateChassisBody.MaxVibrationSine,
		MinNonOperatingTemperature:      *req.CreateChassisBody.MinNonOperatingTemperature,
		MinOperatingTemperature:         *req.CreateChassisBody.MinOperatingTemperature,
		MinVibrationRandom:              *req.CreateChassisBody.MinVibrationRandom,
		MinVibrationSine:                *req.CreateChassisBody.MinVibrationSine,
		Name:                            *req.CreateChassisBody.Name,
		Overload:                        *req.CreateChassisBody.Overload,
		PeakFrequencySpectrum1:          *req.CreateChassisBody.PeakFrequencySpectrum1,
		PeakFrequencySpectrum2:          *req.CreateChassisBody.PeakFrequencySpectrum2,
		PeakOverloadSpectrum1:           *req.CreateChassisBody.PeakOverloadSpectrum1,
		PeakOverloadSpectrum2:           *req.CreateChassisBody.PeakOverloadSpectrum2,
		PowerHandlingCapabilityPerBoard: *req.CreateChassisBody.PowerHandlingCapabilityPerBoard,
		ShockResponseSpectrum:           *req.CreateChassisBody.ShockResponseSpectrum,
		Size:                            *req.CreateChassisBody.Size,
		Slots:                           *req.CreateChassisBody.Slots,
		TemperaturePerBoard:             *req.CreateChassisBody.TemperaturePerBoard,
		Weight:                          *req.CreateChassisBody.Weight,
		Width:                           *req.CreateChassisBody.Width,
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

func (h *Handler) UpdateChassisHandler(req api.UpdateChassisParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update chassis request, id:" + string(req.ID))
	ctx := req.HTTPRequest.Context()

	// Проверка роли пользователя
	if principal.Role == 0 {
		return api.NewCreateChassisForbidden()
	}

	chassis, err := h.chassisUsecase.GetChassisByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateChassisBadRequest().WithPayload(&definition.Error{
			Message: &model.ChassisNotFound,
		})
	}

	// Обновляем только те поля, которые не пустые
	if *req.CreateChassisBody.Axes != 0 {
		chassis.Axes = *req.CreateChassisBody.Axes
	}
	if *req.CreateChassisBody.Height != 0 {
		chassis.Height = *req.CreateChassisBody.Height
	}
	if *req.CreateChassisBody.Length != 0 {
		chassis.Length = *req.CreateChassisBody.Length
	}
	if *req.CreateChassisBody.MaxNonOperatingTemperature != 0 {
		chassis.MaxNonOperatingTemperature = *req.CreateChassisBody.MaxNonOperatingTemperature
	}
	if *req.CreateChassisBody.MaxOperatingTemperature != 0 {
		chassis.MaxOperatingTemperature = *req.CreateChassisBody.MaxOperatingTemperature
	}
	if *req.CreateChassisBody.MaxVibrationRandom != 0 {
		chassis.MaxVibrationRandom = *req.CreateChassisBody.MaxVibrationRandom
	}
	if *req.CreateChassisBody.MaxVibrationSine != 0 {
		chassis.MaxVibrationSine = *req.CreateChassisBody.MaxVibrationSine
	}
	if *req.CreateChassisBody.MinNonOperatingTemperature != 0 {
		chassis.MinNonOperatingTemperature = *req.CreateChassisBody.MinNonOperatingTemperature
	}
	if *req.CreateChassisBody.MinOperatingTemperature != 0 {
		chassis.MinOperatingTemperature = *req.CreateChassisBody.MinOperatingTemperature
	}
	if *req.CreateChassisBody.MinVibrationRandom != 0 {
		chassis.MinVibrationRandom = *req.CreateChassisBody.MinVibrationRandom
	}
	if *req.CreateChassisBody.MinVibrationSine != 0 {
		chassis.MinVibrationSine = *req.CreateChassisBody.MinVibrationSine
	}
	if *req.CreateChassisBody.Name != "" {
		chassis.Name = *req.CreateChassisBody.Name
	}
	if *req.CreateChassisBody.Overload != 0 {
		chassis.Overload = *req.CreateChassisBody.Overload
	}
	if *req.CreateChassisBody.PeakFrequencySpectrum1 != 0 {
		chassis.PeakFrequencySpectrum1 = *req.CreateChassisBody.PeakFrequencySpectrum1
	}
	if *req.CreateChassisBody.PeakFrequencySpectrum2 != 0 {
		chassis.PeakFrequencySpectrum2 = *req.CreateChassisBody.PeakFrequencySpectrum2
	}
	if *req.CreateChassisBody.PeakOverloadSpectrum1 != 0 {
		chassis.PeakOverloadSpectrum1 = *req.CreateChassisBody.PeakOverloadSpectrum1
	}
	if *req.CreateChassisBody.PeakOverloadSpectrum2 != 0 {
		chassis.PeakOverloadSpectrum2 = *req.CreateChassisBody.PeakOverloadSpectrum2
	}
	if *req.CreateChassisBody.PowerHandlingCapabilityPerBoard != 0 {
		chassis.PowerHandlingCapabilityPerBoard = *req.CreateChassisBody.PowerHandlingCapabilityPerBoard
	}
	if *req.CreateChassisBody.ShockResponseSpectrum != 0 {
		chassis.ShockResponseSpectrum = *req.CreateChassisBody.ShockResponseSpectrum
	}
	if *req.CreateChassisBody.Size != "" {
		chassis.Size = *req.CreateChassisBody.Size
	}
	if *req.CreateChassisBody.Slots != 0 {
		chassis.Slots = *req.CreateChassisBody.Slots
	}
	if *req.CreateChassisBody.TemperaturePerBoard != 0 {
		chassis.TemperaturePerBoard = *req.CreateChassisBody.TemperaturePerBoard
	}
	if *req.CreateChassisBody.Weight != 0 {
		chassis.Weight = *req.CreateChassisBody.Weight
	}
	if *req.CreateChassisBody.Width != 0 {
		chassis.Width = *req.CreateChassisBody.Width
	}

	err = h.chassisUsecase.UpdateChassisByID(ctx, chassis)
	if err != nil {
		zap.L().Error("failed to update chassis", zap.Error(err))
		return api.NewUpdateChassisInternalServerError()
	}

	return api.NewUpdateChassisOK().WithPayload(&definition.Chassis{
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

func updateChassisInt64Field(chassis *model.Chassis, value *int64, updateFunc func(int64)) {
	if value != nil && *value != 0 {
		updateFunc(*value)
	}
}

func updateChassisField(chassis *model.Chassis, value *float64, updateFunc func(float64)) {
	if value != nil && *value != 0 {
		updateFunc(*value)
	}
}

func updateChassisStringField(chassis *model.Chassis, value *string, updateFunc func(string)) {
	if value != nil && *value != "" {
		updateFunc(*value)
	}
}
