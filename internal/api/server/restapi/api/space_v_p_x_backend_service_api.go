// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// NewSpaceVPXBackendServiceAPI creates a new SpaceVPXBackendService instance
func NewSpaceVPXBackendServiceAPI(spec *loads.Document) *SpaceVPXBackendServiceAPI {
	return &SpaceVPXBackendServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CreateChassisVPXHandler: CreateChassisVPXHandlerFunc(func(params CreateChassisVPXParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation CreateChassisVPX has not yet been implemented")
		}),
		CreateCubeSatFrameHandler: CreateCubeSatFrameHandlerFunc(func(params CreateCubeSatFrameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation CreateCubeSatFrame has not yet been implemented")
		}),
		CreateProjectHandler: CreateProjectHandlerFunc(func(params CreateProjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation CreateProject has not yet been implemented")
		}),
		CreateSolarPanelHandler: CreateSolarPanelHandlerFunc(func(params CreateSolarPanelParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation CreateSolarPanel has not yet been implemented")
		}),
		DeleteChassisVPXHandler: DeleteChassisVPXHandlerFunc(func(params DeleteChassisVPXParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteChassisVPX has not yet been implemented")
		}),
		DeleteCubeSatFrameHandler: DeleteCubeSatFrameHandlerFunc(func(params DeleteCubeSatFrameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteCubeSatFrame has not yet been implemented")
		}),
		DeleteProjectHandler: DeleteProjectHandlerFunc(func(params DeleteProjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteProject has not yet been implemented")
		}),
		GetAvailableChassisVPXHandler: GetAvailableChassisVPXHandlerFunc(func(params GetAvailableChassisVPXParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetAvailableChassisVPX has not yet been implemented")
		}),
		GetChassisVPXByIDHandler: GetChassisVPXByIDHandlerFunc(func(params GetChassisVPXByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetChassisVPXByID has not yet been implemented")
		}),
		GetCubeSatFrameHandler: GetCubeSatFrameHandlerFunc(func(params GetCubeSatFrameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetCubeSatFrame has not yet been implemented")
		}),
		GetCubeSatFramesHandler: GetCubeSatFramesHandlerFunc(func(params GetCubeSatFramesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetCubeSatFrames has not yet been implemented")
		}),
		GetProjectHandler: GetProjectHandlerFunc(func(params GetProjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetProject has not yet been implemented")
		}),
		GetSolarPanelHandler: GetSolarPanelHandlerFunc(func(params GetSolarPanelParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetSolarPanel has not yet been implemented")
		}),
		GetUserMeHandler: GetUserMeHandlerFunc(func(params GetUserMeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetUserMe has not yet been implemented")
		}),
		GetUserProjectsHandler: GetUserProjectsHandlerFunc(func(params GetUserProjectsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetUserProjects has not yet been implemented")
		}),
		LoginAdminHandler: LoginAdminHandlerFunc(func(params LoginAdminParams) middleware.Responder {
			return middleware.NotImplemented("operation LoginAdmin has not yet been implemented")
		}),
		LoginUserHandler: LoginUserHandlerFunc(func(params LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation LoginUser has not yet been implemented")
		}),
		RegisterUserHandler: RegisterUserHandlerFunc(func(params RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation RegisterUser has not yet been implemented")
		}),
		UpdateChassisVPXHandler: UpdateChassisVPXHandlerFunc(func(params UpdateChassisVPXParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UpdateChassisVPX has not yet been implemented")
		}),
		UpdateCubeSatFrameHandler: UpdateCubeSatFrameHandlerFunc(func(params UpdateCubeSatFrameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UpdateCubeSatFrame has not yet been implemented")
		}),
		UpdateProjectHandler: UpdateProjectHandlerFunc(func(params UpdateProjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UpdateProject has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SpaceVPXBackendServiceAPI Space VPX WebSite Backend Service */
type SpaceVPXBackendServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CreateChassisVPXHandler sets the operation handler for the create chassis v p x operation
	CreateChassisVPXHandler CreateChassisVPXHandler
	// CreateCubeSatFrameHandler sets the operation handler for the create cube sat frame operation
	CreateCubeSatFrameHandler CreateCubeSatFrameHandler
	// CreateProjectHandler sets the operation handler for the create project operation
	CreateProjectHandler CreateProjectHandler
	// CreateSolarPanelHandler sets the operation handler for the create solar panel operation
	CreateSolarPanelHandler CreateSolarPanelHandler
	// DeleteChassisVPXHandler sets the operation handler for the delete chassis v p x operation
	DeleteChassisVPXHandler DeleteChassisVPXHandler
	// DeleteCubeSatFrameHandler sets the operation handler for the delete cube sat frame operation
	DeleteCubeSatFrameHandler DeleteCubeSatFrameHandler
	// DeleteProjectHandler sets the operation handler for the delete project operation
	DeleteProjectHandler DeleteProjectHandler
	// GetAvailableChassisVPXHandler sets the operation handler for the get available chassis v p x operation
	GetAvailableChassisVPXHandler GetAvailableChassisVPXHandler
	// GetChassisVPXByIDHandler sets the operation handler for the get chassis v p x by ID operation
	GetChassisVPXByIDHandler GetChassisVPXByIDHandler
	// GetCubeSatFrameHandler sets the operation handler for the get cube sat frame operation
	GetCubeSatFrameHandler GetCubeSatFrameHandler
	// GetCubeSatFramesHandler sets the operation handler for the get cube sat frames operation
	GetCubeSatFramesHandler GetCubeSatFramesHandler
	// GetProjectHandler sets the operation handler for the get project operation
	GetProjectHandler GetProjectHandler
	// GetSolarPanelHandler sets the operation handler for the get solar panel operation
	GetSolarPanelHandler GetSolarPanelHandler
	// GetUserMeHandler sets the operation handler for the get user me operation
	GetUserMeHandler GetUserMeHandler
	// GetUserProjectsHandler sets the operation handler for the get user projects operation
	GetUserProjectsHandler GetUserProjectsHandler
	// LoginAdminHandler sets the operation handler for the login admin operation
	LoginAdminHandler LoginAdminHandler
	// LoginUserHandler sets the operation handler for the login user operation
	LoginUserHandler LoginUserHandler
	// RegisterUserHandler sets the operation handler for the register user operation
	RegisterUserHandler RegisterUserHandler
	// UpdateChassisVPXHandler sets the operation handler for the update chassis v p x operation
	UpdateChassisVPXHandler UpdateChassisVPXHandler
	// UpdateCubeSatFrameHandler sets the operation handler for the update cube sat frame operation
	UpdateCubeSatFrameHandler UpdateCubeSatFrameHandler
	// UpdateProjectHandler sets the operation handler for the update project operation
	UpdateProjectHandler UpdateProjectHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SpaceVPXBackendServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SpaceVPXBackendServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SpaceVPXBackendServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SpaceVPXBackendServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SpaceVPXBackendServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SpaceVPXBackendServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SpaceVPXBackendServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SpaceVPXBackendServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SpaceVPXBackendServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SpaceVPXBackendServiceAPI
func (o *SpaceVPXBackendServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.CreateChassisVPXHandler == nil {
		unregistered = append(unregistered, "CreateChassisVPXHandler")
	}
	if o.CreateCubeSatFrameHandler == nil {
		unregistered = append(unregistered, "CreateCubeSatFrameHandler")
	}
	if o.CreateProjectHandler == nil {
		unregistered = append(unregistered, "CreateProjectHandler")
	}
	if o.CreateSolarPanelHandler == nil {
		unregistered = append(unregistered, "CreateSolarPanelHandler")
	}
	if o.DeleteChassisVPXHandler == nil {
		unregistered = append(unregistered, "DeleteChassisVPXHandler")
	}
	if o.DeleteCubeSatFrameHandler == nil {
		unregistered = append(unregistered, "DeleteCubeSatFrameHandler")
	}
	if o.DeleteProjectHandler == nil {
		unregistered = append(unregistered, "DeleteProjectHandler")
	}
	if o.GetAvailableChassisVPXHandler == nil {
		unregistered = append(unregistered, "GetAvailableChassisVPXHandler")
	}
	if o.GetChassisVPXByIDHandler == nil {
		unregistered = append(unregistered, "GetChassisVPXByIDHandler")
	}
	if o.GetCubeSatFrameHandler == nil {
		unregistered = append(unregistered, "GetCubeSatFrameHandler")
	}
	if o.GetCubeSatFramesHandler == nil {
		unregistered = append(unregistered, "GetCubeSatFramesHandler")
	}
	if o.GetProjectHandler == nil {
		unregistered = append(unregistered, "GetProjectHandler")
	}
	if o.GetSolarPanelHandler == nil {
		unregistered = append(unregistered, "GetSolarPanelHandler")
	}
	if o.GetUserMeHandler == nil {
		unregistered = append(unregistered, "GetUserMeHandler")
	}
	if o.GetUserProjectsHandler == nil {
		unregistered = append(unregistered, "GetUserProjectsHandler")
	}
	if o.LoginAdminHandler == nil {
		unregistered = append(unregistered, "LoginAdminHandler")
	}
	if o.LoginUserHandler == nil {
		unregistered = append(unregistered, "LoginUserHandler")
	}
	if o.RegisterUserHandler == nil {
		unregistered = append(unregistered, "RegisterUserHandler")
	}
	if o.UpdateChassisVPXHandler == nil {
		unregistered = append(unregistered, "UpdateChassisVPXHandler")
	}
	if o.UpdateCubeSatFrameHandler == nil {
		unregistered = append(unregistered, "UpdateCubeSatFrameHandler")
	}
	if o.UpdateProjectHandler == nil {
		unregistered = append(unregistered, "UpdateProjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SpaceVPXBackendServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SpaceVPXBackendServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *SpaceVPXBackendServiceAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SpaceVPXBackendServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SpaceVPXBackendServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SpaceVPXBackendServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the space v p x backend service API
func (o *SpaceVPXBackendServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SpaceVPXBackendServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/chassis_vpx"] = NewCreateChassisVPX(o.context, o.CreateChassisVPXHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cube_sat_frame"] = NewCreateCubeSatFrame(o.context, o.CreateCubeSatFrameHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project"] = NewCreateProject(o.context, o.CreateProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/solar_panel"] = NewCreateSolarPanel(o.context, o.CreateSolarPanelHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/chassis_vpx/{id}"] = NewDeleteChassisVPX(o.context, o.DeleteChassisVPXHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/cube_sat_frame/{id}"] = NewDeleteCubeSatFrame(o.context, o.DeleteCubeSatFrameHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/project/{id}"] = NewDeleteProject(o.context, o.DeleteProjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chassis_vpx/available_chassis"] = NewGetAvailableChassisVPX(o.context, o.GetAvailableChassisVPXHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chassis_vpx/{id}"] = NewGetChassisVPXByID(o.context, o.GetChassisVPXByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cube_sat_frame/{id}"] = NewGetCubeSatFrame(o.context, o.GetCubeSatFrameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cube_sat_frame/available_cube_sat_frame"] = NewGetCubeSatFrames(o.context, o.GetCubeSatFramesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/project/{id}"] = NewGetProject(o.context, o.GetProjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/solar_panel/{id}"] = NewGetSolarPanel(o.context, o.GetSolarPanelHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/get_me"] = NewGetUserMe(o.context, o.GetUserMeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/user_projects"] = NewGetUserProjects(o.context, o.GetUserProjectsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/admin_login"] = NewLoginAdmin(o.context, o.LoginAdminHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = NewLoginUser(o.context, o.LoginUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/register"] = NewRegisterUser(o.context, o.RegisterUserHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/chassis_vpx/{id}"] = NewUpdateChassisVPX(o.context, o.UpdateChassisVPXHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/cube_sat_frame/{id}"] = NewUpdateCubeSatFrame(o.context, o.UpdateCubeSatFrameHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/project/{id}"] = NewUpdateProject(o.context, o.UpdateProjectHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SpaceVPXBackendServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SpaceVPXBackendServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SpaceVPXBackendServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SpaceVPXBackendServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SpaceVPXBackendServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
