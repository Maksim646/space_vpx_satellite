package handler

import (
	"context"
	"fmt"

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
	userUsecase                        model.IUserUsecase
	adminUsecase                       model.IAdminUsecase
	projectUsecase                     model.IProjectUsecase
	chassisUsecase                     model.IChassisUsecase
	cubeSatSolarPanelSideUsecase       model.ISolarPanelSideUsecase
	cubeSatSolarPanelTopUsecase        model.ISolarPanelTopUsecase
	cubeSatFrameUsecase                model.ICubeSatFrameUsecase
	cubeSatPowerSystemUsecase          model.IPowerSystemUsecase
	cubeSatBoardComputingModuleUsecase model.IBoardComputingModuleUsecase
	cubeSatVhfAntennaSystemUsecase     model.IVHFAntennaSystemUsecase
	cubeSatVhfTransceiverUsecase       model.IVHFTransceiverUsecase

	router       http.Handler
	HashSalt     string
	jwtSigninKey string
}

func New(
	userUsecase model.IUserUsecase,
	adminUsecase model.IAdminUsecase,
	projectUsecase model.IProjectUsecase,
	chassisUsecase model.IChassisUsecase,
	cubeSatSolarPanelSideUsecase model.ISolarPanelSideUsecase,
	cubeSatSolarPanelTopUsecase model.ISolarPanelTopUsecase,
	cubeSatFrameUsecase model.ICubeSatFrameUsecase,
	cubeSatPowerSystemUsecase model.IPowerSystemUsecase,
	cubeSatBoardComputingModuleUsecase model.IBoardComputingModuleUsecase,
	cubeSatVhfAntennaSystemUsecase model.IVHFAntennaSystemUsecase,
	cubeSatVhfTransceiverUsecase model.IVHFTransceiverUsecase,

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
		userUsecase:                        userUsecase,
		adminUsecase:                       adminUsecase,
		projectUsecase:                     projectUsecase,
		chassisUsecase:                     chassisUsecase,
		cubeSatSolarPanelSideUsecase:       cubeSatSolarPanelSideUsecase,
		cubeSatSolarPanelTopUsecase:        cubeSatSolarPanelTopUsecase,
		cubeSatFrameUsecase:                cubeSatFrameUsecase,
		cubeSatPowerSystemUsecase:          cubeSatPowerSystemUsecase,
		cubeSatBoardComputingModuleUsecase: cubeSatBoardComputingModuleUsecase,
		cubeSatVhfAntennaSystemUsecase:     cubeSatVhfAntennaSystemUsecase,
		cubeSatVhfTransceiverUsecase:       cubeSatVhfTransceiverUsecase,

		HashSalt:     HashSalt,
		jwtSigninKey: jwtSigninKey,
	}

	zap.L().Error("server http handler request")
	router := api.NewSpaceVPXBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof
	router.BearerAuth = h.ValidateHeader

	// LOAD TEST DATA
	router.LoadTestDataHandler = api.LoadTestDataHandlerFunc(h.LoadTestDataHandler)

	// AUTH
	router.RegisterUserHandler = api.RegisterUserHandlerFunc(h.RegisterUserHandler)
	router.LoginUserHandler = api.LoginUserHandlerFunc(h.LoginUserHandler)
	router.LoginAdminHandler = api.LoginAdminHandlerFunc(h.LoginAdminHandler)

	// CUBE SAT PROJECTS
	router.CreateCubeSatProjectHandler = api.CreateCubeSatProjectHandlerFunc(h.CreateProjectHandler)
	router.UpdateCubeSatProjectHandler = api.UpdateCubeSatProjectHandlerFunc(h.UpdateProjectHandler)
	router.GetCubeSatProjectHandler = api.GetCubeSatProjectHandlerFunc(h.GetProjectByIDHandler)
	router.DeleteCubeSatProjectHandler = api.DeleteCubeSatProjectHandlerFunc(h.DeleteProject)
	router.GetUserCubeSatProjectsHandler = api.GetUserCubeSatProjectsHandlerFunc(h.GetProjectsByUser)

	router.UpdateCubeSatFrameByProjectHandler = api.UpdateCubeSatFrameByProjectHandlerFunc(h.AddCubeSatFrameHandler)
	// router.UpdateCubeSatBoardComputingModuleByProjectHandler = api.UpdateCubeSatBoardComputingModuleByProjectHandlerFunc(h.)

	// CUBE SAT POWER SYSTEM
	router.CreatePowerSystemHandler = api.CreatePowerSystemHandlerFunc(h.CreatePowerSystemHandler)
	router.GetPowerSystemHandler = api.GetPowerSystemHandlerFunc(h.GetPowerSystemHandler)
	router.GetCubeSatPowerSystemsHandler = api.GetCubeSatPowerSystemsHandlerFunc(h.GetCubeSatPowerSystems)
	router.UpdateCubeSatPowerSystemHandler = api.UpdateCubeSatPowerSystemHandlerFunc(h.UpdateCubeSatPowerSystemHandler)
	router.DeleteCubeSatPowerSystemHandler = api.DeleteCubeSatPowerSystemHandlerFunc(h.DeleteCubeSatPowerSystemHandler)

	// CUBE SAT VHF ANTENNA SYSTEM
	router.CreateVHFAntennaSystemHandler = api.CreateVHFAntennaSystemHandlerFunc(h.CreateVHFAntennaSystemHandler)
	router.GetVHFAntennaSystemHandler = api.GetVHFAntennaSystemHandlerFunc(h.GetVHFAntennaSystemHandler)
	router.GetAvailableVHFAntennaSystemsHandler = api.GetAvailableVHFAntennaSystemsHandlerFunc(h.GetAvailableVHFAntennaSystemsHandler)
	router.UpdateVHFAntennaSystemHandler = api.UpdateVHFAntennaSystemHandlerFunc(h.UpdateVHFAntennaSystemHandler)
	router.DeleteVHFAntennaSystemHandler = api.DeleteVHFAntennaSystemHandlerFunc(h.DeleteVHFAntennaSystemHandler)

	// CUBE SAT VHF TRANSCEIVER
	router.CreateCubeSatVHFTransceiverHandler = api.CreateCubeSatVHFTransceiverHandlerFunc(h.CreateCubeSatVHFTransceiverHandler)
	router.UpdateCubeSatVHFTransceiverHandler = api.UpdateCubeSatVHFTransceiverHandlerFunc(h.UpdateCubeSatVHFTransceiverHandler)
	router.GetCubeSatVHFTransceiverHandler = api.GetCubeSatVHFTransceiverHandlerFunc(h.GetCubeSatVHFTransceiverHandler)
	router.GetAvailableCubeSatVHFTransceiversHandler = api.GetAvailableCubeSatVHFTransceiversHandlerFunc(h.GetAvailableCubeSatVHFTransceiversHandler)

	// CUBE SAT BOARD COMPUTER SYSTEM
	router.CreateBoardComputingModuleHandler = api.CreateBoardComputingModuleHandlerFunc(h.CreateBoardComputingModuleHandler)
	router.GetBoardComputingModuleByIDHandler = api.GetBoardComputingModuleByIDHandlerFunc(h.GetBoardComputingModuleHandler)
	router.GetAvailableBoardComputingModulesHandler = api.GetAvailableBoardComputingModulesHandlerFunc(h.GetBoardComputingModules)
	router.UpdateBoardComputingModuleHandler = api.UpdateBoardComputingModuleHandlerFunc(h.UpdateBoardComputingModuleHandler)
	router.DeleteBoardComputingModuleHandler = api.DeleteBoardComputingModuleHandlerFunc(h.DeleteBoardComputingModuleHandler)

	// SOLAR PANEL SIDE
	router.CreateSolarPanelSideHandler = api.CreateSolarPanelSideHandlerFunc(h.CreateSolarPanelSideHandler)
	router.GetSolarPanelSideHandler = api.GetSolarPanelSideHandlerFunc(h.GetSolarPanelSideHandler)
	router.GetCubeSatSolarPanelsSideHandler = api.GetCubeSatSolarPanelsSideHandlerFunc(h.GetCubeSatSolarPanelsSide)
	router.UpdateCubeSatSolarPanelSideHandler = api.UpdateCubeSatSolarPanelSideHandlerFunc(h.UpdateCubeSatSolarPanelSideHandler)
	router.DeleteCubeSatSolarPanelSideHandler = api.DeleteCubeSatSolarPanelSideHandlerFunc(h.DeleteCubeSatSolarPanelSideHandler)

	// SOLAR PANEL TOP
	router.CreateSolarPanelTopHandler = api.CreateSolarPanelTopHandlerFunc(h.CreateSolarPanelTopHandler)
	router.GetSolarPanelTopHandler = api.GetSolarPanelTopHandlerFunc(h.GetSolarPanelTopHandler)
	router.GetCubeSatSolarPanelsTopHandler = api.GetCubeSatSolarPanelsTopHandlerFunc(h.GetCubeSatSolarPanelsTop)
	router.UpdateCubeSatSolarPanelTopHandler = api.UpdateCubeSatSolarPanelTopHandlerFunc(h.UpdateCubeSatSolarPanelTopHandler)
	router.DeleteCubeSatSolarPanelTopHandler = api.DeleteCubeSatSolarPanelTopHandlerFunc(h.DeleteCubeSatSolarPanelTopHandler)

	// CHASSIS
	router.CreateChassisVPXHandler = api.CreateChassisVPXHandlerFunc(h.CreateChassisVPXHandler)
	router.UpdateChassisVPXHandler = api.UpdateChassisVPXHandlerFunc(h.UpdateChassisVPXHandler)
	router.GetChassisVPXByIDHandler = api.GetChassisVPXByIDHandlerFunc(h.GetChassisVPXHandler)
	router.DeleteChassisVPXHandler = api.DeleteChassisVPXHandlerFunc(h.DeleteChassisVPXHandler)
	router.GetAvailableChassisVPXHandler = api.GetAvailableChassisVPXHandlerFunc(h.GetAvailableChassisVPX)

	/// CUBE SAT FRAME
	router.CreateCubeSatFrameHandler = api.CreateCubeSatFrameHandlerFunc(h.CreateCubeSatFrameHandler)
	router.GetCubeSatFrameHandler = api.GetCubeSatFrameHandlerFunc(h.GetCubeSatFrame)
	router.GetCubeSatFramesHandler = api.GetCubeSatFramesHandlerFunc(h.GetAvailableCubeSatFrames)
	router.UpdateCubeSatFrameHandler = api.UpdateCubeSatFrameHandlerFunc(h.UpdateCubeSatFrameHandler)
	router.DeleteCubeSatFrameHandler = api.DeleteCubeSatFrameHandlerFunc(h.DeleteCubeSatFrameHandler)

	// ADD COMPONENTS
	router.UpdateCubeSatBoardComputingModuleByProjectHandler = api.UpdateCubeSatBoardComputingModuleByProjectHandlerFunc(h.AddCubeSatBoardComputingModule)

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

	fmt.Println("1")
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
