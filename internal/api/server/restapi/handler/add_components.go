package handler

import (
	"fmt"
	"strconv"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) AddCubeSatFrameHandler(req api.UpdateCubeSatFrameByProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("add cube sat frame request")
	ctx := req.HTTPRequest.Context()

	if principal.Role != 0 {
		return api.NewUpdateCubeSatFrameByProjectForbidden()
	}

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ProjectID)
	if err != nil {
		zap.L().Error("error fetch cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	frame, err := h.cubeSatFrameUsecase.GetCubeSatFrameByName(ctx, *req.AddCubeSatFrame.FrameName)
	if err != nil {
		zap.L().Error("no such cube sat frame", zap.Error(err))
		return api.NewUpdateCubeSatFrameBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFrameNotFound,
		})
	}

	if project.Size != frame.Size {
		return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr(fmt.Sprintf(model.InvalidSize + strconv.Itoa(int(project.Size)))),
		})
	}

	if project.PowerSystemName.Valid {
		powerSystem, err := h.cubeSatPowerSystemUsecase.GetPowerSystemByName(ctx, project.PowerSystemName.String)
		if powerSystem.DataInterface.String != "" || err == nil {
			if frame.Interface.String != powerSystem.DataInterface.String {
				return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
					Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + powerSystem.DataInterface.String)),
				})
			}
		} else {
			zap.L().Error("error fetch cube sat power system", zap.Error(err))
			return api.NewUpdateCubeSatFrameByProjectInternalServerError()
		}
	}

	// if project.SolarPanelSideName.Valid {
	// 	solarSide, err := h.cubeSatSolarPanelSideUsecase.GetSolarPanelSideByName(ctx, project.SolarPanelSideName.String)
	// 	if solarSide.Interface.String != "" || err == nil {
	// 		if frame.Interface.String != solarSide.Interface.String {
	// 			return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
	// 				Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + solarSide.Interface.String)),
	// 			})
	// 		}
	// 	} else {
	// 		zap.L().Error("error fetch cube sat solar panel side", zap.Error(err))
	// 		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	// 	}
	// }

	// if project.SolarPanelTopName.Valid {
	// 	solarTop, err := h.cubeSatSolarPanelTopUsecase.GetSolarPanelTopByName(ctx, project.SolarPanelTopName.String)
	// 	if solarTop.Interface.String != "" || err == nil {
	// 		if frame.Interface.String != solarTop.Interface.String {
	// 			return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
	// 				Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + solarTop.Interface.String)),
	// 			})
	// 		}
	// 	} else {
	// 		zap.L().Error("error fetch cube sat solar panel top", zap.Error(err))
	// 		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	// 	}
	// }

	// if project.BoardComputingModuleName.Valid {
	// 	boardComputingModule, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByName(ctx, project.BoardComputingModuleName.String)
	// 	if boardComputingModule.Interface.String != "" || err == nil {
	// 		if frame.Interface.String != boardComputingModule.Interface.String {
	// 			return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
	// 				Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + boardComputingModule.Interface.String)),
	// 			})
	// 		}
	// 	} else {
	// 		zap.L().Error("error fetch cube sat board computing module", zap.Error(err))
	// 		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	// 	}
	// }

	// if project.VHFAntennaSystemName.Valid {
	// 	vhfAntennaSystem, err := h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByName(ctx, project.VHFAntennaSystemName.String)
	// 	if vhfAntennaSystem.Interface.String != "" || err == nil {
	// 		if frame.Interface.String != vhfAntennaSystem.Interface.String {
	// 			return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
	// 				Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + vhfAntennaSystem.Interface.String)),
	// 			})
	// 		}
	// 	} else {
	// 		zap.L().Error("error fetch cube sat vhf antenna system", zap.Error(err))
	// 		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	// 	}
	// }

	// if project.VhfTransceiverName.Valid {
	// 	vhfTransceiverSystem, err := h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByName(ctx, project.VhfTransceiverName.String)
	// 	if vhfTransceiverSystem.Interface.String != "" || err == nil {
	// 		if frame.Interface.String != vhfTransceiverSystem.Interface.String {
	// 			return api.NewUpdateCubeSatFrameByProjectBadRequest().WithPayload(&definition.Error{
	// 				Message: useful.StrPtr(fmt.Sprintf(model.InvalidInterface + vhfTransceiverSystem.Interface.String)),
	// 			})
	// 		}
	// 	} else {
	// 		zap.L().Error("error fetch cube sat vhf transceiver system", zap.Error(err))
	// 		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	// 	}
	// }

	project.FrameName.String = frame.Name.String

	err = h.projectUsecase.UpdateProjectByID(ctx, project)
	if err != nil {
		zap.L().Error("error update cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatFrameByProjectInternalServerError()
	}

	result := h.ProjectToDefinition(ctx, project)

	return api.NewUpdateCubeSatFrameByProjectOK().WithPayload(&result)

}

func (h *Handler) AddCubeSatBoardComputingModule(req api.UpdateCubeSatBoardComputingModuleByProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("add cube sat board computing module request")
	ctx := req.HTTPRequest.Context()

	var powerSystem model.CubeSatPowerSystem
	var vhfTransceiver model.VHFTransceiver
	var vhfAntennaSystem model.VHFAntennaSystem

	if principal.Role != 0 {
		return api.NewUpdateCubeSatFrameByProjectForbidden()
	}

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ProjectID)
	if err != nil {
		zap.L().Error("error fetch cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	boardComputingModule, err := h.cubeSatBoardComputingModuleUsecase.GetBoardComputingModuleByName(ctx, *req.AddCubeSatBoardComputingModule.BoardComputingModuleName)
	if err != nil {
		zap.L().Error("no such cube sat board computing module", zap.Error(err))
		return api.NewUpdateCubeSatFrameBadRequest().WithPayload(&definition.Error{
			Message: &model.CubeSatFrameNotFound,
		})
	}

	if project.PowerSystemName.Valid {
		powerSystem, err = h.cubeSatPowerSystemUsecase.GetPowerSystemByName(ctx, project.PowerSystemName.String)
		if err != nil {
			zap.L().Error("error fetch cube sat power system", zap.Error(err))
			return api.NewUpdateCubeSatBoardComputingModuleByProjectInternalServerError()
		}
	}

	if project.PowerSystemName.Valid {
		if !h.PowerSystemBoardComputingModuleСompatibilityСheck(boardComputingModule, powerSystem) {
			zap.L().Error("error power system compatibility", zap.Error(err))
			return api.NewUpdateBoardComputingModuleUnprocessableEntity().WithPayload(&definition.Error{
				Message: &model.PowerSystemDoesNotCompatibility,
			})
		}
	}

	if project.VHFAntennaSystemName.Valid {
		vhfAntennaSystem, err = h.cubeSatVhfAntennaSystemUsecase.GetVHFAntennaSystemByName(ctx, project.VHFAntennaSystemName.String)
		if err != nil {
			zap.L().Error("error fetch cube sat vhf antenna system", zap.Error(err))
			return api.NewUpdateCubeSatBoardComputingModuleByProjectInternalServerError()
		}

		if !h.VhfAntennaSystemBoardComputingSystemСompatibilityСheck(boardComputingModule, vhfAntennaSystem) {
			zap.L().Error("error vhf antenna system compatibility", zap.Error(err))
			return api.NewUpdateBoardComputingModuleUnprocessableEntity().WithPayload(&definition.Error{
				Message: &model.PowerSystemDoesNotCompatibility,
			})
		}
	}

	if project.VhfTransceiverName.Valid {
		vhfTransceiver, err = h.cubeSatVhfTransceiverUsecase.GetVHFTransceiverByName(ctx, project.VhfTransceiverName.String)
		if err != nil {
			zap.L().Error("error fetch cube sat vhf antenna system", zap.Error(err))
			return api.NewUpdateBoardComputingModuleInternalServerError()
		}

		if !h.VhfTransceiverBoardComputingSystemСompatibilityСheck(boardComputingModule, vhfTransceiver) {
			zap.L().Error("error vhf transceiver compatibility", zap.Error(err))
			return api.NewUpdateBoardComputingModuleUnprocessableEntity().WithPayload(&definition.Error{
				Message: &model.PowerSystemDoesNotCompatibility,
			})
		}
	}

	project.BoardComputingModuleName.String = boardComputingModule.Name.String

	err = h.projectUsecase.UpdateProjectByID(ctx, project)
	if err != nil {
		zap.L().Error("error update cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatBoardComputingModuleByProjectInternalServerError()
	}

	result := h.ProjectToDefinition(ctx, project)

	return api.NewUpdateCubeSatBoardComputingModuleByProjectOK().WithPayload(&result)
}

// func (h *Handler) AddCubeSatSolarPanelSideHandler(req api.UpdateCubeSatSolarPanelSideByProjectParams, principal *definition.Principal) middleware.Responder {
// 	zap.L().Info("add cube sat solar panel side request")
// 	ctx := req.HTTPRequest.Context()

// 	if principal.Role != 0 {
// 		return api.NewUpdateCubeSatSolarPanelSideForbidden()
// 	}

// 	project, err := h.projectUsecase.GetProjectByID(ctx, req.ProjectID)
// 	if err != nil {
// 		zap.L().Error("error fetch cube sat project", zap.Error(err))
// 		return api.NewUpdateCubeSatSolarPanelSideBadRequest().WithPayload(&definition.Error{
// 			Message: &model.SolarPanelSideNotFound,
// 		})
// 	}

// }
