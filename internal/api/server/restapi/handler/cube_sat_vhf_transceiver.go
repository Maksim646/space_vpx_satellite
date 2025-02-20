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

func (h *Handler) CreateCubeSatVHFTransceiverHandler(req api.CreateCubeSatVHTransceiverParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create CubeSat VHF transceiver request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateCubeSatVHTransceiverForbidden()
	}

	cubeSatVHFTransceiver := model.VHFTransceiver{
		Name:                    sql.NullString{String: req.CreateVHFTransceiverBody.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.Width, Valid: true},
		Height:                  sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.Height, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.Weight, Valid: true},
		SupplyVoltage:           sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.SupplyVoltage, Valid: true},
		PowerConsumption:        sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.PowerConsumption, Valid: true},
		Interface:               sql.NullString{String: req.CreateVHFTransceiverBody.Interface, Valid: true},
		OperatingFrequency:      sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.OperatingFrequency, Valid: true},
		OutputPower:             sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.OutputPower, Valid: true},
		SensitivityReceiver:     sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.SensitivityReceiver, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.MaxOperatingTemperature, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.CreateVHFTransceiverBody.MinOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: req.CreateVHFTransceiverBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: req.CreateVHFTransceiverBody.MechanicalShock, Valid: true},
	}

	transceiverID, err := h.cubeSatVhfTransceiverUsecase.CreateVHFTransceiver(ctx, cubeSatVHFTransceiver)
	if err != nil {
		zap.L().Error("error creating new CubeSat VHF transceiver", zap.Error(err))
		return api.NewCreateCubeSatVHTransceiverInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error creating new CubeSat VHF transceiver"),
		})
	}

	transceiverResult, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByID(ctx, transceiverID)
	if err != nil {
		return api.NewCreateCubeSatVHTransceiverBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFTransceiverNotFound,
		})
	}
	return api.NewCreateCubeSatVHTransceiverOK().WithPayload(&definition.VHFTransceiver{
		ID:                      transceiverResult.ID,
		Name:                    transceiverResult.Name.String,
		Length:                  transceiverResult.Length.Float64,
		Width:                   transceiverResult.Width.Float64,
		Height:                  transceiverResult.Height.Float64,
		Weight:                  transceiverResult.Weight.Float64,
		SupplyVoltage:           transceiverResult.SupplyVoltage.Float64,
		PowerConsumption:        transceiverResult.PowerConsumption.Float64,
		Interface:               transceiverResult.Interface.String,
		OperatingFrequency:      transceiverResult.OperatingFrequency.Float64,
		OutputPower:             transceiverResult.OutputPower.Float64,
		SensitivityReceiver:     transceiverResult.SensitivityReceiver.Float64,
		MaxOperatingTemperature: transceiverResult.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: transceiverResult.MinOperatingTemperature.Float64,
		MechanicalVibration:     transceiverResult.MechanicalVibration.Int64,
		MechanicalShock:         transceiverResult.MechanicalShock.Int64,
		CreatedAt:               transceiverResult.CreatedAt.Unix(),
		UpdatedAt:               transceiverResult.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) GetCubeSatVHFTransceiverHandler(req api.GetCubeSatVHTransceiverParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get CubeSat VHF transceiver request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetCubeSatVHTransceiverForbidden()
	}

	transceiver, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching CubeSat VHF transceiver", zap.Error(err))
		return api.NewGetCubeSatVHTransceiverBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFTransceiverNotFound,
		})
	}

	return api.NewGetCubeSatVHTransceiverOK().WithPayload(&definition.VHFTransceiver{
		ID:                      transceiver.ID,
		Name:                    transceiver.Name.String,
		Length:                  transceiver.Length.Float64,
		Width:                   transceiver.Width.Float64,
		Height:                  transceiver.Height.Float64,
		Weight:                  transceiver.Weight.Float64,
		SupplyVoltage:           transceiver.SupplyVoltage.Float64,
		PowerConsumption:        transceiver.PowerConsumption.Float64,
		Interface:               transceiver.Interface.String,
		OperatingFrequency:      transceiver.OperatingFrequency.Float64,
		OutputPower:             transceiver.OutputPower.Float64,
		SensitivityReceiver:     transceiver.SensitivityReceiver.Float64,
		MaxOperatingTemperature: transceiver.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: transceiver.MinOperatingTemperature.Float64,
		MechanicalVibration:     transceiver.MechanicalVibration.Int64,
		MechanicalShock:         transceiver.MechanicalShock.Int64,
		CreatedAt:               transceiver.CreatedAt.Unix(),
		UpdatedAt:               transceiver.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) UpdateCubeSatVHFTransceiverHandler(req api.UpdateCubeSatVHTransceiverParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update CubeSat VHF transceiver request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateCubeSatVHTransceiverForbidden()
	}

	transceiver, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching CubeSat VHF transceiver", zap.Error(err))
		return api.NewUpdateCubeSatVHTransceiverBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFTransceiverNotFound,
		})
	}

	if req.UpdateVHFTransceiverBody.Name != "" {
		transceiver.Name.String = req.UpdateVHFTransceiverBody.Name
		transceiver.Name.Valid = true
	}

	if req.UpdateVHFTransceiverBody.Length != 0 {
		transceiver.Length.Float64 = req.UpdateVHFTransceiverBody.Length
		transceiver.Length.Valid = true
	}

	if req.UpdateVHFTransceiverBody.Width != 0 {
		transceiver.Width.Float64 = req.UpdateVHFTransceiverBody.Width
		transceiver.Width.Valid = true
	}

	if req.UpdateVHFTransceiverBody.Height != 0 {
		transceiver.Height.Float64 = req.UpdateVHFTransceiverBody.Height
		transceiver.Height.Valid = true
	}

	if req.UpdateVHFTransceiverBody.Weight != 0 {
		transceiver.Weight.Float64 = req.UpdateVHFTransceiverBody.Weight
		transceiver.Weight.Valid = true
	}

	if req.UpdateVHFTransceiverBody.SupplyVoltage != 0 {
		transceiver.SupplyVoltage.Float64 = req.UpdateVHFTransceiverBody.SupplyVoltage
		transceiver.SupplyVoltage.Valid = true
	}

	if req.UpdateVHFTransceiverBody.PowerConsumption != 0 {
		transceiver.PowerConsumption.Float64 = req.UpdateVHFTransceiverBody.PowerConsumption
		transceiver.PowerConsumption.Valid = true
	}

	if req.UpdateVHFTransceiverBody.Interface != "" {
		transceiver.Interface.String = req.UpdateVHFTransceiverBody.Interface
	}

	if req.UpdateVHFTransceiverBody.OperatingFrequency != 0 {
		transceiver.OperatingFrequency.Float64 = req.UpdateVHFTransceiverBody.OperatingFrequency
	}

	if req.UpdateVHFTransceiverBody.OutputPower != 0 {
		transceiver.OutputPower.Float64 = req.UpdateVHFTransceiverBody.OutputPower
	}

	if req.UpdateVHFTransceiverBody.SensitivityReceiver != 0 {
		transceiver.SensitivityReceiver.Float64 = req.UpdateVHFTransceiverBody.SensitivityReceiver
	}

	if req.UpdateVHFTransceiverBody.MaxOperatingTemperature != 0 {
		transceiver.MaxOperatingTemperature.Float64 = req.UpdateVHFTransceiverBody.MaxOperatingTemperature
	}

	if req.UpdateVHFTransceiverBody.MinOperatingTemperature != 0 {
		transceiver.MinOperatingTemperature.Float64 = req.UpdateVHFTransceiverBody.MinOperatingTemperature
	}

	if req.UpdateVHFTransceiverBody.MechanicalVibration != 0 {
		transceiver.MechanicalVibration.Int64 = req.UpdateVHFTransceiverBody.MechanicalVibration
	}

	if req.UpdateVHFTransceiverBody.MechanicalShock != 0 {
		transceiver.MechanicalShock.Int64 = req.UpdateVHFTransceiverBody.MechanicalShock
	}

	transceiver.UpdatedAt.Time = time.Now()

	err = h.cubeSatVhfTransceiverUsecase.UpdateVHFTransceiver(ctx, transceiver)
	if err != nil {
		zap.L().Error("error updating CubeSat VHF transceiver", zap.Error(err))
		return api.NewUpdateCubeSatVHTransceiverInternalServerError()
	}

	newTransceiver, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByID(ctx, transceiver.ID)
	if err != nil {
		zap.L().Error("error fetching updated CubeSat VHF transceiver", zap.Error(err))
		return api.NewUpdateCubeSatVHTransceiverInternalServerError()
	}

	return api.NewUpdateCubeSatVHTransceiverOK().WithPayload(&definition.VHFTransceiver{
		ID:                      newTransceiver.ID,
		Name:                    newTransceiver.Name.String,
		Length:                  newTransceiver.Length.Float64,
		Width:                   newTransceiver.Width.Float64,
		Height:                  newTransceiver.Height.Float64,
		Weight:                  newTransceiver.Weight.Float64,
		SupplyVoltage:           newTransceiver.SupplyVoltage.Float64,
		PowerConsumption:        newTransceiver.PowerConsumption.Float64,
		Interface:               newTransceiver.Interface.String,
		OperatingFrequency:      newTransceiver.OperatingFrequency.Float64,
		OutputPower:             newTransceiver.OutputPower.Float64,
		SensitivityReceiver:     newTransceiver.SensitivityReceiver.Float64,
		MaxOperatingTemperature: newTransceiver.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: newTransceiver.MinOperatingTemperature.Float64,
		MechanicalVibration:     newTransceiver.MechanicalVibration.Int64,
		MechanicalShock:         newTransceiver.MechanicalShock.Int64,
		CreatedAt:               newTransceiver.CreatedAt.Unix(),
		UpdatedAt:               newTransceiver.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) DeleteCubeSatVHFTransceiverHandler(req api.DeleteCubeSatVHTransceiverParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete CubeSat VHF transceiver request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteCubeSatVHTransceiverForbidden()
	}

	transceiver, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetching CubeSat VHF transceiver", zap.Error(err))
		return api.NewDeleteCubeSatVHTransceiverBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFTransceiverNotFound,
		})
	}

	err = h.cubeSatVhfTransceiverUsecase.DeleteVHFTransceiver(ctx, transceiver.ID)
	if err != nil {
		zap.L().Error("error deleting CubeSat VHF transceiver", zap.Error(err))
		return api.NewDeleteCubeSatVHTransceiverInternalServerError()
	}

	return api.NewDeleteCubeSatVHTransceiverOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("CubeSat VHF transceiver deleted successfully"),
	})
}

func (h *Handler) GetAvailableCubeSatVHFTransceiversHandler(req api.GetAvailableCubeSatVHTransceiversParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch CubeSat VHF transceivers request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetAvailableCubeSatVHTransceiversForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultVHFTransceiversSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.VHFTransceiversSortByCreatedAt
		default:
			sortParams = model.DefaultVHFTransceiversSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterCubeSatVHFTransceiverByMinFrequency != nil {
		filters[model.FilterVHFTransceiverByMinOperatingFrequency] = *req.FilterCubeSatVHFTransceiverByMinFrequency
	}

	if req.FilterCubeSatVHFTransceiverByMaxFrequency != nil {
		filters[model.FilterVHFTransceiverByMaxOperatingFrequency] = *req.FilterCubeSatVHFTransceiverByMaxFrequency
	}

	transceivers, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiversByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetching CubeSat VHF transceivers", zap.Error(err))
		return api.NewGetAvailableCubeSatVHTransceiversBadRequest().WithPayload(&definition.Error{
			Message: &model.VHFTransceiversNotFound,
		})
	}

	transceiversDefinition := h.CubeSatVHFTransceiversToDefinition(ctx, transceivers)

	return api.NewGetAvailableCubeSatVHTransceiversOK().WithPayload(&definition.CubeSatVHFTransceivers{
		Count:           useful.Int64Ptr(int64(len(transceiversDefinition))),
		VhfTransceivers: transceiversDefinition,
	})
}

func (h *Handler) CubeSatVHFTransceiversToDefinition(ctx context.Context, transceivers []model.VHFTransceiver) []*definition.VHFTransceiver {
	transceiversDefinition := make([]*definition.VHFTransceiver, len(transceivers))

	for i := range transceivers {
		transceiversDefinition[i] = &definition.VHFTransceiver{
			ID:                      transceivers[i].ID,
			Name:                    transceivers[i].Name.String,
			Length:                  transceivers[i].Length.Float64,
			Width:                   transceivers[i].Width.Float64,
			Height:                  transceivers[i].Height.Float64,
			Weight:                  transceivers[i].Weight.Float64,
			SupplyVoltage:           transceivers[i].SupplyVoltage.Float64,
			PowerConsumption:        transceivers[i].PowerConsumption.Float64,
			Interface:               transceivers[i].Interface.String,
			OperatingFrequency:      transceivers[i].OperatingFrequency.Float64,
			OutputPower:             transceivers[i].OutputPower.Float64,
			SensitivityReceiver:     transceivers[i].SensitivityReceiver.Float64,
			MaxOperatingTemperature: transceivers[i].MaxOperatingTemperature.Float64,
			MinOperatingTemperature: transceivers[i].MinOperatingTemperature.Float64,
			MechanicalVibration:     transceivers[i].MechanicalVibration.Int64,
			MechanicalShock:         transceivers[i].MechanicalShock.Int64,
			CreatedAt:               transceivers[i].CreatedAt.Unix(),
			UpdatedAt:               transceivers[i].UpdatedAt.Time.Unix(),
		}
	}

	return transceiversDefinition
}
