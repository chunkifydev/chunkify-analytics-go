// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	shimjson "github.com/chunkifydev/chunkify-analytics-go/internal/encoding/json"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
)

// Project origins and keys.
//
// ProjectKeyService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectKeyService] method instead.
type ProjectKeyService struct {
	Options []option.RequestOption
}

// NewProjectKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectKeyService(opts ...option.RequestOption) (r ProjectKeyService) {
	r = ProjectKeyService{}
	r.Options = opts
	return
}

// Create a named public ingestion key or secret key in the authenticated project.
// Secret keys default to full project access; set read_only to true for analytics
// reads and exports only. Requires a full-access project secret key. Read-only
// keys receive 403 and the installation admin key is not accepted. The key is
// active immediately and its value is returned once; store it before leaving this
// response. Each project supports at most 64 active keys, including its initial
// keys.
func (r *ProjectKeyService) New(ctx context.Context, body ProjectKeyNewParams, opts ...option.RequestOption) (res *CreatedProjectKey, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/project/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List key IDs, names, and types. Key values are not returned. Requires a
// full-access private key; read-only keys receive 403.
func (r *ProjectKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *ProjectKeys, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/project/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace a key in this project. The old value stops working immediately. Save the
// returned replacement. Requires a full-access private key; read-only keys
// receive 403. Only the selected key is rotated. If its replacement is lost, a
// trusted backend can rotate it again through the installation admin API.
func (r *ProjectKeyService) Rotate(ctx context.Context, keyID string, opts ...option.RequestOption) (res *RotatedProjectKey, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/project/keys/%s/rotate", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ProjectKeyNewParams struct {
	CreateProjectKeyRequest CreateProjectKeyRequestParam
	paramObj
}

func (r ProjectKeyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProjectKeyRequest)
}
func (r *ProjectKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
