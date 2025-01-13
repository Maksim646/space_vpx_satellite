package handler

import (
	//"context"

	"net/http"
	"strings"

	//"strings"

	"encoding/json"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"go.uber.org/zap"

	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/go-openapi/loads"
)

type Handler struct {
	router      http.Handler
	userUsecase model.IUserUsecase
}

func New(
	userUsecase model.IUserUsecase,
	version string,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		userUsecase: userUsecase,
	}

	router := api.NewSpaceVPXBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof
	// router.BearerAuthenticator = h.ValidateHeader

	// CREATE USER
	router.RegisterUserHandler = api.RegisterUserHandlerFunc(h.RegisterUserHandler)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// func (h *Handler) ValidateHeader(bearerHeader string) (*definition.Principal, error) {
// 	ctx := context.Background()

// 	bearerToken := strings.TrimPrefix(bearerHeader, "Bearer ")
// 	userID, roleID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if roleID == 0 {
// 		_, err = h.userUsecase.GetUserByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if err := h.userUsecase.UpdateLastVisitTime(ctx, userID); err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		_, err = h.adminUsecase.GetByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if err := h.adminUsecase.UpdateLastVisitTime(ctx, userID); err != nil {
// 			return nil, err
// 		}

// 	}

// 	return &definition.Principal{ID: userID}, nil
// }
