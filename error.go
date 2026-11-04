// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	"github.com/chunkifydev/chunkify-analytics-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/pagination"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
)

// Private retained-event and view-cohort queries.
//
// ErrorService contains methods and other services that help with interacting with
// the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewErrorService] method instead.
type ErrorService struct {
	Options []option.RequestOption
}

// NewErrorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewErrorService(opts ...option.RequestOption) (r ErrorService) {
	r = ErrorService{}
	r.Options = opts
	return
}

// Filters select first retained starts. Error events from the selected view
// identity contribute even after to. Group by exact JSON error_code and fatal
// status, distinguishing numeric codes from text and missing from false. Sort by
// affected_views descending, then canonical JSON code and fatal status ascending
// for ties. error_events counts distinct event IDs per group, affected_views
// counts distinct view IDs. Groups are not additive across codes. Cursors exclude
// later arrivals.
func (r *ErrorService) List(ctx context.Context, query ErrorListParams, opts ...option.RequestOption) (res *pagination.ErrorsPage[ErrorGroup], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/errors"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Filters select first retained starts. Error events from the selected view
// identity contribute even after to. Group by exact JSON error_code and fatal
// status, distinguishing numeric codes from text and missing from false. Sort by
// affected_views descending, then canonical JSON code and fatal status ascending
// for ties. error_events counts distinct event IDs per group, affected_views
// counts distinct view IDs. Groups are not additive across codes. Cursors exclude
// later arrivals.
func (r *ErrorService) ListAutoPaging(ctx context.Context, query ErrorListParams, opts ...option.RequestOption) *pagination.ErrorsPageAutoPager[ErrorGroup] {
	return pagination.NewErrorsPageAutoPager(r.List(ctx, query, opts...))
}

type ErrorGroup struct {
	AffectedViews int64                    `json:"affected_views" api:"required"`
	ErrorCode     ErrorGroupErrorCodeUnion `json:"error_code" api:"required"`
	ErrorEvents   int64                    `json:"error_events" api:"required"`
	Fatal         bool                     `json:"fatal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedViews respjson.Field
		ErrorCode     respjson.Field
		ErrorEvents   respjson.Field
		Fatal         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorGroup) RawJSON() string { return r.JSON.raw }
func (r *ErrorGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ErrorGroupErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type ErrorGroupErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ErrorGroupErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ErrorGroupErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ErrorGroupErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ErrorGroupErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ErrorPage struct {
	Errors     []ErrorGroup `json:"errors" api:"required"`
	NextCursor string       `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorPage) RawJSON() string { return r.JSON.raw }
func (r *ErrorPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ErrorListParams struct {
	// Exact browser name, such as Chrome. Missing values do not match.
	BrowserName param.Opt[string] `query:"browser_name,omitzero" json:"-"`
	// Exact, case-sensitive city name from the local lookup, with 1 to 256 Unicode
	// characters. Use country alongside city to distinguish places with the same name.
	// Missing locations do not match.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Exact uppercase ISO 3166-1 alpha-2 country code, such as GB. Missing locations
	// do not match.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Use next_cursor from the previous page. Keep the same filters, sort_by and
	// order. Video cursors from earlier versions are invalid; start a new pagination
	// sequence after upgrading.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Inclusive RFC3339 occurrence-time lower bound, with timezone. Summary/video
	// queries apply it to the selected view_started event. Must be before to.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Maximum rows on this page. Events and videos always support pagination; summary,
	// quality and engagement support it only with group_by.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Exact operating-system name, such as macOS. Missing values do not match.
	OsName param.Opt[string] `query:"os_name,omitzero" json:"-"`
	// Exact, case-sensitive page origin captured by the player, such as
	// https://example.com. Metrics use the first retained view start; raw events use
	// each event context. Missing values do not match.
	PageOrigin param.Opt[string] `query:"page_origin,omitzero" json:"-"`
	// Exact, case-sensitive page path captured by the player, such as /watch/video.
	// Filter by page_origin as well to distinguish identical paths on different sites.
	// Metrics use the first retained view start; raw events use each event context.
	// Missing values do not match.
	PagePath param.Opt[string] `query:"page_path,omitzero" json:"-"`
	// Exclusive RFC3339 occurrence-time upper bound, with timezone. Summary/video
	// queries apply it to the selected view_started event.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Exact video ID. Combined with all other filters using AND.
	VideoID param.Opt[string] `query:"video_id,omitzero" json:"-"`
	// Filter one playback attempt by view ID. Metrics use its first retained start.
	ViewID param.Opt[string] `query:"view_id,omitzero" json:"-"`
	// Exact supplied viewer ID. Missing/null viewer IDs do not match.
	ViewerID param.Opt[string] `query:"viewer_id,omitzero" json:"-"`
	// Exact continent code from the local lookup. Missing locations do not match.
	//
	// Any of "AF", "AN", "AS", "EU", "NA", "OC", "SA".
	Continent ErrorListParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType ErrorListParamsDeviceType `query:"device_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ErrorListParams]'s query parameters as `url.Values`.
func (r ErrorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type ErrorListParamsContinent string

const (
	ErrorListParamsContinentAf ErrorListParamsContinent = "AF"
	ErrorListParamsContinentAn ErrorListParamsContinent = "AN"
	ErrorListParamsContinentAs ErrorListParamsContinent = "AS"
	ErrorListParamsContinentEu ErrorListParamsContinent = "EU"
	ErrorListParamsContinentNa ErrorListParamsContinent = "NA"
	ErrorListParamsContinentOc ErrorListParamsContinent = "OC"
	ErrorListParamsContinentSa ErrorListParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type ErrorListParamsDeviceType string

const (
	ErrorListParamsDeviceTypeDesktop ErrorListParamsDeviceType = "desktop"
	ErrorListParamsDeviceTypeMobile  ErrorListParamsDeviceType = "mobile"
	ErrorListParamsDeviceTypeTablet  ErrorListParamsDeviceType = "tablet"
	ErrorListParamsDeviceTypeUnknown ErrorListParamsDeviceType = "unknown"
)
