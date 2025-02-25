package handler

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateProjectHandler(req api.CreateCubeSatProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create project request")
	ctx := req.HTTPRequest.Context()
	cubeSatProject := model.CubeSatProject{
		Name:                     *req.CreateCubeSatProject.ProjectName,
		UserID:                   principal.ID,
		Size:                     req.CreateCubeSatProject.Size,
		FrameName:                sql.NullString{String: req.CreateCubeSatProject.FrameName, Valid: req.CreateCubeSatProject.FrameName != ""},
		SolarPanelSideName:       sql.NullString{String: req.CreateCubeSatProject.SolarPanelSideName, Valid: req.CreateCubeSatProject.SolarPanelSideName != ""},
		SolarPanelTopName:        sql.NullString{String: req.CreateCubeSatProject.SolarPanelTopName, Valid: req.CreateCubeSatProject.SolarPanelTopName != ""},
		PowerSystemName:          sql.NullString{String: req.CreateCubeSatProject.PowerSystemName, Valid: req.CreateCubeSatProject.PowerSystemName != ""},
		BoardComputingModuleName: sql.NullString{String: req.CreateCubeSatProject.BoardComputingModuleName, Valid: req.CreateCubeSatProject.BoardComputingModuleName != ""},
		VHFAntennaSystemName:     sql.NullString{String: req.CreateCubeSatProject.VhfAntennaSystemName, Valid: req.CreateCubeSatProject.VhfAntennaSystemName != ""},
		VhfTransceiverName:       sql.NullString{String: req.CreateCubeSatProject.VhfTransceiverName, Valid: req.CreateCubeSatProject.VhfTransceiverName != ""},
	}

	projectID, err := h.projectUsecase.CreatedProject(ctx, cubeSatProject)
	if err != nil {
		zap.L().Error("error create project", zap.Error(err))
		return api.NewCreateCubeSatProjectInternalServerError()
	}

	return api.NewCreateCubeSatProjectOK().WithPayload(&definition.CreateCubeSatProjectResponse{
		ProjectID:   useful.StrPtr(projectID),
		ProjectName: *req.CreateCubeSatProject.ProjectName,
	})
}

func (h *Handler) GetProjectByIDHandler(req api.GetCubeSatProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get project request, userID: " + req.ID)
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewGetCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	projectResult := h.ProjectToDefinition(ctx, project)

	return api.NewGetCubeSatProjectOK().WithPayload(&projectResult)
}

func (h *Handler) UpdateProjectHandler(req api.UpdateCubeSatProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update project request")
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error fetch cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	if *req.UpdateCubeSatProjectBody.ProjectName != "" {
		project.Name = *req.UpdateCubeSatProjectBody.ProjectName
	}

	if req.UpdateCubeSatProjectBody.Size != 0 {
		project.Size = req.UpdateCubeSatProjectBody.Size
	}

	err = h.projectUsecase.UpdateProjectByID(ctx, project)
	if err != nil {
		return api.NewUpdateCubeSatProjectInternalServerError()
	}

	newProject, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateCubeSatProjectInternalServerError()
	}

	projectResult := h.ProjectToDefinition(ctx, newProject)

	return api.NewUpdateCubeSatProjectOK().WithPayload(&projectResult)
}

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

func (h *Handler) DeleteProject(req api.DeleteCubeSatProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete project request")
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewDeleteCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	err = h.projectUsecase.DeleteProject(ctx, project.ID)
	if err != nil {
		return api.NewDeleteCubeSatProjectInternalServerError()
	}

	return api.NewDeleteCubeSatProjectOK().WithPayload(&definition.CreateCubeSatProjectResponse{
		ProjectID: useful.StrPtr(project.ID),
	})
}

func (h *Handler) GetProjectsByUser(req api.GetUserCubeSatProjectsParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get user projects request")
	ctx := req.HTTPRequest.Context()

	filters := make(map[string]interface{})
	filters[model.FilterCubeSatProjectsByUser] = principal.ID

	projects, projectsCount, err := h.projectUsecase.GetProjectsByFilters(ctx, req.Offset, req.Limit, *req.SortField, filters)
	if err != nil {
		return api.NewDeleteCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectsNotFound,
		})
	}

	projectsResult := h.ProjectsToDefinition(ctx, projects)

	projectsResponse := definition.CubeSatProjects{
		Count:    useful.Int64Ptr(projectsCount),
		Projects: projectsResult,
	}

	return api.NewGetUserCubeSatProjectsOK().WithPayload(&projectsResponse)
}

func (h *Handler) ProjectToDefinition(ctx context.Context, project model.CubeSatProject) definition.CubeSatProject {
	projectResult := &definition.CubeSatProject{
		ID:                       project.ID,
		UserID:                   project.UserID,
		ProjectName:              &project.Name,
		FrameName:                project.FrameName.String,
		SolarPanelSideName:       project.SolarPanelSideName.String,
		SolarPanelTopName:        project.SolarPanelTopName.String,
		PowerSystemName:          project.PowerSystemName.String,
		BoardComputingModuleName: project.BoardComputingModuleName.String,
		VhfAntennaSystemName:     project.VHFAntennaSystemName.String,
		VhfTransceiverName:       project.VhfTransceiverName.String,
		CreatedAt:                project.CreatedAt.Unix(),
		UpdatedAt:                project.UpdatedAt.Time.Unix(),
	}

	return *projectResult
}

func (h *Handler) ProjectsToDefinition(ctx context.Context, projects []model.CubeSatProject) []*definition.CubeSatProject {
	projectsData := make([]*definition.CubeSatProject, len(projects))

	for i := range projects {
		projectsData[i] = &definition.CubeSatProject{
			ID:                       projects[i].ID,
			ProjectName:              &projects[i].Name,
			UserID:                   projects[i].UserID,
			FrameName:                projects[i].FrameName.String,
			SolarPanelSideName:       projects[i].SolarPanelSideName.String,
			SolarPanelTopName:        projects[i].SolarPanelTopName.String,
			PowerSystemName:          projects[i].PowerSystemName.String,
			BoardComputingModuleName: projects[i].BoardComputingModuleName.String,
			VhfAntennaSystemName:     projects[i].VHFAntennaSystemName.String,
			VhfTransceiverName:       projects[i].VhfTransceiverName.String,
			CreatedAt:                projects[i].CreatedAt.Unix(),
			UpdatedAt:                projects[i].UpdatedAt.Time.Unix(),
		}
	}
	return projectsData
}
