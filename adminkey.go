// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
)

// Project origins and keys.
//
// AdminKeyService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminKeyService] method instead.
type AdminKeyService struct {
	Options []option.RequestOption
}

// NewAdminKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminKeyService(opts ...option.RequestOption) (r AdminKeyService) {
	r = AdminKeyService{}
	r.Options = opts
	return
}

// List key IDs, names and permissions by project ID using CHUNKIFY_ADMIN_KEY. No
// project credential is required. Key values are never returned. Intended for a
// trusted backend that verifies project ownership.
func (r *AdminKeyService) List(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectKeys, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithAdminKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/projects/%s/keys", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace one existing public or secret key using CHUNKIFY_ADMIN_KEY, even when
// the project secret was lost. The old value stops working immediately. Key ID,
// type and read-only access are preserved; other keys are unchanged. Save the
// returned replacement. Intended for a trusted backend that verifies project
// ownership.
func (r *AdminKeyService) Rotate(ctx context.Context, keyID string, body AdminKeyRotateParams, opts ...option.RequestOption) (res *RotatedProjectKey, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithAdminKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if body.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/projects/%s/keys/%s/rotate", body.ProjectID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AdminKeyRotateParams struct {
	ProjectID string `path:"project_id" api:"required" json:"-"`
	paramObj
}
