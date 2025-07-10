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

func (h *Handler) CreatePowerSystemHandler(req api.CreatePowerSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create power system request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreatePowerSystemForbidden()
	}

	powerSystem := model.CubeSatPowerSystem{
		Name:                              sql.NullString{String: req.CreatePowerSystemBody.Name, Valid: true},
		Length:                            sql.NullFloat64{Float64: *req.CreatePowerSystemBody.Length, Valid: true},
		Width:                             sql.NullFloat64{Float64: *req.CreatePowerSystemBody.Width, Valid: true},
		Height:                            sql.NullFloat64{Float64: *req.CreatePowerSystemBody.Height, Valid: true},
		Weight:                            sql.NullFloat64{Float64: req.CreatePowerSystemBody.Weight, Valid: true},
		SolarPanelChannels:                sql.NullInt64{Int64: req.CreatePowerSystemBody.SolarPanelChannels, Valid: true},
		SolarPanelsType:                   sql.NullString{String: req.CreatePowerSystemBody.SolarPanelsType, Valid: true},
		SolarPanelVoltageMin:              sql.NullFloat64{Float64: req.CreatePowerSystemBody.SolarPanelVoltageMin, Valid: true},
		SolarPanelVoltageMax:              sql.NullFloat64{Float64: req.CreatePowerSystemBody.SolarPanelVoltageMax, Valid: true},
		SolarPanelCurrentPerChannelMax:    sql.NullFloat64{Float64: req.CreatePowerSystemBody.SolarPanelCurrentPerChannelMax, Valid: true},
		TotalCurrentOfSolarPanelsMax:      sql.NullFloat64{Float64: req.CreatePowerSystemBody.TotalCurrentOfSolarPanelsMax, Valid: true},
		OutputChannels:                    sql.NullInt64{Int64: req.CreatePowerSystemBody.OutputChannels, Valid: true},
		SystemBusVoltageSolarPanels:       sql.NullFloat64{Float64: req.CreatePowerSystemBody.SystemBusVoltageSolarPanels, Valid: true},
		MaxSystemBusVoltageOutputChannels: sql.NullFloat64{Float64: req.CreatePowerSystemBody.MaxSystemBusVoltageOutputChannels, Valid: true},
		MinSystemBusVoltageOutputChannels: sql.NullFloat64{Float64: req.CreatePowerSystemBody.MinSystemBusVoltageOutputChannels, Valid: true},
		CurrentOutputChannelsMax:          sql.NullFloat64{Float64: req.CreatePowerSystemBody.CurrentOutputChannelsMax, Valid: true},
		TotalOutputCurrent:                sql.NullFloat64{Float64: req.CreatePowerSystemBody.TotalOutputCurrent, Valid: true},
		DataInterface:                     sql.NullString{String: req.CreatePowerSystemBody.DataInterface, Valid: true},
		MaxOperatingTemperature:           sql.NullFloat64{Float64: req.CreatePowerSystemBody.MaxOperatingTemperature, Valid: true},
		MinOperatingTemperature:           sql.NullFloat64{Float64: req.CreatePowerSystemBody.MinOperatingTemperature, Valid: true},
		MechanicalVibration:               sql.NullInt64{Int64: req.CreatePowerSystemBody.MechanicalVibration, Valid: true},
		MechanicalShock:                   sql.NullInt64{Int64: req.CreatePowerSystemBody.MechanicalShock, Valid: true},
	}

	PowerSystemID, err := h.cubeSatPowerSystemUsecase.CreatePowerSystem(ctx, powerSystem)
	if err != nil {
		zap.L().Error("error create new power system", zap.Error(err))
		return api.NewCreatePowerSystemInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new power system"),
		})
	}

	powerSystemResult, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByID(ctx, PowerSystemID)
	if err != nil {
		return api.NewCreatePowerSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.PowerSystemNotFound,
		})
	}
	return api.NewCreatePowerSystemOK().WithPayload(&definition.CubeSatPowerSystem{
		Name:                              powerSystemResult.Name.String,
		Length:                            powerSystemResult.Length.Float64,
		Width:                             powerSystemResult.Width.Float64,
		Height:                            powerSystemResult.Height.Float64,
		Weight:                            powerSystemResult.Weight.Float64,
		SolarPanelChannels:                powerSystemResult.SolarPanelChannels.Int64,
		SolarPanelsType:                   powerSystemResult.SolarPanelsType.String,
		SolarPanelVoltageMin:              powerSystemResult.SolarPanelVoltageMin.Float64,
		SolarPanelVoltageMax:              powerSystemResult.SolarPanelVoltageMax.Float64,
		SolarPanelCurrentPerChannelMax:    powerSystemResult.SolarPanelCurrentPerChannelMax.Float64,
		TotalCurrentOfSolarPanelsMax:      powerSystemResult.TotalCurrentOfSolarPanelsMax.Float64,
		OutputChannels:                    powerSystemResult.OutputChannels.Int64,
		SystemBusVoltageSolarPanels:       powerSystemResult.SystemBusVoltageSolarPanels.Float64,
		MaxSystemBusVoltageOutputChannels: powerSystemResult.MaxSystemBusVoltageOutputChannels.Float64,
		MinSystemBusVoltageOutputChannels: powerSystemResult.MinSystemBusVoltageOutputChannels.Float64,
		CurrentOutputChannelsMax:          powerSystemResult.CurrentOutputChannelsMax.Float64,
		TotalOutputCurrent:                powerSystemResult.TotalOutputCurrent.Float64,
		DataInterface:                     powerSystemResult.DataInterface.String,
		MaxOperatingTemperature:           powerSystemResult.MaxOperatingTemperature.Float64,
		MinOperatingTemperature:           powerSystemResult.MinOperatingTemperature.Float64,
		MechanicalVibration:               powerSystemResult.MechanicalVibration.Int64,
		MechanicalShock:                   powerSystemResult.MechanicalShock.Int64,
		CreatedAt:                         powerSystemResult.CreatedAt.Unix(),
		UpdatedAt:                         powerSystemResult.UpdatedAt.Time.Unix(),
		ID:                                &powerSystemResult.ID,
	})
}

func (h *Handler) GetPowerSystemHandler(req api.GetPowerSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get power system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetPowerSystemForbidden()
	}

	powerSystem, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch power system", zap.Error(err))
		return api.NewGetPowerSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.PowerSystemNotFound,
		})
	}

	return api.NewGetPowerSystemOK().WithPayload(&definition.CubeSatPowerSystem{
		Name:                              powerSystem.Name.String,
		Length:                            powerSystem.Length.Float64,
		Width:                             powerSystem.Width.Float64,
		Height:                            powerSystem.Height.Float64,
		Weight:                            powerSystem.Weight.Float64,
		SolarPanelChannels:                powerSystem.SolarPanelChannels.Int64,
		SolarPanelsType:                   powerSystem.SolarPanelsType.String,
		SolarPanelVoltageMin:              powerSystem.SolarPanelVoltageMin.Float64,
		SolarPanelVoltageMax:              powerSystem.SolarPanelVoltageMax.Float64,
		SolarPanelCurrentPerChannelMax:    powerSystem.SolarPanelCurrentPerChannelMax.Float64,
		TotalCurrentOfSolarPanelsMax:      powerSystem.TotalCurrentOfSolarPanelsMax.Float64,
		OutputChannels:                    powerSystem.OutputChannels.Int64,
		SystemBusVoltageSolarPanels:       powerSystem.SystemBusVoltageSolarPanels.Float64,
		MaxSystemBusVoltageOutputChannels: powerSystem.MaxSystemBusVoltageOutputChannels.Float64,
		MinSystemBusVoltageOutputChannels: powerSystem.MinSystemBusVoltageOutputChannels.Float64,
		CurrentOutputChannelsMax:          powerSystem.CurrentOutputChannelsMax.Float64,
		TotalOutputCurrent:                powerSystem.TotalOutputCurrent.Float64,
		DataInterface:                     powerSystem.DataInterface.String,
		MaxOperatingTemperature:           powerSystem.MaxOperatingTemperature.Float64,
		MinOperatingTemperature:           powerSystem.MinOperatingTemperature.Float64,
		MechanicalVibration:               powerSystem.MechanicalVibration.Int64,
		MechanicalShock:                   powerSystem.MechanicalShock.Int64,
		CreatedAt:                         powerSystem.CreatedAt.Unix(),
		UpdatedAt:                         powerSystem.UpdatedAt.Time.Unix(),
		ID:                                &powerSystem.ID,
	})
}

func (h *Handler) UpdateCubeSatPowerSystemHandler(req api.UpdateCubeSatPowerSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update power system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateCubeSatPowerSystemForbidden()
	}

	powerSystem, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch power system", zap.Error(err))
		return api.NewUpdateCubeSatPowerSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.PowerSystemNotFound,
		})
	}

	if req.UpdatePowerSystemBody.Name != "" {
		powerSystem.Name.String = req.UpdatePowerSystemBody.Name
		powerSystem.Name.Valid = true
	}

	if req.UpdatePowerSystemBody.Length != nil {
		powerSystem.Length.Float64 = *req.UpdatePowerSystemBody.Length
		powerSystem.Length.Valid = true
	}

	if req.UpdatePowerSystemBody.Width != nil {
		powerSystem.Width.Float64 = *req.UpdatePowerSystemBody.Width
		powerSystem.Width.Valid = true
	}

	if req.UpdatePowerSystemBody.Height != nil {
		powerSystem.Height.Float64 = *req.UpdatePowerSystemBody.Height
		powerSystem.Height.Valid = true
	}

	if req.UpdatePowerSystemBody.Weight != 0 {
		powerSystem.Weight.Float64 = req.UpdatePowerSystemBody.Weight
		powerSystem.Weight.Valid = true
	}

	if req.UpdatePowerSystemBody.DataInterface != "" {
		powerSystem.DataInterface.String = req.UpdatePowerSystemBody.DataInterface
	}

	if req.UpdatePowerSystemBody.SolarPanelChannels != 0 {
		powerSystem.SolarPanelChannels.Int64 = req.UpdatePowerSystemBody.SolarPanelChannels
	}

	if req.UpdatePowerSystemBody.SolarPanelsType != "" {
		powerSystem.SolarPanelsType.String = req.UpdatePowerSystemBody.SolarPanelsType
	}

	if req.UpdatePowerSystemBody.SolarPanelVoltageMin != 0 {
		powerSystem.SolarPanelVoltageMin.Float64 = req.UpdatePowerSystemBody.SolarPanelVoltageMin
	}

	if req.UpdatePowerSystemBody.SolarPanelVoltageMax != 0 {
		powerSystem.SolarPanelVoltageMax.Float64 = req.UpdatePowerSystemBody.SolarPanelVoltageMax
	}

	if req.UpdatePowerSystemBody.SolarPanelCurrentPerChannelMax != 0 {
		powerSystem.SolarPanelCurrentPerChannelMax.Float64 = req.UpdatePowerSystemBody.SolarPanelCurrentPerChannelMax
	}

	if req.UpdatePowerSystemBody.TotalCurrentOfSolarPanelsMax != 0 {
		powerSystem.TotalCurrentOfSolarPanelsMax.Float64 = req.UpdatePowerSystemBody.TotalCurrentOfSolarPanelsMax
	}

	if req.UpdatePowerSystemBody.OutputChannels != 0 {
		powerSystem.OutputChannels.Int64 = req.UpdatePowerSystemBody.OutputChannels
	}

	if req.UpdatePowerSystemBody.SystemBusVoltageSolarPanels != 0 {
		powerSystem.SystemBusVoltageSolarPanels.Float64 = req.UpdatePowerSystemBody.SystemBusVoltageSolarPanels
	}

	if req.UpdatePowerSystemBody.MaxSystemBusVoltageOutputChannels != 0 {
		powerSystem.MaxSystemBusVoltageOutputChannels.Float64 = req.UpdatePowerSystemBody.MaxSystemBusVoltageOutputChannels
	}

	if req.UpdatePowerSystemBody.MinSystemBusVoltageOutputChannels != 0 {
		powerSystem.MinSystemBusVoltageOutputChannels.Float64 = req.UpdatePowerSystemBody.MinSystemBusVoltageOutputChannels
	}

	if req.UpdatePowerSystemBody.CurrentOutputChannelsMax != 0 {
		powerSystem.CurrentOutputChannelsMax.Float64 = req.UpdatePowerSystemBody.CurrentOutputChannelsMax
	}

	if req.UpdatePowerSystemBody.TotalOutputCurrent != 0 {
		powerSystem.TotalOutputCurrent.Float64 = req.UpdatePowerSystemBody.TotalOutputCurrent
	}

	if req.UpdatePowerSystemBody.MaxOperatingTemperature != 0 {
		powerSystem.MaxOperatingTemperature.Float64 = req.UpdatePowerSystemBody.MaxOperatingTemperature
	}

	if req.UpdatePowerSystemBody.MinOperatingTemperature != 0 {
		powerSystem.MinOperatingTemperature.Float64 = req.UpdatePowerSystemBody.MinOperatingTemperature
	}

	if req.UpdatePowerSystemBody.MechanicalVibration != 0 {
		powerSystem.MechanicalVibration.Int64 = req.UpdatePowerSystemBody.MechanicalVibration
	}

	if req.UpdatePowerSystemBody.MechanicalShock != 0 {
		powerSystem.MechanicalShock.Int64 = req.UpdatePowerSystemBody.MechanicalShock
	}

	powerSystem.UpdatedAt.Time = time.Now()

	err = h.cubeSatPowerSystemUsecase.UpdatePowerSystem(ctx, powerSystem)
	if err != nil {
		zap.L().Error("error update power system", zap.Error(err))
		return api.NewUpdateCubeSatPowerSystemInternalServerError()
	}

	newPowerSystem, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByID(ctx, powerSystem.ID)
	if err != nil {
		zap.L().Error("error fetch updated power system", zap.Error(err))
		return api.NewUpdateCubeSatPowerSystemInternalServerError()
	}

	return api.NewUpdateCubeSatPowerSystemOK().WithPayload(&definition.CubeSatPowerSystem{
		Name:                              newPowerSystem.Name.String,
		Length:                            newPowerSystem.Length.Float64,
		Width:                             newPowerSystem.Width.Float64,
		Height:                            newPowerSystem.Height.Float64,
		Weight:                            newPowerSystem.Weight.Float64,
		SolarPanelChannels:                newPowerSystem.SolarPanelChannels.Int64,
		SolarPanelsType:                   newPowerSystem.SolarPanelsType.String,
		SolarPanelVoltageMin:              newPowerSystem.SolarPanelVoltageMin.Float64,
		SolarPanelVoltageMax:              newPowerSystem.SolarPanelVoltageMax.Float64,
		SolarPanelCurrentPerChannelMax:    newPowerSystem.SolarPanelCurrentPerChannelMax.Float64,
		TotalCurrentOfSolarPanelsMax:      newPowerSystem.TotalCurrentOfSolarPanelsMax.Float64,
		OutputChannels:                    newPowerSystem.OutputChannels.Int64,
		SystemBusVoltageSolarPanels:       newPowerSystem.SystemBusVoltageSolarPanels.Float64,
		MaxSystemBusVoltageOutputChannels: newPowerSystem.MaxSystemBusVoltageOutputChannels.Float64,
		MinSystemBusVoltageOutputChannels: newPowerSystem.MinSystemBusVoltageOutputChannels.Float64,
		CurrentOutputChannelsMax:          newPowerSystem.CurrentOutputChannelsMax.Float64,
		TotalOutputCurrent:                newPowerSystem.TotalOutputCurrent.Float64,
		DataInterface:                     newPowerSystem.DataInterface.String,
		MaxOperatingTemperature:           newPowerSystem.MaxOperatingTemperature.Float64,
		MinOperatingTemperature:           newPowerSystem.MinOperatingTemperature.Float64,
		MechanicalVibration:               newPowerSystem.MechanicalVibration.Int64,
		MechanicalShock:                   newPowerSystem.MechanicalShock.Int64,
		CreatedAt:                         newPowerSystem.CreatedAt.Unix(),
		UpdatedAt:                         newPowerSystem.UpdatedAt.Time.Unix(),
		ID:                                &newPowerSystem.ID,
	})
}

func (h *Handler) DeleteCubeSatPowerSystemHandler(req api.DeleteCubeSatPowerSystemParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete power system request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteCubeSatPowerSystemForbidden()
	}

	powerSystem, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch power system", zap.Error(err))
		return api.NewDeleteCubeSatPowerSystemBadRequest().WithPayload(&definition.Error{
			Message: &model.PowerSystemNotFound,
		})
	}

	err = h.cubeSatPowerSystemUsecase.DeletePowerSystem(ctx, powerSystem.ID)
	if err != nil {
		zap.L().Error("error delete power system", zap.Error(err))
		return api.NewDeleteCubeSatPowerSystemInternalServerError()
	}

	return api.NewDeleteCubeSatPowerSystemOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("Cube sat power system deleted successfully"),
	})
}

func (h *Handler) GetCubeSatPowerSystems(req api.GetCubeSatPowerSystemsParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch power systems request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetCubeSatPowerSystemsForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultCubeSatPowerSystemsSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.CubeSatPowerSystemsSortByCreatedAt
		default:
			sortParams = model.DefaultCubeSatPowerSystemsSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterCubeSatPowerSystemByLengthMax != nil {
		filters[model.FilterCubeSatPowerSystemByMaxLength] = *req.FilterCubeSatPowerSystemByLengthMax
	}

	if req.FilterCubeSatPowerSystemByLengthMin != nil {
		filters[model.FilterCubeSatPowerSystemByMinLength] = *req.FilterCubeSatPowerSystemByLengthMin
	}

	powerSystems, err := h.cubeSatPowerSystemUsecase.GetPowerSystemsByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetch power systems", zap.Error(err))
		return api.NewGetCubeSatPowerSystemsBadRequest().WithPayload(&definition.Error{
			Message: &model.PowerSystemsNotFound,
		})
	}

	powerSystemsDefinition := h.PowerSystemsToDefinition(ctx, powerSystems)

	return api.NewGetCubeSatPowerSystemsOK().WithPayload(&definition.CubeSatPowerSystems{
		Count:               useful.Int64Ptr(int64(len(powerSystemsDefinition))),
		CubeSatPowerSystems: powerSystemsDefinition,
	})
}

func (h *Handler) PowerSystemsToDefinition(ctx context.Context, cubeSatPowerSystems []model.CubeSatPowerSystem) []*definition.CubeSatPowerSystem {
	cubeSatPowerSystemsDefinition := make([]*definition.CubeSatPowerSystem, len(cubeSatPowerSystems))

	for i := range cubeSatPowerSystems {
		cubeSatPowerSystemsDefinition[i] = &definition.CubeSatPowerSystem{
			ID:                                &cubeSatPowerSystems[i].ID,
			Name:                              cubeSatPowerSystems[i].Name.String,
			Length:                            cubeSatPowerSystems[i].Length.Float64,
			Width:                             cubeSatPowerSystems[i].Width.Float64,
			Height:                            cubeSatPowerSystems[i].Height.Float64,
			Weight:                            cubeSatPowerSystems[i].Weight.Float64,
			SolarPanelChannels:                cubeSatPowerSystems[i].SolarPanelChannels.Int64,
			SolarPanelsType:                   cubeSatPowerSystems[i].SolarPanelsType.String,
			SolarPanelVoltageMin:              cubeSatPowerSystems[i].SolarPanelVoltageMin.Float64,
			SolarPanelVoltageMax:              cubeSatPowerSystems[i].SolarPanelVoltageMax.Float64,
			SolarPanelCurrentPerChannelMax:    cubeSatPowerSystems[i].SolarPanelCurrentPerChannelMax.Float64,
			TotalCurrentOfSolarPanelsMax:      cubeSatPowerSystems[i].TotalCurrentOfSolarPanelsMax.Float64,
			OutputChannels:                    cubeSatPowerSystems[i].OutputChannels.Int64,
			SystemBusVoltageSolarPanels:       cubeSatPowerSystems[i].SystemBusVoltageSolarPanels.Float64,
			MaxSystemBusVoltageOutputChannels: cubeSatPowerSystems[i].MaxSystemBusVoltageOutputChannels.Float64,
			MinSystemBusVoltageOutputChannels: cubeSatPowerSystems[i].MinSystemBusVoltageOutputChannels.Float64,
			CurrentOutputChannelsMax:          cubeSatPowerSystems[i].CurrentOutputChannelsMax.Float64,
			TotalOutputCurrent:                cubeSatPowerSystems[i].TotalOutputCurrent.Float64,
			DataInterface:                     cubeSatPowerSystems[i].DataInterface.String,
			MaxOperatingTemperature:           cubeSatPowerSystems[i].MaxOperatingTemperature.Float64,
			MinOperatingTemperature:           cubeSatPowerSystems[i].MinOperatingTemperature.Float64,
			MechanicalVibration:               cubeSatPowerSystems[i].MechanicalVibration.Int64,
			MechanicalShock:                   cubeSatPowerSystems[i].MechanicalShock.Int64,
			CreatedAt:                         cubeSatPowerSystems[i].CreatedAt.Unix(),
			UpdatedAt:                         cubeSatPowerSystems[i].UpdatedAt.Time.Unix(),
		}
	}

	return cubeSatPowerSystemsDefinition
}
