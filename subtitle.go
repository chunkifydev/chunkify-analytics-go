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
// SubtitleService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubtitleService] method instead.
type SubtitleService struct {
	Options []option.RequestOption
}

// NewSubtitleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubtitleService(opts ...option.RequestOption) (r SubtitleService) {
	r = SubtitleService{}
	r.Options = opts
	return
}

// Summarize watch time per subtitle language. Expand each view's cumulative
// subtitle_playing_time array once; a null language is time with subtitles off.
// Use only the selected final event, or latest snapshot when no final exists, for
// the filtered view-start cohort. Missing or empty measurements contribute no
// time; older SDKs send none. Rank by playing time descending, then value
// ascending with null first; limit bounds returned rows only, while totals cover
// every matching row. Views can contribute to several rows and must not be summed.
// No cursor, group_by, sort_by or order. Export resource=subtitles includes all
// matching rows.
func (r *SubtitleService) Get(ctx context.Context, query SubtitleGetParams, opts ...option.RequestOption) (res *SubtitleReport, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/subtitles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Watch time with one subtitle language, or with subtitles off.
type SubtitleMetric struct {
	// Subtitle language tag; und for a track without a language; null for subtitles
	// off.
	Language string `json:"language" api:"required"`
	// Cumulative measured playing time in this language, summed once per attempt.
	PlayingTimeMs float64 `json:"playing_time_ms" api:"required"`
	// Share of all measured playing time, including rows outside the returned limit.
	PlayingTimeRatio float64 `json:"playing_time_ratio" api:"required"`
	// Distinct attempts with positive playing time in this language.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language         respjson.Field
		PlayingTimeMs    respjson.Field
		PlayingTimeRatio respjson.Field
		Views            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtitleMetric) RawJSON() string { return r.JSON.raw }
func (r *SubtitleMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtitleReport struct {
	// Most-watched first, then value ascending with null first. Defaults to 100 rows,
	// maximum 1000. CSV exports include all matching rows.
	Languages []SubtitleMetric `json:"languages" api:"required"`
	// Attempts with positive measured playing time in this report. Older SDKs send no
	// measurement.
	SampleViews int64 `json:"sample_views" api:"required"`
	// Number of distinct values before applying the row limit.
	TotalLanguages int64 `json:"total_languages" api:"required"`
	// Total measured playing time across all matching rows.
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms" api:"required"`
	// All attempts in the filtered start cohort.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Languages          respjson.Field
		SampleViews        respjson.Field
		TotalLanguages     respjson.Field
		TotalPlayingTimeMs respjson.Field
		Views              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtitleReport) RawJSON() string { return r.JSON.raw }
func (r *SubtitleReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtitleGetParams struct {
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
	Continent SubtitleGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType SubtitleGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubtitleGetParams]'s query parameters as `url.Values`.
func (r SubtitleGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type SubtitleGetParamsContinent string

const (
	SubtitleGetParamsContinentAf SubtitleGetParamsContinent = "AF"
	SubtitleGetParamsContinentAn SubtitleGetParamsContinent = "AN"
	SubtitleGetParamsContinentAs SubtitleGetParamsContinent = "AS"
	SubtitleGetParamsContinentEu SubtitleGetParamsContinent = "EU"
	SubtitleGetParamsContinentNa SubtitleGetParamsContinent = "NA"
	SubtitleGetParamsContinentOc SubtitleGetParamsContinent = "OC"
	SubtitleGetParamsContinentSa SubtitleGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type SubtitleGetParamsDeviceType string

const (
	SubtitleGetParamsDeviceTypeDesktop SubtitleGetParamsDeviceType = "desktop"
	SubtitleGetParamsDeviceTypeMobile  SubtitleGetParamsDeviceType = "mobile"
	SubtitleGetParamsDeviceTypeTablet  SubtitleGetParamsDeviceType = "tablet"
	SubtitleGetParamsDeviceTypeUnknown SubtitleGetParamsDeviceType = "unknown"
)
