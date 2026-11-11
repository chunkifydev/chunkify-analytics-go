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
// PlaybackRateService contains methods and other services that help with
// interacting with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaybackRateService] method instead.
type PlaybackRateService struct {
	Options []option.RequestOption
}

// NewPlaybackRateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaybackRateService(opts ...option.RequestOption) (r PlaybackRateService) {
	r = PlaybackRateService{}
	r.Options = opts
	return
}

// Summarize watch time per playback rate. Expand each view's cumulative
// playback_rate_playing_time array once; rates are rounded to two decimals.
// Playing time is wall-clock time, so a minute at 2 covers about two minutes of
// media. Use only the selected final event, or latest snapshot when no final
// exists, for the filtered view-start cohort. Missing or empty measurements
// contribute no time; older SDKs send none. Rank by playing time descending, then
// value ascending with null first; limit bounds returned rows only, while totals
// cover every matching row. Views can contribute to several rows and must not be
// summed. No cursor, group_by, sort_by or order. Export resource=playback_rates
// includes all matching rows.
func (r *PlaybackRateService) Get(ctx context.Context, query PlaybackRateGetParams, opts ...option.RequestOption) (res *PlaybackRateReport, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/playback-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Watch time at one playback rate.
type PlaybackRateMetric struct {
	// Cumulative measured playing time at this rate, summed once per attempt.
	PlayingTimeMs float64 `json:"playing_time_ms" api:"required"`
	// Share of all measured playing time, including rows outside the returned limit.
	PlayingTimeRatio float64 `json:"playing_time_ratio" api:"required"`
	// Playback rate, rounded to two decimals.
	Rate float64 `json:"rate" api:"required"`
	// Distinct attempts with positive playing time at this rate.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlayingTimeMs    respjson.Field
		PlayingTimeRatio respjson.Field
		Rate             respjson.Field
		Views            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackRateMetric) RawJSON() string { return r.JSON.raw }
func (r *PlaybackRateMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackRateReport struct {
	// Most-watched first, then value ascending with null first. Defaults to 100 rows,
	// maximum 1000. CSV exports include all matching rows.
	Rates []PlaybackRateMetric `json:"rates" api:"required"`
	// Attempts with positive measured playing time in this report. Older SDKs send no
	// measurement.
	SampleViews int64 `json:"sample_views" api:"required"`
	// Total measured playing time across all matching rows.
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms" api:"required"`
	// Number of distinct values before applying the row limit.
	TotalRates int64 `json:"total_rates" api:"required"`
	// All attempts in the filtered start cohort.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rates              respjson.Field
		SampleViews        respjson.Field
		TotalPlayingTimeMs respjson.Field
		TotalRates         respjson.Field
		Views              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackRateReport) RawJSON() string { return r.JSON.raw }
func (r *PlaybackRateReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackRateGetParams struct {
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
	Continent PlaybackRateGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType PlaybackRateGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlaybackRateGetParams]'s query parameters as `url.Values`.
func (r PlaybackRateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type PlaybackRateGetParamsContinent string

const (
	PlaybackRateGetParamsContinentAf PlaybackRateGetParamsContinent = "AF"
	PlaybackRateGetParamsContinentAn PlaybackRateGetParamsContinent = "AN"
	PlaybackRateGetParamsContinentAs PlaybackRateGetParamsContinent = "AS"
	PlaybackRateGetParamsContinentEu PlaybackRateGetParamsContinent = "EU"
	PlaybackRateGetParamsContinentNa PlaybackRateGetParamsContinent = "NA"
	PlaybackRateGetParamsContinentOc PlaybackRateGetParamsContinent = "OC"
	PlaybackRateGetParamsContinentSa PlaybackRateGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type PlaybackRateGetParamsDeviceType string

const (
	PlaybackRateGetParamsDeviceTypeDesktop PlaybackRateGetParamsDeviceType = "desktop"
	PlaybackRateGetParamsDeviceTypeMobile  PlaybackRateGetParamsDeviceType = "mobile"
	PlaybackRateGetParamsDeviceTypeTablet  PlaybackRateGetParamsDeviceType = "tablet"
	PlaybackRateGetParamsDeviceTypeUnknown PlaybackRateGetParamsDeviceType = "unknown"
)
