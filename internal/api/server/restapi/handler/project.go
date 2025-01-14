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

func (h *Handler) CreateProjectHandler(req api.CreateProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("create project request")
	ctx := req.HTTPRequest.Context()

	projectID, err := h.projectUsecase.CreatedProject(ctx, *req.CreateProject.ProjectName, principal.ID)
	if err != nil {
		zap.L().Error("error create project", zap.Error(err))
		return api.NewCreateProjectInternalServerError()
	}

	return api.NewCreateProjectOK().WithPayload(&definition.CreateProjectResponse{
		ProjectID: useful.StrPtr(projectID),
	})
}

func (h *Handler) GetProjectByIDHandler(req api.GetProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get project request, userID: " + req.ID)
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewGetProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	projectResult := h.ProjectToDefinition(ctx, project)

	return api.NewGetProjectOK().WithPayload(&projectResult)

}

func (h *Handler) UpdateProjectHandler(req api.UpdateProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("update project request")
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	if *req.UpdateProjectBody.ProjectName != "" {
		err := h.projectUsecase.UpdateProjectByID(ctx, project.ID, *req.UpdateProjectBody.ProjectName)
		if err != nil {
			return api.NewUpdateProjectInternalServerError()
		}
	}

	newProject, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateProjectInternalServerError()
	}

	projectResult := h.ProjectToDefinition(ctx, newProject)

	return api.NewUpdateProjectOK().WithPayload(&projectResult)

}

func (h *Handler) DeleteProject(req api.DeleteProjectParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("delete project request")
	ctx := req.HTTPRequest.Context()

	project, err := h.projectUsecase.GetProjectByID(ctx, req.ID)
	if err != nil {
		return api.NewDeleteProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectNotFound,
		})
	}

	err = h.projectUsecase.DeleteProject(ctx, project.ID)
	if err != nil {
		return api.NewDeleteProjectInternalServerError()
	}

	return api.NewDeleteProjectOK().WithPayload(&definition.CreateProjectResponse{
		ProjectID: useful.StrPtr(project.ID),
	})
}

func (h *Handler) GetProjectsByUser(req api.GetUserProjectsParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get user projects request")
	ctx := req.HTTPRequest.Context()

	filters := make(map[string]interface{})
	filters[model.FilterProjectsByUser] = principal.ID

	projects, projectsCount, err := h.projectUsecase.GetProjectsByFilters(ctx, req.Offset, req.Limit, *req.SortField, filters)
	if err != nil {
		return api.NewDeleteProjectBadRequest().WithPayload(&definition.Error{
			Message: &model.ProjectsNotFound,
		})
	}

	projectsResult := h.ProjectsToDefinition(ctx, projects)

	projectsResponse := definition.Projects{
		Count:    useful.Int64Ptr(projectsCount),
		Projects: projectsResult,
	}

	return api.NewGetUserProjectsOK().WithPayload(&projectsResponse)
}

func (h *Handler) ProjectToDefinition(ctx context.Context, project model.Project) definition.Project {
	projectResult := &definition.Project{
		ID:          project.ID,
		UserID:      project.UserID,
		ProjectName: &project.Name,
		CreatedAt:   project.CreatedAt.Unix(),
		UpdatedAt:   project.UpdatedAt.Time.Unix(),
	}

	return *projectResult
}

func (h *Handler) ProjectsToDefinition(ctx context.Context, projects []model.Project) []*definition.Project {
	projectsData := make([]*definition.Project, len(projects))

	for i, _ := range projects {
		projectsData[i] = &definition.Project{
			ID:          projects[i].ID,
			ProjectName: &projects[i].Name,
			UserID:      projects[i].UserID,
			CreatedAt:   projects[i].CreatedAt.Unix(),
			UpdatedAt:   projects[i].UpdatedAt.Time.Unix(),
		}
	}
	return projectsData
}
