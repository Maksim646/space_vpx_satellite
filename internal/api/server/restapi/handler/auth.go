package handler

import (
	"strings"

	"github.com/Maksim646/space_vpx_satellite/pkg/hash"
	"github.com/Maksim646/space_vpx_satellite/pkg/jsonwebtoken"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) RegisterUserHandler(req api.RegisterUserParams) middleware.Responder {
	zap.L().Info("register request")
	ctx := req.HTTPRequest.Context()

	_, err := h.userUsecase.GetUserByEmail(ctx, *req.RegisterUser.Email)
	if err == nil {
		return api.NewRegisterUserBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr("user already exists"),
		})
	}

	passwordHash, err := hash.GenerateHash(*req.RegisterUser.Password, h.HashSalt)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	userID, err := h.userUsecase.CreateUser(ctx, *req.RegisterUser.Email, *req.RegisterUser.Name, passwordHash)

	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return api.NewRegisterUserInternalServerError()
	}

	user, err := h.userUsecase.GetUserByID(ctx, userID)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	token, err := jsonwebtoken.GenerateToken(user.ID, 0, h.jwtSigninKey)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	return api.NewRegisterUserOK().WithPayload(&definition.LoginResponse{
		AccessToken: token,
	})
}

func (h *Handler) LoginUserHandler(req api.LoginUserParams) middleware.Responder {
	zap.L().Info("login request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByEmail(ctx, *req.LoginUser.Email)
	if err != nil {
		return api.NewLoginUserBadRequest().WithPayload(&definition.Error{
			Message: &model.UserNotFound,
		})
	}

	passwordHash, err := hash.GenerateHash(*req.LoginUser.Password, h.HashSalt)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	if strings.Compare(user.PasswordHash, passwordHash) != 0 {

		return api.NewLoginUserBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr("invalid login or password"),
		})
	}

	token, err := jsonwebtoken.GenerateToken(user.ID, 0, h.jwtSigninKey)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	return api.NewLoginUserOK().WithPayload(&definition.LoginResponse{
		AccessToken: token,
	})

}

func (h *Handler) LoginAdminHandler(req api.LoginAdminParams) middleware.Responder {
	zap.L().Info("login request")
	ctx := req.HTTPRequest.Context()

	admin, err := h.adminUsecase.GetAdminByEmail(ctx, *req.LoginAdmin.Email)
	if err != nil {
		return api.NewLoginUserBadRequest().WithPayload(&definition.Error{
			Message: &model.UserNotFound,
		})
	}

	passwordHash, err := hash.GenerateHash(*req.LoginAdmin.Password, h.HashSalt)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	if strings.Compare(admin.PasswordHash, passwordHash) != 0 {

		return api.NewLoginUserBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr("invalid login or password"),
		})
	}

	token, err := jsonwebtoken.GenerateToken(admin.ID, 1, h.jwtSigninKey)
	if err != nil {
		return api.NewRegisterUserInternalServerError()
	}

	return api.NewLoginUserOK().WithPayload(&definition.LoginResponse{
		AccessToken: token,
	})

}
