package handler

import (
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) RegisterUserHandler(req api.RegisterUserParams) middleware.Responder {
	zap.L().Info("register request")
	ctx := req.HTTPRequest.Context()

	_, err := h.userUsecase.CreateUser(ctx, *req.RegisterUser.Email, *req.RegisterUser.Name, *req.RegisterUser.Password)

	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}
	return api.NewRegisterUserOK()
}
