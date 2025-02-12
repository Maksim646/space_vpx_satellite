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

func (h *Handler) CreateSolarPanelTopHandler(req api.CreateSolarPanelTopParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create solar_panel_top request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateSolarPanelTopForbidden()
	}

	solarPanelTop := model.SolarPanelTop{
		Name:                    sql.NullString{String: req.CreateSolarPanelTopBody.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: *req.CreateSolarPanelTopBody.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: *req.CreateSolarPanelTopBody.Width, Valid: true},
		Height:                  sql.NullFloat64{Float64: *req.CreateSolarPanelTopBody.Height, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Weight, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.MinOperatingTemperature, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.MaxOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.MechanicalShock, Valid: true},
		CoilArea:                sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.CoilArea, Valid: true},
		CoilResistance:          sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.CoilResistance, Valid: true},
		Efficiency:              sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Efficiency, Valid: true},
		Imp:                     sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Imp, Valid: true},
		Interface:               sql.NullString{String: req.CreateSolarPanelTopBody.Interface, Valid: true},
		Isc:                     sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Isc, Valid: true},
		Vmp:                     sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Vmp, Valid: true},
		Voc:                     sql.NullFloat64{Float64: req.CreateSolarPanelTopBody.Voc, Valid: true},
	}

	SolarPanelTopID, err := h.cubeSatSolarPanelTopUsecase.CreateSolarPanelTop(ctx, solarPanelTop)
	if err != nil {
		zap.L().Error("error create new solar panel top", zap.Error(err))
		return api.NewCreateSolarPanelTopInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new solar panel top"),
		})
	}

	solarPanelTopResult, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByID(ctx, SolarPanelTopID)
	if err != nil {
		return api.NewCreateSolarPanelTopBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelTopNotFound,
		})
	}
	return api.NewCreateSolarPanelTopOK().WithPayload(&definition.SolarPanelTop{
		Name:                    solarPanelTopResult.Name.String,
		CoilArea:                solarPanelTopResult.CoilArea.Float64,
		CoilResistance:          solarPanelTopResult.CoilResistance.Float64,
		CreatedAt:               solarPanelTopResult.CreatedAt.Unix(),
		Efficiency:              solarPanelTopResult.Efficiency.Float64,
		Height:                  solarPanelTopResult.Height.Float64,
		ID:                      &solarPanelTopResult.ID,
		Imp:                     solarPanelTopResult.Imp.Float64,
		Interface:               solarPanelTopResult.Interface.String,
		Isc:                     solarPanelTopResult.Isc.Float64,
		Length:                  solarPanelTopResult.Length.Float64,
		MaxOperatingTemperature: solarPanelTopResult.MaxOperatingTemperature.Float64,
		MechanicalShock:         solarPanelTopResult.MechanicalShock.Float64,
		MechanicalVibration:     solarPanelTopResult.MechanicalVibration.Float64,
		MinOperatingTemperature: solarPanelTopResult.MinOperatingTemperature.Float64,
		UpdatedAt:               solarPanelTopResult.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanelTopResult.Vmp.Float64,
		Voc:                     solarPanelTopResult.Voc.Float64,
		Weight:                  solarPanelTopResult.Weight.Float64,
		Width:                   solarPanelTopResult.Width.Float64,
	})
}

func (h *Handler) GetSolarPanelTopHandler(req api.GetSolarPanelTopParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get solar panel top request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetSolarPanelTopForbidden()
	}

	solarPanelTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel top", zap.Error(err))
		return api.NewGetSolarPanelTopBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelTopNotFound,
		})
	}
	return api.NewGetSolarPanelTopOK().WithPayload(&definition.SolarPanelTop{
		CoilArea:                solarPanelTop.CoilArea.Float64,
		CoilResistance:          solarPanelTop.CoilResistance.Float64,
		CreatedAt:               solarPanelTop.CreatedAt.Unix(),
		Efficiency:              solarPanelTop.Efficiency.Float64,
		Height:                  solarPanelTop.Height.Float64,
		ID:                      &solarPanelTop.ID,
		Imp:                     solarPanelTop.Imp.Float64,
		Interface:               solarPanelTop.Interface.String,
		Isc:                     solarPanelTop.Isc.Float64,
		Length:                  solarPanelTop.Length.Float64,
		MaxOperatingTemperature: solarPanelTop.MaxOperatingTemperature.Float64,
		MechanicalShock:         solarPanelTop.MechanicalShock.Float64,
		MechanicalVibration:     solarPanelTop.MechanicalVibration.Float64,
		MinOperatingTemperature: solarPanelTop.MinOperatingTemperature.Float64,
		UpdatedAt:               solarPanelTop.UpdatedAt.Time.Unix(),
		Vmp:                     solarPanelTop.Vmp.Float64,
		Voc:                     solarPanelTop.Voc.Float64,
		Weight:                  solarPanelTop.Weight.Float64,
		Width:                   solarPanelTop.Width.Float64,
	})
}

func (h *Handler) UpdateCubeSatSolarPanelTopHandler(req api.UpdateCubeSatSolarPanelTopParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update solar panel top request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateCubeSatSolarPanelTopForbidden()
	}

	solarPanelTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel top", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelTopBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelTopNotFound,
		})
	}

	if req.UpdateSolarPanelTopBody.Name != "" {
		solarPanelTop.Name.String = req.UpdateSolarPanelTopBody.Name
		solarPanelTop.Name.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Length != nil {
		solarPanelTop.Length.Float64 = *req.UpdateSolarPanelTopBody.Length
		solarPanelTop.Length.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Width != nil {
		solarPanelTop.Width.Float64 = *req.UpdateSolarPanelTopBody.Width
		solarPanelTop.Width.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Height != nil {
		solarPanelTop.Height.Float64 = *req.UpdateSolarPanelTopBody.Height
		solarPanelTop.Height.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Weight != 0 {
		solarPanelTop.Weight.Float64 = req.UpdateSolarPanelTopBody.Weight
		solarPanelTop.Weight.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Interface != "" {
		solarPanelTop.Interface.String = req.UpdateSolarPanelTopBody.Interface
		solarPanelTop.Interface.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Voc != 0 {
		solarPanelTop.Voc.Float64 = req.UpdateSolarPanelTopBody.Voc
		solarPanelTop.Voc.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Isc != 0 {
		solarPanelTop.Isc.Float64 = req.UpdateSolarPanelTopBody.Isc
		solarPanelTop.Isc.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Vmp != 0 {
		solarPanelTop.Vmp.Float64 = req.UpdateSolarPanelTopBody.Vmp
		solarPanelTop.Vmp.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Imp != 0 {
		solarPanelTop.Imp.Float64 = req.UpdateSolarPanelTopBody.Imp
		solarPanelTop.Imp.Valid = true
	}

	if req.UpdateSolarPanelTopBody.Efficiency != 0 {
		solarPanelTop.Efficiency.Float64 = req.UpdateSolarPanelTopBody.Efficiency
		solarPanelTop.Efficiency.Valid = true
	}

	if req.UpdateSolarPanelTopBody.CoilArea != 0 {
		solarPanelTop.CoilArea.Float64 = req.UpdateSolarPanelTopBody.CoilArea
		solarPanelTop.CoilArea.Valid = true
	}

	if req.UpdateSolarPanelTopBody.CoilResistance != 0 {
		solarPanelTop.CoilResistance.Float64 = req.UpdateSolarPanelTopBody.CoilResistance
		solarPanelTop.CoilResistance.Valid = true
	}

	if req.UpdateSolarPanelTopBody.MaxOperatingTemperature != 0 {
		solarPanelTop.MaxOperatingTemperature.Float64 = req.UpdateSolarPanelTopBody.MaxOperatingTemperature
		solarPanelTop.MaxOperatingTemperature.Valid = true
	}

	if req.UpdateSolarPanelTopBody.MinOperatingTemperature != 0 {
		solarPanelTop.MinOperatingTemperature.Float64 = req.UpdateSolarPanelTopBody.MinOperatingTemperature
		solarPanelTop.MinOperatingTemperature.Valid = true
	}

	if req.UpdateSolarPanelTopBody.MechanicalVibration != 0 {
		solarPanelTop.MechanicalVibration.Float64 = req.UpdateSolarPanelTopBody.MechanicalVibration
		solarPanelTop.MechanicalVibration.Valid = true
	}

	if req.UpdateSolarPanelTopBody.MechanicalShock != 0 {
		solarPanelTop.MechanicalShock.Float64 = req.UpdateSolarPanelTopBody.MechanicalShock
		solarPanelTop.MechanicalShock.Valid = true
	}

	solarPanelTop.UpdatedAt.Time = time.Now()

	err = h.cubeSatSolarPanelTopUsecase.UpdateSolarPanelTop(ctx, solarPanelTop)
	if err != nil {
		zap.L().Error("error update solar panel top", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelTopInternalServerError()
	}

	newSolarPanelTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByID(ctx, solarPanelTop.ID)
	if err != nil {
		zap.L().Error("error fetch dolar panel side", zap.Error(err))
		return api.NewUpdateCubeSatSolarPanelTopInternalServerError()
	}

	return api.NewUpdateCubeSatSolarPanelTopOK().WithPayload(&definition.SolarPanelTop{
		CoilArea:                newSolarPanelTop.CoilArea.Float64,
		CoilResistance:          newSolarPanelTop.CoilResistance.Float64,
		CreatedAt:               newSolarPanelTop.CreatedAt.Unix(),
		Efficiency:              newSolarPanelTop.Efficiency.Float64,
		Height:                  newSolarPanelTop.Height.Float64,
		ID:                      &newSolarPanelTop.ID,
		Imp:                     newSolarPanelTop.Imp.Float64,
		Interface:               newSolarPanelTop.Interface.String,
		Isc:                     newSolarPanelTop.Isc.Float64,
		Length:                  newSolarPanelTop.Length.Float64,
		MaxOperatingTemperature: newSolarPanelTop.MaxOperatingTemperature.Float64,
		MechanicalShock:         newSolarPanelTop.MechanicalShock.Float64,
		MechanicalVibration:     newSolarPanelTop.MechanicalVibration.Float64,
		MinOperatingTemperature: newSolarPanelTop.MinOperatingTemperature.Float64,
		UpdatedAt:               newSolarPanelTop.UpdatedAt.Time.Unix(),
		Vmp:                     newSolarPanelTop.Vmp.Float64,
		Voc:                     newSolarPanelTop.Voc.Float64,
		Weight:                  newSolarPanelTop.Weight.Float64,
		Width:                   newSolarPanelTop.Width.Float64,
	})
}

func (h *Handler) DeleteCubeSatSolarPanelTopHandler(req api.DeleteCubeSatSolarPanelTopParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete solar panel top request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteCubeSatSolarPanelTopForbidden()
	}

	solarPanelTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch solar panel top", zap.Error(err))
		return api.NewDeleteCubeSatSolarPanelTopBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelTopNotFound,
		})
	}

	err = h.cubeSatSolarPanelTopUsecase.DeleteSolarPanelTop(ctx, solarPanelTop.ID)
	if err != nil {
		zap.L().Error("error delete solar panel top", zap.Error(err))
		return api.NewDeleteCubeSatSolarPanelTopInternalServerError()
	}

	return api.NewDeleteCubeSatSolarPanelTopOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("Cube sat solar panel top deleted successfully"),
	})
}

func (h *Handler) GetCubeSatSolarPanelsTop(req api.GetCubeSatSolarPanelsTopParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch solar panels top request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetCubeSatSolarPanelsTopForbidden()
	}

	var sortParams string
	if req.SortField == nil {
		sortParams = model.DefaultSolarPanelTopSort
	} else {
		switch *req.SortField {
		case "created_at":
			sortParams = model.SolarPanelTopSortByCreatedAt
		default:
			sortParams = model.DefaultSolarPanelTopSort
		}
	}

	filters := make(map[string]interface{})

	if req.FilterCubeSatSolarPanelByLengthMax != nil {
		filters[model.FilterSolarPanelTopByMaxLength] = *req.FilterCubeSatSolarPanelByLengthMax
	}

	if req.FilterCubeSatSolarPanelByLengthMin != nil {
		filters[model.FilterSolarPanelTopByMinLength] = *req.FilterCubeSatSolarPanelByLengthMin
	}

	solarPanelsTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetch solar panels top", zap.Error(err))
		return api.NewGetCubeSatSolarPanelsTopBadRequest().WithPayload(&definition.Error{
			Message: &model.SolarPanelsTopNotFound,
		})
	}

	solarPanelsTopDefinition := h.SolarPanelsTopToDefinition(ctx, solarPanelsTop)

	return api.NewGetCubeSatSolarPanelsTopOK().WithPayload(&definition.CubeSatSolarPanelsTop{
		Count:                 useful.Int64Ptr(int64(len(solarPanelsTopDefinition))),
		CubeSatSolarPanelsTop: solarPanelsTopDefinition,
	})
}

func (h *Handler) SolarPanelsTopToDefinition(ctx context.Context, cubeSatSolarPanelsTop []model.SolarPanelTop) []*definition.SolarPanelTop {
	cubeSatSolarPanelsTopDefinition := make([]*definition.SolarPanelTop, len(cubeSatSolarPanelsTop))

	for i := range cubeSatSolarPanelsTop {
		cubeSatSolarPanelsTopDefinition[i] = &definition.SolarPanelTop{
			Name:                    cubeSatSolarPanelsTop[i].Name.String,
			CoilArea:                cubeSatSolarPanelsTop[i].CoilArea.Float64,
			CoilResistance:          cubeSatSolarPanelsTop[i].CoilResistance.Float64,
			CreatedAt:               cubeSatSolarPanelsTop[i].CreatedAt.Unix(),
			Efficiency:              cubeSatSolarPanelsTop[i].Efficiency.Float64,
			Height:                  cubeSatSolarPanelsTop[i].Height.Float64,
			ID:                      &cubeSatSolarPanelsTop[i].ID,
			Imp:                     cubeSatSolarPanelsTop[i].Imp.Float64,
			Interface:               cubeSatSolarPanelsTop[i].Interface.String,
			Isc:                     cubeSatSolarPanelsTop[i].Isc.Float64,
			Length:                  cubeSatSolarPanelsTop[i].Length.Float64,
			MaxOperatingTemperature: cubeSatSolarPanelsTop[i].MaxOperatingTemperature.Float64,
			MechanicalShock:         cubeSatSolarPanelsTop[i].MechanicalShock.Float64,
			MechanicalVibration:     cubeSatSolarPanelsTop[i].MechanicalVibration.Float64,
			MinOperatingTemperature: cubeSatSolarPanelsTop[i].MinOperatingTemperature.Float64,
			UpdatedAt:               cubeSatSolarPanelsTop[i].UpdatedAt.Time.Unix(),
			Vmp:                     cubeSatSolarPanelsTop[i].Vmp.Float64,
			Voc:                     cubeSatSolarPanelsTop[i].Voc.Float64,
			Weight:                  cubeSatSolarPanelsTop[i].Weight.Float64,
			Width:                   cubeSatSolarPanelsTop[i].Width.Float64,
		}
	}

	return cubeSatSolarPanelsTopDefinition
}
