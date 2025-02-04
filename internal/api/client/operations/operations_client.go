// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateChassisVPX(params *CreateChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChassisVPXOK, error)

	CreateCubeSatFrame(params *CreateCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCubeSatFrameOK, error)

	CreateProject(params *CreateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProjectOK, error)

	CreateSolarPanel(params *CreateSolarPanelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSolarPanelOK, error)

	DeleteChassisVPX(params *DeleteChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteChassisVPXOK, error)

	DeleteProject(params *DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectOK, error)

	GetAvailableChassisVPX(params *GetAvailableChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableChassisVPXOK, error)

	GetChassisVPXByID(params *GetChassisVPXByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChassisVPXByIDOK, error)

	GetCubeSatFrame(params *GetCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCubeSatFrameOK, error)

	GetProject(params *GetProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectOK, error)

	GetSolarPanel(params *GetSolarPanelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSolarPanelOK, error)

	GetUserMe(params *GetUserMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserMeOK, error)

	GetUserProjects(params *GetUserProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserProjectsOK, error)

	LoginAdmin(params *LoginAdminParams, opts ...ClientOption) (*LoginAdminOK, error)

	LoginUser(params *LoginUserParams, opts ...ClientOption) (*LoginUserOK, error)

	RegisterUser(params *RegisterUserParams, opts ...ClientOption) (*RegisterUserOK, error)

	UpdateChassisVPX(params *UpdateChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateChassisVPXOK, error)

	UpdateCubeSatFrame(params *UpdateCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCubeSatFrameOK, error)

	UpdateProject(params *UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateChassisVPX creates chasis
*/
func (a *Client) CreateChassisVPX(params *CreateChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChassisVPXOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateChassisVPXParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateChassisVPX",
		Method:             "POST",
		PathPattern:        "/chassis_vpx",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChassisVPXReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateChassisVPXOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateChassisVPX: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateCubeSatFrame creates cube sat frame
*/
func (a *Client) CreateCubeSatFrame(params *CreateCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCubeSatFrameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCubeSatFrameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCubeSatFrame",
		Method:             "POST",
		PathPattern:        "/cube_sat_frame",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCubeSatFrameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCubeSatFrameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCubeSatFrame: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateProject creates project
*/
func (a *Client) CreateProject(params *CreateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateProject",
		Method:             "POST",
		PathPattern:        "/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSolarPanel creates solar panel
*/
func (a *Client) CreateSolarPanel(params *CreateSolarPanelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSolarPanelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSolarPanelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSolarPanel",
		Method:             "POST",
		PathPattern:        "/solar_panel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSolarPanelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSolarPanelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSolarPanel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteChassisVPX deletes chassis v p x by ID
*/
func (a *Client) DeleteChassisVPX(params *DeleteChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteChassisVPXOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChassisVPXParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteChassisVPX",
		Method:             "DELETE",
		PathPattern:        "/chassis_vpx/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteChassisVPXReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteChassisVPXOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteChassisVPX: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteProject deletes project
*/
func (a *Client) DeleteProject(params *DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteProject",
		Method:             "DELETE",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAvailableChassisVPX gets available chassis v p x
*/
func (a *Client) GetAvailableChassisVPX(params *GetAvailableChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableChassisVPXOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableChassisVPXParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAvailableChassisVPX",
		Method:             "GET",
		PathPattern:        "/chassis_vpx/available_chassis",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableChassisVPXReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAvailableChassisVPXOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAvailableChassisVPX: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetChassisVPXByID gets chassis v p x by ID
*/
func (a *Client) GetChassisVPXByID(params *GetChassisVPXByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChassisVPXByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChassisVPXByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetChassisVPXByID",
		Method:             "GET",
		PathPattern:        "/chassis_vpx/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetChassisVPXByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetChassisVPXByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetChassisVPXByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCubeSatFrame gets cube sat frame
*/
func (a *Client) GetCubeSatFrame(params *GetCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCubeSatFrameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCubeSatFrameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCubeSatFrame",
		Method:             "GET",
		PathPattern:        "/cube_sat_frame/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCubeSatFrameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCubeSatFrameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCubeSatFrame: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProject gets project
*/
func (a *Client) GetProject(params *GetProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProject",
		Method:             "GET",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSolarPanel gets solar panel
*/
func (a *Client) GetSolarPanel(params *GetSolarPanelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSolarPanelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSolarPanelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSolarPanel",
		Method:             "GET",
		PathPattern:        "/solar_panel/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSolarPanelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSolarPanelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSolarPanel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserMe gets user me
*/
func (a *Client) GetUserMe(params *GetUserMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserMeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserMe",
		Method:             "GET",
		PathPattern:        "/user/get_me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUserMe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserProjects gets user projects
*/
func (a *Client) GetUserProjects(params *GetUserProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserProjects",
		Method:             "GET",
		PathPattern:        "/projects/user_projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUserProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LoginAdmin logins admin
*/
func (a *Client) LoginAdmin(params *LoginAdminParams, opts ...ClientOption) (*LoginAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginAdminParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LoginAdmin",
		Method:             "POST",
		PathPattern:        "/auth/admin_login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginAdminReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginAdminOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LoginAdmin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LoginUser logins user
*/
func (a *Client) LoginUser(params *LoginUserParams, opts ...ClientOption) (*LoginUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LoginUser",
		Method:             "POST",
		PathPattern:        "/auth/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LoginUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RegisterUser registers user
*/
func (a *Client) RegisterUser(params *RegisterUserParams, opts ...ClientOption) (*RegisterUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegisterUser",
		Method:             "POST",
		PathPattern:        "/auth/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegisterUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateChassisVPX updates chassis v p x
*/
func (a *Client) UpdateChassisVPX(params *UpdateChassisVPXParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateChassisVPXOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateChassisVPXParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateChassisVPX",
		Method:             "PATCH",
		PathPattern:        "/chassis_vpx/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateChassisVPXReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateChassisVPXOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateChassisVPX: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCubeSatFrame updates cube sat frame
*/
func (a *Client) UpdateCubeSatFrame(params *UpdateCubeSatFrameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCubeSatFrameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCubeSatFrameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCubeSatFrame",
		Method:             "PATCH",
		PathPattern:        "/cube_sat_frame/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCubeSatFrameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCubeSatFrameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCubeSatFrame: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateProject updates project
*/
func (a *Client) UpdateProject(params *UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProject",
		Method:             "PATCH",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
