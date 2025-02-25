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

func (h *Handler) CreateBoardComputeringModuleHandler(req api.CreateBoardComputingModuleParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create board computering module request")
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewCreateBoardComputingModuleForbidden()
	}

	boardComputeringModule := model.BoardComputingModule{
		Name:                    sql.NullString{String: req.BoardComputingModule.Name, Valid: true},
		Length:                  sql.NullFloat64{Float64: req.BoardComputingModule.Length, Valid: true},
		Width:                   sql.NullFloat64{Float64: req.BoardComputingModule.Width, Valid: true},
		Height:                  sql.NullFloat64{Float64: req.BoardComputingModule.Height, Valid: true},
		Weight:                  sql.NullFloat64{Float64: req.BoardComputingModule.Weight, Valid: true},
		SupplyVoltage:           sql.NullFloat64{Float64: req.BoardComputingModule.SupplyVoltage, Valid: true},
		PowerConsumption:        sql.NullFloat64{Float64: req.BoardComputingModule.PowerConsumption, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: req.BoardComputingModule.MinOperatingTemperature, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: req.BoardComputingModule.MaxOperatingTemperature, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: req.BoardComputingModule.MechanicalVibration, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: req.BoardComputingModule.MechanicalShock, Valid: true},
		Interface:               sql.NullString{String: req.BoardComputingModule.Interface, Valid: true},
	}

	boardComputeringModuleID, err := h.cubeSatBoardComputingModuleUsecase.CreateBoardComputingModule(ctx, boardComputeringModule)
	if err != nil {
		zap.L().Error("error create new board computering module", zap.Error(err))
		return api.NewCreateBoardComputingModuleInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("error create new board computering module"),
		})
	}

	boardComputeringModuleResult, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByID(ctx, boardComputeringModuleID)
	if err != nil {
		return api.NewCreateBoardComputingModuleBadRequest().WithPayload(&definition.Error{
			Message: &model.BoardComputingModuleNotFound,
		})
	}
	return api.NewCreateBoardComputingModuleCreated().WithPayload(&definition.BoardComputingModule{
		ID:                      boardComputeringModuleResult.ID,
		Name:                    boardComputeringModuleResult.Name.String,
		Length:                  boardComputeringModuleResult.Length.Float64,
		Width:                   boardComputeringModuleResult.Width.Float64,
		Height:                  boardComputeringModuleResult.Height.Float64,
		Weight:                  boardComputeringModuleResult.Weight.Float64,
		SupplyVoltage:           boardComputeringModuleResult.SupplyVoltage.Float64,
		PowerConsumption:        boardComputeringModuleResult.PowerConsumption.Float64,
		CreatedAt:               boardComputeringModuleResult.CreatedAt.Unix(),
		Interface:               boardComputeringModuleResult.Interface.String,
		MaxOperatingTemperature: boardComputeringModuleResult.MaxOperatingTemperature.Float64,
		MechanicalShock:         boardComputeringModuleResult.MechanicalShock.Int64,
		MechanicalVibration:     boardComputeringModuleResult.MechanicalVibration.Int64,
		MinOperatingTemperature: boardComputeringModuleResult.MinOperatingTemperature.Float64,
		UpdatedAt:               boardComputeringModuleResult.UpdatedAt.Time.Unix(),
	})
}

func (h *Handler) GetBoardComputingModuleHandler(req api.GetBoardComputingModuleByIDParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get board computing module request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetBoardComputingModuleByIDForbidden()
	}

	boardComputingModule, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch board computing module", zap.Error(err))
		return api.NewGetBoardComputingModuleByIDBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr(model.BoardComputingModuleNotFound),
		})
	}

	payload := &definition.BoardComputingModule{
		ID:                      boardComputingModule.ID,
		Name:                    boardComputingModule.Name.String,
		Length:                  boardComputingModule.Length.Float64,
		Width:                   boardComputingModule.Width.Float64,
		Height:                  boardComputingModule.Height.Float64,
		Weight:                  boardComputingModule.Weight.Float64,
		SupplyVoltage:           boardComputingModule.SupplyVoltage.Float64,
		PowerConsumption:        boardComputingModule.PowerConsumption.Float64,
		Interface:               boardComputingModule.Interface.String,
		MaxOperatingTemperature: boardComputingModule.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: boardComputingModule.MinOperatingTemperature.Float64,
		MechanicalVibration:     boardComputingModule.MechanicalVibration.Int64,
		MechanicalShock:         boardComputingModule.MechanicalShock.Int64,
		UpdatedAt:               boardComputingModule.UpdatedAt.Time.Unix(),
	}

	return api.NewGetBoardComputingModuleByIDOK().WithPayload(payload)
}

func (h *Handler) UpdateBoardComputingModuleHandler(req api.UpdateBoardComputingModuleParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update board computing module request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewUpdateBoardComputingModuleForbidden()
	}

	boardComputingModule, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch board computing module", zap.Error(err))
		return api.NewUpdateBoardComputingModuleBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr(model.BoardComputingModuleNotFound),
		})
	}

	if req.BoardComputingModule.Name != "" {
		boardComputingModule.Name.String = req.BoardComputingModule.Name
		boardComputingModule.Name.Valid = true
	}

	if req.BoardComputingModule.Length != 0 {
		boardComputingModule.Length.Float64 = req.BoardComputingModule.Length
		boardComputingModule.Length.Valid = true
	}

	if req.BoardComputingModule.Width != 0 {
		boardComputingModule.Width.Float64 = req.BoardComputingModule.Width
		boardComputingModule.Width.Valid = true
	}

	if req.BoardComputingModule.Height != 0 {
		boardComputingModule.Height.Float64 = req.BoardComputingModule.Height
		boardComputingModule.Height.Valid = true
	}

	if req.BoardComputingModule.Weight != 0 {
		boardComputingModule.Weight.Float64 = req.BoardComputingModule.Weight
		boardComputingModule.Weight.Valid = true
	}

	if req.BoardComputingModule.SupplyVoltage != 0 {
		boardComputingModule.SupplyVoltage.Float64 = req.BoardComputingModule.SupplyVoltage
		boardComputingModule.SupplyVoltage.Valid = true
	}
	if req.BoardComputingModule.PowerConsumption != 0 {
		boardComputingModule.PowerConsumption.Float64 = req.BoardComputingModule.PowerConsumption
		boardComputingModule.PowerConsumption.Valid = true
	}

	if req.BoardComputingModule.Interface != "" {
		boardComputingModule.Interface.String = req.BoardComputingModule.Interface
		boardComputingModule.Interface.Valid = true
	}

	if req.BoardComputingModule.MaxOperatingTemperature != 0 {
		boardComputingModule.MaxOperatingTemperature.Float64 = req.BoardComputingModule.MaxOperatingTemperature
		boardComputingModule.MaxOperatingTemperature.Valid = true
	}

	if req.BoardComputingModule.MinOperatingTemperature != 0 {
		boardComputingModule.MinOperatingTemperature.Float64 = req.BoardComputingModule.MinOperatingTemperature
		boardComputingModule.MinOperatingTemperature.Valid = true
	}

	if req.BoardComputingModule.MechanicalVibration != 0 {
		boardComputingModule.MechanicalVibration.Int64 = req.BoardComputingModule.MechanicalVibration
	}

	if req.BoardComputingModule.MechanicalShock != 0 {
		boardComputingModule.MechanicalShock.Int64 = req.BoardComputingModule.MechanicalShock
	}

	boardComputingModule.UpdatedAt.Time = time.Now()

	err = h.cubeSatBoardComputingModuleUsecase.UpdateBoardComputingModule(ctx, boardComputingModule)
	if err != nil {
		zap.L().Error("error update board computing module", zap.Error(err))
		return api.NewUpdateBoardComputingModuleInternalServerError()
	}

	newBoardComputingModule, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByID(ctx, boardComputingModule.ID)
	if err != nil {
		zap.L().Error("error fetch board computing module", zap.Error(err))
		return api.NewUpdateBoardComputingModuleInternalServerError()
	}

	// Construct the response payload
	payload := &definition.BoardComputingModule{
		ID:                      newBoardComputingModule.ID,
		Name:                    newBoardComputingModule.Name.String,
		Length:                  newBoardComputingModule.Length.Float64,
		Width:                   newBoardComputingModule.Width.Float64,
		Height:                  newBoardComputingModule.Height.Float64,
		Weight:                  newBoardComputingModule.Weight.Float64,
		SupplyVoltage:           newBoardComputingModule.SupplyVoltage.Float64,
		PowerConsumption:        newBoardComputingModule.PowerConsumption.Float64,
		Interface:               newBoardComputingModule.Interface.String,
		MaxOperatingTemperature: newBoardComputingModule.MaxOperatingTemperature.Float64,
		MinOperatingTemperature: newBoardComputingModule.MinOperatingTemperature.Float64,
		MechanicalVibration:     newBoardComputingModule.MechanicalVibration.Int64,
		MechanicalShock:         newBoardComputingModule.MechanicalShock.Int64,
		UpdatedAt:               newBoardComputingModule.UpdatedAt.Time.Unix(),
	}

	return api.NewUpdateBoardComputingModuleOK().WithPayload(payload)
}

func (h *Handler) DeleteBoardComputingModuleHandler(req api.DeleteBoardComputingModuleParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete board computing module request, id:" + strconv.Itoa(int(req.ID)))
	ctx := req.HTTPRequest.Context()

	if principal.Role == 0 {
		return api.NewDeleteBoardComputingModuleForbidden()
	}

	// Check if BoardComputingModule exists
	_, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch board computing module", zap.Error(err))
		return api.NewDeleteBoardComputingModuleBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr(model.BoardComputingModuleNotFound),
		})
	}

	err = h.cubeSatBoardComputingModuleUsecase.DeleteBoardComputingModule(ctx, req.ID)
	if err != nil {
		zap.L().Error("error delete board computing module", zap.Error(err))
		return api.NewDeleteBoardComputingModuleInternalServerError()
	}

	return api.NewDeleteBoardComputingModuleNoContent()
}

func (h *Handler) GetBoardComputingModules(req api.GetAvailableBoardComputingModulesParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("fetch board computing modules request, userID: " + principal.ID)
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewGetAvailableBoardComputingModulesForbidden()
	}

	sortParams := model.DefaultBoardComputingModulesSort

	if req.SortField != nil && req.SortDirection != nil {
		sortField := *req.SortField
		sortDirection := *req.SortDirection

		switch sortField {
		case "created_at":
			sortParams = "created_at"
		case "name":
			sortParams = "name"
		case "length":
			sortParams = "length"
		case "width":
			sortParams = "width"
		case "height":
			sortParams = "height"
		case "weight":
			sortParams = "weight"
		case "supply_voltage":
			sortParams = "supply_voltage"
		case "power_consumption":
			sortParams = "power_consumption"
		case "max_operating_temperature":
			sortParams = "max_operating_temperature"
		case "min_operating_temperature":
			sortParams = "min_operating_temperature"
		default:
			sortParams = model.DefaultBoardComputingModulesSort
		}

		if sortDirection == "desc" {
			sortParams = sortParams + " DESC" // Assuming you concatenate " DESC" for descending
		}
	}

	filters := make(map[string]interface{})

	if req.FilterBoardComputingModuleByLengthMax != nil {
		filters[model.FilterBoardComputingModuleByMaxLength] = *req.FilterBoardComputingModuleByLengthMax
	}

	if req.FilterBoardComputingModuleByLengthMin != nil {
		filters[model.FilterBoardComputingModuleByMinLength] = *req.FilterBoardComputingModuleByLengthMin
	}

	boardComputingModules, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModulesByFilters(ctx, req.Offset, req.Limit, sortParams, filters)
	if err != nil {
		zap.L().Error("error fetch board computing modules", zap.Error(err))
		return api.NewGetAvailableBoardComputingModulesBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr(model.BoardComputingModulesNotFound),
		})
	}

	boardComputingModulesDefinition := h.BoardComputingModulesToDefinition(ctx, boardComputingModules)

	return api.NewGetAvailableBoardComputingModulesOK().WithPayload(&definition.BoardComputingModuleList{
		Count:                 useful.Int64Ptr(int64(len(boardComputingModulesDefinition))),
		BoardComputingModules: boardComputingModulesDefinition,
	})
}

func (h *Handler) BoardComputingModulesToDefinition(ctx context.Context, boardComputingModules []model.BoardComputingModule) []*definition.BoardComputingModule {
	boardComputingModulesDefinition := make([]*definition.BoardComputingModule, len(boardComputingModules))

	for i := range boardComputingModules {
		boardComputingModulesDefinition[i] = &definition.BoardComputingModule{
			ID:                      boardComputingModules[i].ID,
			Name:                    boardComputingModules[i].Name.String,
			Length:                  boardComputingModules[i].Length.Float64,
			Width:                   boardComputingModules[i].Width.Float64,
			Height:                  boardComputingModules[i].Height.Float64,
			Weight:                  boardComputingModules[i].Weight.Float64,
			SupplyVoltage:           boardComputingModules[i].SupplyVoltage.Float64,
			PowerConsumption:        boardComputingModules[i].PowerConsumption.Float64,
			Interface:               boardComputingModules[i].Interface.String,
			MaxOperatingTemperature: boardComputingModules[i].MaxOperatingTemperature.Float64,
			MinOperatingTemperature: boardComputingModules[i].MinOperatingTemperature.Float64,
			MechanicalVibration:     boardComputingModules[i].MechanicalVibration.Int64,
			MechanicalShock:         boardComputingModules[i].MechanicalShock.Int64,
			UpdatedAt:               boardComputingModules[i].UpdatedAt.Time.Unix(),
		}
	}
	return boardComputingModulesDefinition
}
