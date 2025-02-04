package handler

import (
	"fmt"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

// func (h *Handler) CreateSolarPanelHandler(req api.CreateSolarPanelParams, principal *definition.Principal) middleware.Responder {
// 	zap.L().Info("create solar_panel request")
// 	ctx := req.HTTPRequest.Context()

// 	if principal.Role == 0 {
// 		return api.NewCreateSolarPanelForbidden()
// 	}

// 	solar_panel := model.Chassis{

// 		Height:                          *req.CreateChassisVPXBody.Height,
// 		Length:                          *req.CreateChassisVPXBody.Length,
// 	}

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
	zap.L().Info("get solar_panel request, id:" + string(req.ID))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetSolarPanelForbidden()
	}

	solarPanel, err := h.solarPanelUsecase.GetSolarPanelByID(ctx, req.ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(solarPanel.Isc)
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
