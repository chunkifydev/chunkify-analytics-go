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
// AdminService contains methods and other services that help with interacting with
// the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	Options []option.RequestOption
	// Project origins and keys.
	Keys AdminKeyService
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r AdminService) {
	r = AdminService{}
	r.Options = opts
	r.Keys = NewAdminKeyService(opts...)
	return
}

// Generate a project ID and initial keys, optionally setting retention. Requires
// the installation admin key. The project and its retention are active before
// success is returned.
func (r *AdminService) NewProject(ctx context.Context, body AdminNewProjectParams, opts ...option.RequestOption) (res *CreatedProject, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithAdminKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The property Name is required.
type CreateProjectRequestParam struct {
	// Project display name.
	Name string `json:"name" api:"required"`
	// Retention by server receipt time, in seconds. 0 disables expiration; null
	// inherits the installation default. Omission also inherits the default.
	RetentionSeconds param.Opt[int64] `json:"retention_seconds,omitzero"`
	// Additional project origins. Shared TOML origins are always included.
	AllowedOrigins []string `json:"allowed_origins,omitzero"`
	paramObj
}

func (r CreateProjectRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProjectRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProjectRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatedProject struct {
	// Effective retention in seconds after resolving the installation default. 0
	// disables expiration.
	EffectiveRetentionSeconds int64 `json:"effective_retention_seconds" api:"required"`
	// Project display name.
	Name      string `json:"name" api:"required"`
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Public key for sending events.
	PublicKey       string `json:"public_key" api:"required"`
	RestartRequired bool   `json:"restart_required" api:"required"`
	// Retention by server receipt time, in seconds. 0 disables expiration; null
	// inherits the installation default.
	RetentionSeconds int64 `json:"retention_seconds" api:"required"`
	// Secret key for this project. Save it now; it is returned once.
	SecretKey string `json:"secret_key" api:"required"`
	Warning   string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffectiveRetentionSeconds respjson.Field
		Name                      respjson.Field
		ProjectID                 respjson.Field
		PublicKey                 respjson.Field
		RestartRequired           respjson.Field
		RetentionSeconds          respjson.Field
		SecretKey                 respjson.Field
		Warning                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedProject) RawJSON() string { return r.JSON.raw }
func (r *CreatedProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminNewProjectParams struct {
	CreateProjectRequest CreateProjectRequestParam
	paramObj
}

func (r AdminNewProjectParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProjectRequest)
}
func (r *AdminNewProjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
