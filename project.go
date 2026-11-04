// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"net/http"
	"slices"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	shimjson "github.com/chunkifydev/chunkify-analytics-go/internal/encoding/json"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
)

// Project origins and keys.
//
// ProjectService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
	// Project origins and keys.
	Keys ProjectKeyService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	r.Keys = NewProjectKeyService(opts...)
	return
}

// Update the name, additional origins or retention of your project. Omitted fields
// stay unchanged. Set retention_seconds to null to inherit the installation
// default, or 0 to disable expiration. Changes apply to existing and new events
// without a restart. Requires a full-access private key; read-only keys
// receive 403.
func (r *ProjectService) Update(ctx context.Context, body ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/project"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Revoke all project keys. Stored events remain until their retention expires.
// Requires a full-access private key; read-only keys receive 403.
func (r *ProjectService) Delete(ctx context.Context, opts ...option.RequestOption) (res *DeletedProject, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/project"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The properties Name, Permission are required.
type CreateProjectKeyRequestParam struct {
	// Key display name, up to 128 UTF-8 bytes. Whitespace-only names are rejected.
	Name string `json:"name" api:"required"`
	// Any of "public", "secret".
	Permission CreateProjectKeyRequestPermission `json:"permission,omitzero" api:"required"`
	// Only secret keys can be read-only. False grants full project management access
	// to secret keys; public keys only ingest events.
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	paramObj
}

func (r CreateProjectKeyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProjectKeyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProjectKeyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateProjectKeyRequestPermission string

const (
	CreateProjectKeyRequestPermissionPublic CreateProjectKeyRequestPermission = "public"
	CreateProjectKeyRequestPermissionSecret CreateProjectKeyRequestPermission = "secret"
)

// A newly created key. Save the key value now; later key listings contain only
// metadata.
type CreatedProjectKey struct {
	// New key value. Save it now; it is returned once.
	Key   string `json:"key" api:"required"`
	KeyID string `json:"key_id" api:"required"`
	// Any of "public", "secret".
	Permission CreatedProjectKeyPermission `json:"permission" api:"required"`
	ProjectID  string                      `json:"project_id" api:"required"`
	// True for a private key limited to analytics reads through HTTP and MCP. Public
	// keys always report false. Rotation preserves this setting.
	ReadOnly        bool   `json:"read_only" api:"required"`
	RestartRequired bool   `json:"restart_required" api:"required"`
	Warning         string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key             respjson.Field
		KeyID           respjson.Field
		Permission      respjson.Field
		ProjectID       respjson.Field
		ReadOnly        respjson.Field
		RestartRequired respjson.Field
		Warning         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedProjectKey) RawJSON() string { return r.JSON.raw }
func (r *CreatedProjectKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatedProjectKeyPermission string

const (
	CreatedProjectKeyPermissionPublic CreatedProjectKeyPermission = "public"
	CreatedProjectKeyPermissionSecret CreatedProjectKeyPermission = "secret"
)

type DeletedProject struct {
	Deleted         bool   `json:"deleted" api:"required"`
	ProjectID       string `json:"project_id" api:"required"`
	RestartRequired bool   `json:"restart_required" api:"required"`
	Warning         string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted         respjson.Field
		ProjectID       respjson.Field
		RestartRequired respjson.Field
		Warning         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeletedProject) RawJSON() string { return r.JSON.raw }
func (r *DeletedProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Project struct {
	AllowedOrigins []string `json:"allowed_origins" api:"required"`
	// Combined shared and project origins.
	EffectiveAllowedOrigins []string `json:"effective_allowed_origins" api:"required"`
	// Effective retention in seconds after resolving the installation default. 0
	// disables expiration.
	EffectiveRetentionSeconds int64  `json:"effective_retention_seconds" api:"required"`
	Name                      string `json:"name" api:"required"`
	ProjectID                 string `json:"project_id" api:"required"`
	RestartRequired           bool   `json:"restart_required" api:"required"`
	// Retention by server receipt time, in seconds. 0 disables expiration; null
	// inherits the installation default.
	RetentionSeconds int64  `json:"retention_seconds" api:"required"`
	Warning          string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedOrigins            respjson.Field
		EffectiveAllowedOrigins   respjson.Field
		EffectiveRetentionSeconds respjson.Field
		Name                      respjson.Field
		ProjectID                 respjson.Field
		RestartRequired           respjson.Field
		RetentionSeconds          respjson.Field
		Warning                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectKey struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "public", "secret".
	Permission ProjectKeyPermission `json:"permission" api:"required"`
	// True for a private key limited to analytics reads through HTTP and MCP. Public
	// keys always report false. Rotation preserves this setting.
	ReadOnly bool `json:"read_only" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Permission  respjson.Field
		ReadOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectKey) RawJSON() string { return r.JSON.raw }
func (r *ProjectKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectKeyPermission string

const (
	ProjectKeyPermissionPublic ProjectKeyPermission = "public"
	ProjectKeyPermissionSecret ProjectKeyPermission = "secret"
)

type ProjectKeys struct {
	Keys      []ProjectKey `json:"keys" api:"required"`
	ProjectID string       `json:"project_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys        respjson.Field
		ProjectID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectKeys) RawJSON() string { return r.JSON.raw }
func (r *ProjectKeys) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParam struct {
	// Retention by server receipt time, in seconds. 0 disables expiration; null
	// inherits the installation default. Omission leaves the current setting
	// unchanged. Applies to existing and new events; removed data cannot be recovered.
	RetentionSeconds param.Opt[int64]  `json:"retention_seconds,omitzero"`
	Name             param.Opt[string] `json:"name,omitzero"`
	// Additional project origins. Shared TOML origins are always included.
	AllowedOrigins []string `json:"allowed_origins,omitzero"`
	paramObj
}

func (r ProjectUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RotatedProjectKey struct {
	// New key value. Save it now; it is returned once.
	Key   string `json:"key" api:"required"`
	KeyID string `json:"key_id" api:"required"`
	// Any of "public", "secret".
	Permission RotatedProjectKeyPermission `json:"permission" api:"required"`
	ProjectID  string                      `json:"project_id" api:"required"`
	// True for a private key limited to analytics reads through HTTP and MCP. Public
	// keys always report false. Rotation preserves this setting.
	ReadOnly        bool   `json:"read_only" api:"required"`
	RestartRequired bool   `json:"restart_required" api:"required"`
	Warning         string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key             respjson.Field
		KeyID           respjson.Field
		Permission      respjson.Field
		ProjectID       respjson.Field
		ReadOnly        respjson.Field
		RestartRequired respjson.Field
		Warning         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RotatedProjectKey) RawJSON() string { return r.JSON.raw }
func (r *RotatedProjectKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RotatedProjectKeyPermission string

const (
	RotatedProjectKeyPermissionPublic RotatedProjectKeyPermission = "public"
	RotatedProjectKeyPermissionSecret RotatedProjectKeyPermission = "secret"
)

type ProjectUpdateParams struct {
	ProjectUpdate ProjectUpdateParam
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ProjectUpdate)
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
