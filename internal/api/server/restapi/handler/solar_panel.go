package handler

import (
	"database/sql"
	"strconv"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateSolarPanelHandler(req api.CreateSolarPanelParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create solar_panel request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateSolarPanelForbidden()
	}

	solarPanel := model.SolarPanel{

		Length:                  sql.NullFloat64{Float64: *req.CreateSolarPanelBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: *req.CreateSolarPanelBody.Width, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.CreateSolarPanelBody.Weight, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelBody.MinOperatingTemperature, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelBody.MaxOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: req.CreateSolarPanelBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: req.CreateSolarPanelBody.MechanicalShock, Valid: true},
		CoilArea:                sql.NullFloat64{Float64: req.CreateSolarPanelBody.CoilArea, Valid: true},
		CoilResistance:          sql.NullInt64{Int64: req.CreateSolarPanelBody.CoilResistance, Valid: true},
		Efficiency:              sql.NullInt64{Int64: req.CreateSolarPanelBody.Efficiency, Valid: true},
		Imp:                     sql.NullFloat64{Float64: req.CreateSolarPanelBody.Imp, Valid: true},
		Interface:               sql.NullString{String: req.CreateSolarPanelBody.Interface, Valid: true},
		Isc:                     sql.NullFloat64{Float64: req.CreateSolarPanelBody.Isc, Valid: true},
		Vmp:                     sql.NullFloat64{Float64: req.CreateSolarPanelBody.Vmp, Valid: true},
		Voc:                     sql.NullFloat64{Float64: req.CreateSolarPanelBody.Voc, Valid: true},
		Height:                  sql.NullFloat64{Float64: *req.CreateSolarPanelBody.Height, Valid: true},
	}

	SolarPanelID, err := h.solarPanelUsecase.CreateSolarPanel(ctx, solarPanel)
	if err != nil {
		zap.L().Error("error create new solar panel", zap.Error(err))
		return api.NewCreateSolarPanelInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new solar panel"),
		})
	}

	solarPanelResult, err := h.solarPanelUsecase.GetSolarPanelByID(ctx, SolarPanelID)
	if err != nil {

		return api.NewCreateSolarPanelBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}
	return api.NewCreateSolarPanelOK().WithPayload(&definition.SolarPanel{
		CoilArea:                solarPanelResult.CoilArea.Float64,
		CoilResistance:          solarPanelResult.CoilResistance.Int64,
		CreatedAt:               solarPanelResult.CreatedAt.Unix(),
		Efficiency:              solarPanelResult.Efficiency.Int64,
		Height:                  solarPanelResult.Height.Float64,
		ID:                      &solarPanelResult.ID,
		Imp:                     solarPanelResult.Imp.Float64,
		Interface:               solarPanelResult.Interface.String,
		Isc:                     solarPanelResult.Isc.Float64,
		Length:                  solarPanelResult.Length.Float64,
		MaxOperatingTemperature: solarPanelResult.MaxOperatingTemperature.Float64,
		MechanicalShock:         solarPanelResult.MechanicalShock.Int64,
		MechanicalVibration:     solarPanelResult.MechanicalShock.Int64,
		MinOperatingTemperature: solarPanelResult.MaxOperatingTemperature.Float64,
		UpdatedAt:               solarPanelResult.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanelResult.Vmp.Float64,
		Voc:                     solarPanelResult.Voc.Float64,
		Weight:                  solarPanelResult.Weight.Float64,
		Width:                   solarPanelResult.Width.Float64,
	})

}

// func (h *Handler) DeleteSolarPanelHandler(req api.DeleteSolarPanelParams, principal *definition.Principal) middleware.Responder {
// 	zap.L().Info("update solar_panel request, id:" + string(req.ID))
// 	ctx := req.HTTPRequest.Context()

// 	// Проверка роли пользователя
// 	if principal.Role == 0 {
// 		return api.NewDeleteSolarPanelForbidden()
// 	}

// 	solar_panel, err := h.solarPanelUsecase.GetSolarPanelByID(ctx, req.ID)
// 	if err != nil {
// 		return api.NewDeleteChassisVPXNotFound().WithPayload(&definition.Error{
// 			Message: &model.ChassisNotFound,
// 		})
// 	}

// 	err = h.solarPanelUsecase.DeleteSolarPanelByID(ctx, solar_panel.ID)
// 	if err != nil {
// 		zap.L().Error("failed to delete solar_panel", zap.Error(err))
// 		return api.NewDeleteSolarPanelInternalServerError()
// 	}

// 	return api.NewDeleteSolarPanelOK().WithPayload(&definition.Error{
// 		Message: useful.StrPtr("SolarPanel deleted successfully"),
// 	})
// }

func (h *Handler) GetSolarPanelHandler(req api.GetSolarPanelParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get solar_panel request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetSolarPanelForbidden()
	}

	solarPanel, err := h.solarPanelUsecase.GetSolarPanelByID(ctx, req.ID)
	if err != nil {

		return api.NewGetSolarPanelBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}
	return api.NewGetSolarPanelOK().WithPayload(&definition.SolarPanel{
		CoilArea:                solarPanel.CoilArea.Float64,
		CoilResistance:          solarPanel.CoilResistance.Int64,
		CreatedAt:               solarPanel.CreatedAt.Unix(),
		Efficiency:              solarPanel.Efficiency.Int64,
		Height:                  solarPanel.Height.Float64,
		ID:                      &solarPanel.ID,
		Imp:                     solarPanel.Imp.Float64,
		Interface:               solarPanel.Interface.String,
		Isc:                     solarPanel.Isc.Float64,
		Length:                  solarPanel.Length.Float64,
		MaxOperatingTemperature: solarPanel.MaxOperatingTemperature.Float64,
		MechanicalShock:         solarPanel.MechanicalShock.Int64,
		MechanicalVibration:     solarPanel.MechanicalShock.Int64,
		MinOperatingTemperature: solarPanel.MaxOperatingTemperature.Float64,
		UpdatedAt:               solarPanel.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanel.Vmp.Float64,
		Voc:                     solarPanel.Voc.Float64,
		Weight:                  solarPanel.Weight.Float64,
		Width:                   solarPanel.Width.Float64,
	})
}
