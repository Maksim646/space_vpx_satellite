package handler

import (
	"fmt"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) RegisterUserHandler(req api.RegisterUserParams) middleware.Responder {
	zap.L().Info("register request")
	ctx := req.HTTPRequest.Context()

	_, err := h.userUsecase.GetUserByEmail(ctx, *req.RegisterUser.Email)
	if err == nil {
		return api.NewRegisterUserBadRequest().WithPayload(&definition.Error{
			Message: "user already exists",
		})
	}

	userID, err := h.userUsecase.CreateUser(ctx, *req.RegisterUser.Email, *req.RegisterUser.Name, *req.RegisterUser.Password)

	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return api.NewRegisterUserInternalServerError()
	}

	fmt.Println(userID)

	return api.NewRegisterUserOK()
}
