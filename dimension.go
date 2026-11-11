// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"fmt"
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
// DimensionService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDimensionService] method instead.
type DimensionService struct {
	Options []option.RequestOption
}

// NewDimensionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDimensionService(opts ...option.RequestOption) (r DimensionService) {
	r = DimensionService{}
	r.Options = opts
	return
}

// Values and view counts come from the first retained view starts, after applying
// the same cohort filters as summary. Values sort by views descending, then value
// ascending with null first for ties; missing context is separate from the literal
// unknown. Cursors bind the dimension, filters, project and ingestion watermark.
func (r *DimensionService) ListValues(ctx context.Context, dimension DimensionListValuesParamsDimension, query DimensionListValuesParams, opts ...option.RequestOption) (res *pagination.DimensionValuesPage[DimensionValuesValue], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("v1/dimensions/%v", dimension)
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

// Values and view counts come from the first retained view starts, after applying
// the same cohort filters as summary. Values sort by views descending, then value
// ascending with null first for ties; missing context is separate from the literal
// unknown. Cursors bind the dimension, filters, project and ingestion watermark.
func (r *DimensionService) ListValuesAutoPaging(ctx context.Context, dimension DimensionListValuesParamsDimension, query DimensionListValuesParams, opts ...option.RequestOption) *pagination.DimensionValuesPageAutoPager[DimensionValuesValue] {
	return pagination.NewDimensionValuesPageAutoPager(r.ListValues(ctx, dimension, query, opts...))
}

type DimensionValues struct {
	// One custom dimension's values, such as custom.plan.
	Dimension  DimensionValuesDimension `json:"dimension" api:"required"`
	NextCursor string                   `json:"next_cursor" api:"required"`
	Values     []DimensionValuesValue   `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension   respjson.Field
		NextCursor  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DimensionValues) RawJSON() string { return r.JSON.raw }
func (r *DimensionValues) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DimensionValuesDimension string

const (
	DimensionValuesDimensionBrowserName      DimensionValuesDimension = "browser_name"
	DimensionValuesDimensionOsName           DimensionValuesDimension = "os_name"
	DimensionValuesDimensionDeviceType       DimensionValuesDimension = "device_type"
	DimensionValuesDimensionVideoID          DimensionValuesDimension = "video_id"
	DimensionValuesDimensionViewerID         DimensionValuesDimension = "viewer_id"
	DimensionValuesDimensionUserID           DimensionValuesDimension = "user_id"
	DimensionValuesDimensionVideoTitle       DimensionValuesDimension = "video_title"
	DimensionValuesDimensionContinent        DimensionValuesDimension = "continent"
	DimensionValuesDimensionCountry          DimensionValuesDimension = "country"
	DimensionValuesDimensionCity             DimensionValuesDimension = "city"
	DimensionValuesDimensionPageOrigin       DimensionValuesDimension = "page_origin"
	DimensionValuesDimensionPagePath         DimensionValuesDimension = "page_path"
	DimensionValuesDimensionSubtitleLanguage DimensionValuesDimension = "subtitle_language"
	DimensionValuesDimensionPlaybackRate     DimensionValuesDimension = "playback_rate"
	DimensionValuesDimensionCustomKeys       DimensionValuesDimension = "custom_keys"
)

type DimensionValuesValue struct {
	Value string `json:"value" api:"required"`
	Views int64  `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DimensionValuesValue) RawJSON() string { return r.JSON.raw }
func (r *DimensionValuesValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DimensionListValuesParams struct {
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
	// Playback rate from 0.01 to 16 with at most two decimals, such as 1, 1.5 or 2;
	// 2.0 and 2 are the same rate. Metrics use the rate on the first retained view
	// start; raw events use each event's value.
	PlaybackRate param.Opt[string] `query:"playback_rate,omitzero" json:"-"`
	// Exact, case-sensitive, non-blank subtitle language tag, such as fr or en-US; und
	// for a subtitles track without a language. Metrics use the language showing on
	// the first retained view start; raw events use each event's value. Views without
	// subtitles at start do not match.
	SubtitleLanguage param.Opt[string] `query:"subtitle_language,omitzero" json:"-"`
	// Exclusive RFC3339 occurrence-time upper bound, with timezone. Summary/video
	// queries apply it to the selected view_started event.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Exact integrator-supplied user ID. Selects views by their start event; event
	// lists match each event's own value. Missing/null user IDs do not match. Reports
	// with this filter read one stored row per view.
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Exact video ID; must not be blank. Combined with all other filters using AND.
	VideoID param.Opt[string] `query:"video_id,omitzero" json:"-"`
	// Exact, case-sensitive video title. Selects views by the title on their first
	// retained start; event lists match each event's own value. Views without a title
	// do not match. Reports with this filter read one stored row per view.
	VideoTitle param.Opt[string] `query:"video_title,omitzero" json:"-"`
	// Filter one playback attempt by view ID. Metrics use its first retained start.
	ViewID param.Opt[string] `query:"view_id,omitzero" json:"-"`
	// Exact viewer ID. The Chunkify player SDK generates one per browser; missing/null
	// viewer IDs do not match.
	ViewerID param.Opt[string] `query:"viewer_id,omitzero" json:"-"`
	// Exact continent code from the local lookup. Missing locations do not match.
	//
	// Any of "AF", "AN", "AS", "EU", "NA", "OC", "SA".
	Continent DimensionListValuesParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType DimensionListValuesParamsDeviceType `query:"device_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DimensionListValuesParams]'s query parameters as
// `url.Values`.
func (r DimensionListValuesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DimensionListValuesParamsDimension string

const (
	DimensionListValuesParamsDimensionBrowserName      DimensionListValuesParamsDimension = "browser_name"
	DimensionListValuesParamsDimensionOsName           DimensionListValuesParamsDimension = "os_name"
	DimensionListValuesParamsDimensionDeviceType       DimensionListValuesParamsDimension = "device_type"
	DimensionListValuesParamsDimensionVideoID          DimensionListValuesParamsDimension = "video_id"
	DimensionListValuesParamsDimensionViewerID         DimensionListValuesParamsDimension = "viewer_id"
	DimensionListValuesParamsDimensionUserID           DimensionListValuesParamsDimension = "user_id"
	DimensionListValuesParamsDimensionVideoTitle       DimensionListValuesParamsDimension = "video_title"
	DimensionListValuesParamsDimensionContinent        DimensionListValuesParamsDimension = "continent"
	DimensionListValuesParamsDimensionCountry          DimensionListValuesParamsDimension = "country"
	DimensionListValuesParamsDimensionCity             DimensionListValuesParamsDimension = "city"
	DimensionListValuesParamsDimensionPageOrigin       DimensionListValuesParamsDimension = "page_origin"
	DimensionListValuesParamsDimensionPagePath         DimensionListValuesParamsDimension = "page_path"
	DimensionListValuesParamsDimensionSubtitleLanguage DimensionListValuesParamsDimension = "subtitle_language"
	DimensionListValuesParamsDimensionPlaybackRate     DimensionListValuesParamsDimension = "playback_rate"
	DimensionListValuesParamsDimensionCustomKeys       DimensionListValuesParamsDimension = "custom_keys"
)

// Exact continent code from the local lookup. Missing locations do not match.
type DimensionListValuesParamsContinent string

const (
	DimensionListValuesParamsContinentAf DimensionListValuesParamsContinent = "AF"
	DimensionListValuesParamsContinentAn DimensionListValuesParamsContinent = "AN"
	DimensionListValuesParamsContinentAs DimensionListValuesParamsContinent = "AS"
	DimensionListValuesParamsContinentEu DimensionListValuesParamsContinent = "EU"
	DimensionListValuesParamsContinentNa DimensionListValuesParamsContinent = "NA"
	DimensionListValuesParamsContinentOc DimensionListValuesParamsContinent = "OC"
	DimensionListValuesParamsContinentSa DimensionListValuesParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type DimensionListValuesParamsDeviceType string

const (
	DimensionListValuesParamsDeviceTypeDesktop DimensionListValuesParamsDeviceType = "desktop"
	DimensionListValuesParamsDeviceTypeMobile  DimensionListValuesParamsDeviceType = "mobile"
	DimensionListValuesParamsDeviceTypeTablet  DimensionListValuesParamsDeviceType = "tablet"
	DimensionListValuesParamsDeviceTypeUnknown DimensionListValuesParamsDeviceType = "unknown"
)
