package handler

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) GetUserMe(req api.GetUserMeParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("get user me request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByID(ctx, principal.ID)
	if err != nil {
		return api.NewGetUserMeBadRequest().WithPayload(&definition.Error{
			Message: &model.UserNotFound,
		})
	}

	userResult := h.UserToDefinition(ctx, user)

	return api.NewGetUserMeOK().WithPayload(&userResult)
}

func (h *Handler) UserToDefinition(ctx context.Context, user model.User) definition.User {
	resultUser := definition.User{
		ID:        user.ID,
		Email:     &user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Time.Unix(),
	}

	return resultUser
}
