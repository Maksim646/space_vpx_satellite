package handler

import (
	"context"

	"net/http"
	"strings"

	//"strings"

	"encoding/json"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/jsonwebtoken"
	"go.uber.org/zap"

	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/go-openapi/loads"
)

type Handler struct {
	userUsecase         model.IUserUsecase
	adminUsecase        model.IAdminUsecase
	projectUsecase      model.IProjectUsecase
	chassisUsecase      model.IChassisUsecase
	solarPanelUsecase   model.ISolarPanelUsecase
	cubeSatFrameUsecase model.ICubeSatFrameUsecase

	router       http.Handler
	HashSalt     string
	jwtSigninKey string
}

func New(
	userUsecase model.IUserUsecase,
	adminUsecase model.IAdminUsecase,
	projectUsecase model.IProjectUsecase,
	chassisUsecase model.IChassisUsecase,
	solarPanelUsecase model.ISolarPanelUsecase,
	cubeSatFrameUsecase model.ICubeSatFrameUsecase,

	version string,
	HashSalt string,
	jwtSigninKey string,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		userUsecase:         userUsecase,
		adminUsecase:        adminUsecase,
		projectUsecase:      projectUsecase,
		chassisUsecase:      chassisUsecase,
		solarPanelUsecase:   solarPanelUsecase,
		cubeSatFrameUsecase: cubeSatFrameUsecase,

		HashSalt:     HashSalt,
		jwtSigninKey: jwtSigninKey,
	}

	zap.L().Error("server http handler request")
	router := api.NewSpaceVPXBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof
	router.BearerAuth = h.ValidateHeader

	// AUTH
	router.RegisterUserHandler = api.RegisterUserHandlerFunc(h.RegisterUserHandler)
	router.LoginUserHandler = api.LoginUserHandlerFunc(h.LoginUserHandler)
	router.LoginAdminHandler = api.LoginAdminHandlerFunc(h.LoginAdminHandler)

	// CUBESATPROJECTS
	router.CreateCubeSatProjectHandler = api.CreateCubeSatProjectHandlerFunc(h.CreateProjectHandler)
	router.UpdateCubeSatProjectHandler = api.UpdateCubeSatProjectHandlerFunc(h.UpdateProjectHandler)
	router.GetCubeSatProjectHandler = api.GetCubeSatProjectHandlerFunc(h.GetProjectByIDHandler)
	router.DeleteCubeSatProjectHandler = api.DeleteCubeSatProjectHandlerFunc(h.DeleteProject)
	router.GetUserCubeSatProjectsHandler = api.GetUserCubeSatProjectsHandlerFunc(h.GetProjectsByUser)

	// SOLARPANELSide
	router.CreateSolarPanelSideHandler = api.CreateSolarPanelSideHandlerFunc(h.CreateSolarPanelSideHandler)
	router.GetSolarPanelSideHandler = api.GetSolarPanelSideHandlerFunc(h.GetSolarPanelSideHandler)
	router.GetCubeSatSolarPanelsSideHandler = api.GetCubeSatSolarPanelsSideHandlerFunc(h.GetCubeSatSolarPanelsSide)
	router.UpdateCubeSatSolarPanelSideHandler = api.UpdateCubeSatSolarPanelSideHandlerFunc(h.UpdateCubeSatSolarPanelSideHandler)
	router.DeleteCubeSatSolarPanelSideHandler = api.DeleteCubeSatSolarPanelSideHandlerFunc(h.DeleteCubeSatSolarPanelSideHandler)

	// CHASSIS
	router.CreateChassisVPXHandler = api.CreateChassisVPXHandlerFunc(h.CreateChassisVPXHandler)
	router.UpdateChassisVPXHandler = api.UpdateChassisVPXHandlerFunc(h.UpdateChassisVPXHandler)
	router.GetChassisVPXByIDHandler = api.GetChassisVPXByIDHandlerFunc(h.GetChassisVPXHandler)
	router.DeleteChassisVPXHandler = api.DeleteChassisVPXHandlerFunc(h.DeleteChassisVPXHandler)
	router.GetAvailableChassisVPXHandler = api.GetAvailableChassisVPXHandlerFunc(h.GetAvailableChassisVPX)

	/// CUBESATFRAME
	router.CreateCubeSatFrameHandler = api.CreateCubeSatFrameHandlerFunc(h.CreateCubeSatFrameHandler)
	router.GetCubeSatFrameHandler = api.GetCubeSatFrameHandlerFunc(h.GetCubeSatFrame)
	router.GetCubeSatFramesHandler = api.GetCubeSatFramesHandlerFunc(h.GetAvailableCubeSatFrames)
	router.UpdateCubeSatFrameHandler = api.UpdateCubeSatFrameHandlerFunc(h.UpdateCubeSatFrameHandler)
	router.DeleteCubeSatFrameHandler = api.DeleteCubeSatFrameHandlerFunc(h.DeleteCubeSatFrameHandler)

	// USER
	router.GetUserMeHandler = api.GetUserMeHandlerFunc(h.GetUserMe)

	h.router = router.Serve(nil)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Received HTTP request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	if h.router == nil {
		zap.L().Error("h.router is nil")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	zap.L().Info("h.router is not nil, processing request")
	h.router.ServeHTTP(w, r)
}

func (h *Handler) ValidateHeader(bearerHeader string) (*definition.Principal, error) {
	ctx := context.Background()

	bearerToken := strings.TrimPrefix(bearerHeader, "Bearer ")
	userID, roleID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
	if err != nil {
		return nil, err
	}

	if roleID == 0 {
		_, err = h.userUsecase.GetUserByID(ctx, userID)
		if err != nil {
			return nil, err
		}

	} else {
		_, err = h.adminUsecase.GetAdminByID(ctx, userID)
		if err != nil {
			return nil, err
		}

	}

	return &definition.Principal{ID: userID, Role: roleID}, nil
}
