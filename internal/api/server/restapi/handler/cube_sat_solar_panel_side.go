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

func (h *Handler) CreateSolarPanelSideHandler(req api.CreateSolarPanelSideParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create solar_panel request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateSolarPanelSideForbidden()
	}

	solarPanel := model.CubeSatSolarPanelSide{

		Name:                    sql.NullString{String: req.CreateSolarPanelSideBody.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: *req.CreateSolarPanelSideBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: *req.CreateSolarPanelSideBody.Width, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.Weight, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.MinOperatingTemperature, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.MaxOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: req.CreateSolarPanelSideBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: req.CreateSolarPanelSideBody.MechanicalShock, Valid: true},
		CoilArea:                sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.CoilArea, Valid: true},
		CoilResistance:          sql.NullInt64{Int64: req.CreateSolarPanelSideBody.CoilResistance, Valid: true},
		Efficiency:              sql.NullInt64{Int64: req.CreateSolarPanelSideBody.Efficiency, Valid: true},
		Imp:                     sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.Imp, Valid: true},
		Interface:               sql.NullString{String: req.CreateSolarPanelSideBody.Interface, Valid: true},
		Isc:                     sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.Isc, Valid: true},
		Vmp:                     sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.Vmp, Valid: true},
		Voc:                     sql.NullFloat64{Float64: req.CreateSolarPanelSideBody.Voc, Valid: true},
		Height:                  sql.NullFloat64{Float64: *req.CreateSolarPanelSideBody.Height, Valid: true},
	}

	SolarPanelID, err := h.solarPanelUsecase.CreateSolarPanelSide(ctx, solarPanel)
	if err != nil {
		zap.L().Error("error create new solar panel", zap.Error(err))
		return api.NewCreateSolarPanelSideInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new solar panel"),
		})
	}

	solarPanelResult, err := h.solarPanelUsecase.GetSolarPanelSideByID(ctx, SolarPanelID)
	if err != nil {

		return api.NewCreateSolarPanelSideBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}
	return api.NewCreateSolarPanelSideOK().WithPayload(&definition.SolarPanelSide{
		Name:                    solarPanel.Name.String,
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
		MechanicalVibration:     solarPanelResult.MechanicalVibration.Int64,
		MinOperatingTemperature: solarPanelResult.MinOperatingTemperature.Float64,
		UpdatedAt:               solarPanelResult.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanelResult.Vmp.Float64,
		Voc:                     solarPanelResult.Voc.Float64,
		Weight:                  solarPanelResult.Weight.Float64,
		Width:                   solarPanelResult.Width.Float64,
	})

}

func (h *Handler) GetSolarPanelSideHandler(req api.GetSolarPanelSideParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get solar panel request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetSolarPanelSideForbidden()
	}

	solarPanel, err := h.solarPanelUsecase.GetSolarPanelSideByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel", zap.Error(err))
		return api.NewGetSolarPanelSideBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}
	return api.NewGetSolarPanelSideOK().WithPayload(&definition.SolarPanelSide{
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
		MechanicalVibration:     solarPanel.MechanicalVibration.Int64,
		MinOperatingTemperature: solarPanel.MinOperatingTemperature.Float64,
		UpdatedAt:               solarPanel.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanel.Vmp.Float64,
		Voc:                     solarPanel.Voc.Float64,
		Weight:                  solarPanel.Weight.Float64,
		Width:                   solarPanel.Width.Float64,
	})
}

func (h *Handler) UpdateCubeSatSolarPanelSideHandler(req api.UpdateCubeSatSolarPanelSideParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update solar panel request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateCubeSatSolarPanelSideForbidden()
	}

	solarPanel, err := h.solarPanelUsecase.GetSolarPanelSideByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelSideBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}

	if req.UpdateSolarPanelSideBody.Name != "" {
		solarPanel.Name.String = req.UpdateSolarPanelSideBody.Name
		solarPanel.Name.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Length != nil {
		solarPanel.Length.Float64 = *req.UpdateSolarPanelSideBody.Length
		solarPanel.Length.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Width != nil {
		solarPanel.Width.Float64 = *req.UpdateSolarPanelSideBody.Width
		solarPanel.Width.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Height != nil {
		solarPanel.Height.Float64 = *req.UpdateSolarPanelSideBody.Height
		solarPanel.Height.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Weight != 0 {
		solarPanel.Weight.Float64 = req.UpdateSolarPanelSideBody.Weight
		solarPanel.Weight.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Interface != "" {
		solarPanel.Interface.String = req.UpdateSolarPanelSideBody.Interface
		solarPanel.Interface.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Voc != 0 {
		solarPanel.Voc.Float64 = req.UpdateSolarPanelSideBody.Voc
		solarPanel.Voc.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Isc != 0 {
		solarPanel.Isc.Float64 = req.UpdateSolarPanelSideBody.Isc
		solarPanel.Isc.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Vmp != 0 {
		solarPanel.Vmp.Float64 = req.UpdateSolarPanelSideBody.Vmp
		solarPanel.Vmp.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Imp != 0 {
		solarPanel.Imp.Float64 = req.UpdateSolarPanelSideBody.Imp
		solarPanel.Imp.Valid = true
	}

	if req.UpdateSolarPanelSideBody.Efficiency != 0 {
		solarPanel.Efficiency.Int64 = req.UpdateSolarPanelSideBody.Efficiency
		solarPanel.Efficiency.Valid = true
	}

	if req.UpdateSolarPanelSideBody.CoilArea != 0 {
		solarPanel.CoilArea.Float64 = req.UpdateSolarPanelSideBody.CoilArea
		solarPanel.CoilArea.Valid = true
	}

	if req.UpdateSolarPanelSideBody.CoilResistance != 0 {
		solarPanel.CoilResistance.Int64 = req.UpdateSolarPanelSideBody.CoilResistance
		solarPanel.CoilResistance.Valid = true
	}

	if req.UpdateSolarPanelSideBody.MaxOperatingTemperature != 0 {
		solarPanel.MaxOperatingTemperature.Float64 = req.UpdateSolarPanelSideBody.MaxOperatingTemperature
		solarPanel.MaxOperatingTemperature.Valid = true
	}

	if req.UpdateSolarPanelSideBody.MinOperatingTemperature != 0 {
		solarPanel.MinOperatingTemperature.Float64 = req.UpdateSolarPanelSideBody.MinOperatingTemperature
		solarPanel.MinOperatingTemperature.Valid = true
	}

	if req.UpdateSolarPanelSideBody.MechanicalVibration != 0 {
		solarPanel.MechanicalVibration.Int64 = req.UpdateSolarPanelSideBody.MechanicalVibration
		solarPanel.MechanicalVibration.Valid = true
	}

	if req.UpdateSolarPanelSideBody.MechanicalShock != 0 {
		solarPanel.MechanicalShock.Int64 = req.UpdateSolarPanelSideBody.MechanicalShock
		solarPanel.MechanicalShock.Valid = true
	}

	solarPanel.UpdatedAt.Time = time.Now()

	err = h.solarPanelUsecase.UpdateSolarPanelSide(ctx, solarPanel)
	if err != nil {
		zap.L().Error("error update solar panel side", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelSideInternalServerError()
	}

	newSolarPanel, err := h.solarPanelUsecase.GetSolarPanelSideByID(ctx, solarPanel.ID)
	if err != nil {
		zap.L().Error("error fetch dolar panel side", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelSideInternalServerError()
	}

	return api.NewUpdateCubeSatSolarPanelSideOK().WithPayload(&definition.SolarPanelSide{
		CoilArea:                newSolarPanel.CoilArea.Float64,
		CoilResistance:          newSolarPanel.CoilResistance.Int64,
		CreatedAt:               newSolarPanel.CreatedAt.Unix(),
		Efficiency:              newSolarPanel.Efficiency.Int64,
		Height:                  newSolarPanel.Height.Float64,
		ID:                      &newSolarPanel.ID,
		Imp:                     newSolarPanel.Imp.Float64,
		Interface:               newSolarPanel.Interface.String,
		Isc:                     newSolarPanel.Isc.Float64,
		Length:                  newSolarPanel.Length.Float64,
		MaxOperatingTemperature: newSolarPanel.MaxOperatingTemperature.Float64,
		MechanicalShock:         newSolarPanel.MechanicalShock.Int64,
		MechanicalVibration:     newSolarPanel.MechanicalVibration.Int64,
		MinOperatingTemperature: newSolarPanel.MinOperatingTemperature.Float64,
		UpdatedAt:               newSolarPanel.UpdatedAt.Time.Unix(),
		Vmp:                     newSolarPanel.Vmp.Float64,
		Voc:                     newSolarPanel.Voc.Float64,
		Weight:                  newSolarPanel.Weight.Float64,
		Width:                   newSolarPanel.Width.Float64,
	})
}

func (h *Handler) DeleteCubeSatSolarPanelSideHandler(req api.DeleteCubeSatSolarPanelSideParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete solar panel request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteSolarPanelSideForbidden()
	}

	solarPanel, err := h.solarPanelUsecase.GetSolarPanelSideByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel", zap.Error(err))
		return api.NewDeleteCubeSatSolarPanelSideBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelNotFound,
		})
	}

	err = h.solarPanelUsecase.DeleteSolarPanelSide(ctx, solarPanel.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel", zap.Error(err))
		return api.NewDeleteCubeSatSolarPanelSideInternalServerError()
	}

	return api.NewDeleteSolarPanelSideOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("Cube sat solar pane deleted successfully"),
	})
}

func (h *Handler) GetCubeSatSolarPanelsSide(req api.GetCubeSatSolarPanelsSideParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch solar panels request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetCubeSatSolarPanelsSideForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultCubeSatSolarPanelsSideSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.CubeSatSolarPanelsSideSortByCreatedAt
		default:
			sortParams = model.DefaultCubeSatSolarPanelsSideSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterCubeSatSolarPanelByLengthMax != nil {
		filters[model.FilterCubeSatSolarPanelSideByMaxLength] = *req.FilterCubeSatSolarPanelByLengthMax
	}

	if req.FilterCubeSatSolarPanelByLengthMin != nil {
		filters[model.FilterCubeSatSolarPanelSideByMinLength] = *req.FilterCubeSatSolarPanelByLengthMin
	}

	solarPanelsSide, err := h.solarPanelUsecase.GetSolarPanelSideByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetch solar panels", zap.Error(err))
		return api.NewGetCubeSatSolarPanelsSideBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelsNotFound,
		})
	}

	solarPanelsSideDefinition := h.SolarPanelsSideToDefinition(ctx, solarPanelsSide)

	return api.NewGetCubeSatSolarPanelsSideOK().WithPayload(&definition.CubeSatSolarPanelsSide{
		Count:                  useful.Int64Ptr(int64(len(solarPanelsSideDefinition))),
		CubeSatSolarPanelsSide: solarPanelsSideDefinition,
	})
}

func (h *Handler) SolarPanelsSideToDefinition(ctx context.Context, cubeSatSolarPanelsSide []model.CubeSatSolarPanelSide) []*definition.SolarPanelSide {
	cubeSatSolarPanelsSideDefinition := make([]*definition.SolarPanelSide, len(cubeSatSolarPanelsSide))

	for i := range cubeSatSolarPanelsSide {
		cubeSatSolarPanelsSideDefinition[i] = &definition.SolarPanelSide{
			Name:                    cubeSatSolarPanelsSide[i].Name.String,
			CoilArea:                cubeSatSolarPanelsSide[i].CoilArea.Float64,
			CoilResistance:          cubeSatSolarPanelsSide[i].CoilResistance.Int64,
			CreatedAt:               cubeSatSolarPanelsSide[i].CreatedAt.Unix(),
			Efficiency:              cubeSatSolarPanelsSide[i].Efficiency.Int64,
			Height:                  cubeSatSolarPanelsSide[i].Height.Float64,
			ID:                      &cubeSatSolarPanelsSide[i].ID,
			Imp:                     cubeSatSolarPanelsSide[i].Imp.Float64,
			Interface:               cubeSatSolarPanelsSide[i].Interface.String,
			Isc:                     cubeSatSolarPanelsSide[i].Isc.Float64,
			Length:                  cubeSatSolarPanelsSide[i].Length.Float64,
			MaxOperatingTemperature: cubeSatSolarPanelsSide[i].MaxOperatingTemperature.Float64,
			MechanicalShock:         cubeSatSolarPanelsSide[i].MechanicalShock.Int64,
			MechanicalVibration:     cubeSatSolarPanelsSide[i].MechanicalVibration.Int64,
			MinOperatingTemperature: cubeSatSolarPanelsSide[i].MinOperatingTemperature.Float64,
			UpdatedAt:               cubeSatSolarPanelsSide[i].UpdatedAt.Time.Unix(),
			Vmp:                     cubeSatSolarPanelsSide[i].Vmp.Float64,
			Voc:                     cubeSatSolarPanelsSide[i].Voc.Float64,
			Weight:                  cubeSatSolarPanelsSide[i].Weight.Float64,
			Width:                   cubeSatSolarPanelsSide[i].Width.Float64,
		}
	}

	return cubeSatSolarPanelsSideDefinition
}
