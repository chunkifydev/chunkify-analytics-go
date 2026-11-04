// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	"github.com/chunkifydev/chunkify-analytics-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
)

// Private retained-event and view-cohort queries.
//
// RenditionService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRenditionService] method instead.
type RenditionService struct {
	Options []option.RequestOption
}

// NewRenditionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRenditionService(opts ...option.RequestOption) (r RenditionService) {
	r = RenditionService{}
	r.Options = opts
	return
}

// Summarize rendition measurements for the filtered view-start cohort. Use only
// the selected final event, or latest snapshot when no final exists. Expand its
// cumulative rendition_playing_time array once per view. Missing or empty
// measurements contribute no rendition time. Rank by playing time descending;
// limit bounds returned rows only, while totals and average cover every matching
// rendition. Average bitrate is sum(bitrate_bps \* playing_time_ms) divided by
// playing time with a known bitrate. Dimensions, codec, bitrate and frame rate
// together identify a rendition. Views can contribute to several renditions and
// must not be summed. No cursor, group_by, sort_by or order. Export
// resource=renditions includes all matching rows.
func (r *RenditionService) Get(ctx context.Context, query RenditionGetParams, opts ...option.RequestOption) (res *RenditionReport, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/renditions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// One rendition identified by its dimensions, codec, declared bitrate and frame
// rate.
type RenditionMetric struct {
	// Declared rendition bitrate in bits per second, not measured network throughput.
	BitrateBps float64 `json:"bitrate_bps" api:"required"`
	// Reported codec; null when unknown.
	Codec string `json:"codec" api:"required"`
	// Reported frames per second; null when unknown.
	FrameRate float64 `json:"frame_rate" api:"required"`
	// Reported height in pixels; null when unknown.
	Height float64 `json:"height" api:"required"`
	// Cumulative measured playing time for this rendition, summed once per attempt.
	PlayingTimeMs float64 `json:"playing_time_ms" api:"required"`
	// Share of all measured rendition playing time, including rows outside the
	// returned limit.
	PlayingTimeRatio float64 `json:"playing_time_ratio" api:"required"`
	// Distinct attempts with positive playing time for this rendition.
	Views int64 `json:"views" api:"required"`
	// Reported width in pixels; null when unknown.
	Width float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BitrateBps       respjson.Field
		Codec            respjson.Field
		FrameRate        respjson.Field
		Height           respjson.Field
		PlayingTimeMs    respjson.Field
		PlayingTimeRatio respjson.Field
		Views            respjson.Field
		Width            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenditionMetric) RawJSON() string { return r.JSON.raw }
func (r *RenditionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RenditionReport struct {
	// Time-weighted mean declared bitrate. Null without playing time with a known
	// bitrate.
	AverageBitrateBps float64 `json:"average_bitrate_bps" api:"required"`
	// Measured playing time with a known rendition bitrate, used as the average
	// denominator.
	BitrateSamplePlayingTimeMs float64 `json:"bitrate_sample_playing_time_ms" api:"required"`
	// Most-played renditions first. Defaults to 100 rows, maximum 1000. CSV exports
	// include all matching rows.
	Renditions []RenditionMetric `json:"renditions" api:"required"`
	// Attempts with positive measured rendition playing time.
	SampleViews int64 `json:"sample_views" api:"required"`
	// Total measured rendition playing time across all matching rows. Can be less than
	// overall watch time when rendition information is missing.
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms" api:"required"`
	// Number of distinct renditions before applying the row limit.
	TotalRenditions int64 `json:"total_renditions" api:"required"`
	// All attempts in the filtered start cohort.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageBitrateBps          respjson.Field
		BitrateSamplePlayingTimeMs respjson.Field
		Renditions                 respjson.Field
		SampleViews                respjson.Field
		TotalPlayingTimeMs         respjson.Field
		TotalRenditions            respjson.Field
		Views                      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenditionReport) RawJSON() string { return r.JSON.raw }
func (r *RenditionReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RenditionGetParams struct {
	// Exact browser name, such as Chrome. Missing values do not match.
	BrowserName param.Opt[string] `query:"browser_name,omitzero" json:"-"`
	// Exact, case-sensitive city name from the local lookup, with 1 to 256 Unicode
	// characters. Use country alongside city to distinguish places with the same name.
	// Missing locations do not match.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Exact uppercase ISO 3166-1 alpha-2 country code, such as GB. Missing locations
	// do not match.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
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
	Continent RenditionGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType RenditionGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RenditionGetParams]'s query parameters as `url.Values`.
func (r RenditionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type RenditionGetParamsContinent string

const (
	RenditionGetParamsContinentAf RenditionGetParamsContinent = "AF"
	RenditionGetParamsContinentAn RenditionGetParamsContinent = "AN"
	RenditionGetParamsContinentAs RenditionGetParamsContinent = "AS"
	RenditionGetParamsContinentEu RenditionGetParamsContinent = "EU"
	RenditionGetParamsContinentNa RenditionGetParamsContinent = "NA"
	RenditionGetParamsContinentOc RenditionGetParamsContinent = "OC"
	RenditionGetParamsContinentSa RenditionGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type RenditionGetParamsDeviceType string

const (
	RenditionGetParamsDeviceTypeDesktop RenditionGetParamsDeviceType = "desktop"
	RenditionGetParamsDeviceTypeMobile  RenditionGetParamsDeviceType = "mobile"
	RenditionGetParamsDeviceTypeTablet  RenditionGetParamsDeviceType = "tablet"
	RenditionGetParamsDeviceTypeUnknown RenditionGetParamsDeviceType = "unknown"
)
