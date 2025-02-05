package handler

import (
	"context"

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

	projectID, err := h.projectUsecase.CreatedProject(ctx, *req.CreateCubeSatProject.ProjectName, principal.ID)
	if err != nil {
		zap.L().Error("error create project", zap.Error(err))
		return api.NewCreateCubeSatProjectInternalServerError()
	}

	return api.NewCreateCubeSatProjectOK().WithPayload(&definition.CreateCubeSatProjectResponse{
		ProjectID: useful.StrPtr(projectID),
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
		zap.L().Error("error getch cube sat project", zap.Error(err))
		return api.NewUpdateCubeSatProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	if *req.UpdateCubeSatProjectBody.ProjectName != "" {
		project.Name = *req.UpdateCubeSatProjectBody.ProjectName
	}

	if req.UpdateCubeSatProjectBody.CubeSatFrameName != "" {
		project.CubeSatFrameName.String = req.UpdateCubeSatProjectBody.CubeSatFrameName
	}

	if req.UpdateCubeSatProjectBody.SolarPanaelName != "" {
		project.CubeSatSolarPanelName.String = req.UpdateCubeSatProjectBody.SolarPanaelName
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
		ID:                    project.ID,
		UserID:                project.UserID,
		ProjectName:           &project.Name,
		CubeSatFrameName:      project.CubeSatFrameName.String,
		CubeSatSolarPanelName: project.CubeSatSolarPanelName.String,
		CreatedAt:             project.CreatedAt.Unix(),
		UpdatedAt:             project.UpdatedAt.Time.Unix(),
	}

	return *projectResult
}

func (h *Handler) ProjectsToDefinition(ctx context.Context, projects []model.CubeSatProject) []*definition.CubeSatProject {
	projectsData := make([]*definition.CubeSatProject, len(projects))

	for i := range projects {
		projectsData[i] = &definition.CubeSatProject{
			ID:                    projects[i].ID,
			ProjectName:           &projects[i].Name,
			UserID:                projects[i].UserID,
			CubeSatFrameName:      projects[i].CubeSatFrameName.String,
			CubeSatSolarPanelName: projects[i].CubeSatSolarPanelName.String,
			CreatedAt:             projects[i].CreatedAt.Unix(),
			UpdatedAt:             projects[i].UpdatedAt.Time.Unix(),
		}
	}
	return projectsData
}
