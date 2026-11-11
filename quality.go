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
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
	"github.com/chunkifydev/chunkify-analytics-go/shared"
)

// Private retained-event and view-cohort queries.
//
// QualityService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualityService] method instead.
type QualityService struct {
	Options []option.RequestOption
}

// NewQualityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQualityService(opts ...option.RequestOption) (r QualityService) {
	r = QualityService{}
	r.Options = opts
	return
}

// Get startup time, buffering, and observed fatal errors for matching playback
// attempts. Null means unavailable. Optionally group the full filtered view-start
// cohort with group_by=browser_name, os_name, device_type, continent, country or
// city. Without group_by the totals response is unchanged. With group_by, return
// paginated groups sorted by views descending. Equal view counts sort by value
// ascending, null first. Context comes from the first retained start; missing,
// null and nonstring values form a null group. Keep group_by and filters unchanged
// with each cursor. Unique viewers must not be summed across groups. Includes
// mutually exclusive startup outcomes; unknown outcomes remain explicit.
func (r *QualityService) Get(ctx context.Context, query QualityGetParams, opts ...option.RequestOption) (res *QualityResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/quality"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Playback quality for the same start cohort as summary. All rates are fractions.
// Counts expose measurement coverage.
type Quality struct {
	// Duration-weighted buffering. Startup wait, pauses and seeking are excluded by
	// the SDK. Unknown durations are not treated as zero.
	Buffering QualityBuffering `json:"buffering" api:"required"`
	// Observed fatal errors. Lost events can undercount this metric; absence of an
	// error is not proof of successful playback.
	FatalErrors   QualityFatalErrors   `json:"fatal_errors" api:"required"`
	PlaybackScore shared.PlaybackScore `json:"playback_score" api:"required"`
	// Mutually exclusive outcomes. An explicit startup_outcome on the selected final
	// event takes precedence. Otherwise an available startup time, including zero, or
	// positive playing time on the selected measurement proves started. All other
	// views are unknown. sample_views excludes unknown; every ratio uses all selected
	// attempts, including unknown. Missing final delivery does not prove an exit or
	// failure.
	StartupOutcomes shared.StartupOutcomes `json:"startup_outcomes" api:"required"`
	// Time from attempt to first playing, measured once per eligible view. This is not
	// first-frame timing.
	StartupTimeMs QualityStartupTimeMs `json:"startup_time_ms" api:"required"`
	// Selected playback attempts with a retained start, including failed attempts.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffering       respjson.Field
		FatalErrors     respjson.Field
		PlaybackScore   respjson.Field
		StartupOutcomes respjson.Field
		StartupTimeMs   respjson.Field
		Views           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quality) RawJSON() string { return r.JSON.raw }
func (r *Quality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Duration-weighted buffering. Startup wait, pauses and seeking are excluded by
// the SDK. Unknown durations are not treated as zero.
type QualityBuffering struct {
	// Total rebuffer time divided by combined playing and rebuffer time. Null without
	// samples. Multiply by 100 for percent.
	Ratio float64 `json:"ratio" api:"required"`
	// Views with playing and rebuffer durations available and positive combined
	// duration.
	SampleViews int64 `json:"sample_views" api:"required"`
	// Sum of playing time for eligible buffering samples only.
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms" api:"required"`
	// Sum of rebuffer time for the same samples.
	TotalRebufferDurationMs float64 `json:"total_rebuffer_duration_ms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ratio                   respjson.Field
		SampleViews             respjson.Field
		TotalPlayingTimeMs      respjson.Field
		TotalRebufferDurationMs respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityBuffering) RawJSON() string { return r.JSON.raw }
func (r *QualityBuffering) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Observed fatal errors. Lost events can undercount this metric; absence of an
// error is not proof of successful playback.
type QualityFatalErrors struct {
	// Distinct selected views with at least one observed matching fatal
	// playback_error, including before playback starts.
	AffectedViews int64 `json:"affected_views" api:"required"`
	// Affected views divided by all selected attempts; null when views is zero.
	// Multiply by 100 for percent.
	Ratio float64 `json:"ratio" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedViews respjson.Field
		Ratio         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityFatalErrors) RawJSON() string { return r.JSON.raw }
func (r *QualityFatalErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time from attempt to first playing, measured once per eligible view. This is not
// first-frame timing.
type QualityStartupTimeMs struct {
	// Continuous median startup time in milliseconds, or null without samples.
	Median float64 `json:"median" api:"required"`
	// Continuous 95th percentile startup time in milliseconds, or null without
	// samples.
	P95 float64 `json:"p95" api:"required"`
	// Views with a non-null startup measurement on the selected final event or latest
	// snapshot.
	SampleViews int64 `json:"sample_views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Median      respjson.Field
		P95         respjson.Field
		SampleViews respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityStartupTimeMs) RawJSON() string { return r.JSON.raw }
func (r *QualityStartupTimeMs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualityBreakdownPage struct {
	// A custom dimension key, such as custom.plan.
	GroupBy    QualityBreakdownPageGroupBy `json:"group_by" api:"required"`
	Groups     []QualityBreakdownPageGroup `json:"groups" api:"required"`
	NextCursor string                      `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupBy     respjson.Field
		Groups      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityBreakdownPage) RawJSON() string { return r.JSON.raw }
func (r *QualityBreakdownPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualityBreakdownPageGroupBy string

const (
	QualityBreakdownPageGroupByBrowserName      QualityBreakdownPageGroupBy = "browser_name"
	QualityBreakdownPageGroupByOsName           QualityBreakdownPageGroupBy = "os_name"
	QualityBreakdownPageGroupByDeviceType       QualityBreakdownPageGroupBy = "device_type"
	QualityBreakdownPageGroupByContinent        QualityBreakdownPageGroupBy = "continent"
	QualityBreakdownPageGroupByCountry          QualityBreakdownPageGroupBy = "country"
	QualityBreakdownPageGroupByCity             QualityBreakdownPageGroupBy = "city"
	QualityBreakdownPageGroupByPageOrigin       QualityBreakdownPageGroupBy = "page_origin"
	QualityBreakdownPageGroupByPagePath         QualityBreakdownPageGroupBy = "page_path"
	QualityBreakdownPageGroupBySubtitleLanguage QualityBreakdownPageGroupBy = "subtitle_language"
	QualityBreakdownPageGroupByPlaybackRate     QualityBreakdownPageGroupBy = "playback_rate"
	QualityBreakdownPageGroupByVideoTitle       QualityBreakdownPageGroupBy = "video_title"
	QualityBreakdownPageGroupByUserID           QualityBreakdownPageGroupBy = "user_id"
)

type QualityBreakdownPageGroup struct {
	// Playback quality for the same start cohort as summary. All rates are fractions.
	// Counts expose measurement coverage.
	Data  Quality `json:"data" api:"required"`
	Value string  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityBreakdownPageGroup) RawJSON() string { return r.JSON.raw }
func (r *QualityBreakdownPageGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QualityResponseUnion contains all possible properties and values from [Quality],
// [QualityBreakdownPage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type QualityResponseUnion struct {
	// This field is from variant [Quality].
	Buffering QualityBuffering `json:"buffering"`
	// This field is from variant [Quality].
	FatalErrors QualityFatalErrors `json:"fatal_errors"`
	// This field is from variant [Quality].
	PlaybackScore shared.PlaybackScore `json:"playback_score"`
	// This field is from variant [Quality].
	StartupOutcomes shared.StartupOutcomes `json:"startup_outcomes"`
	// This field is from variant [Quality].
	StartupTimeMs QualityStartupTimeMs `json:"startup_time_ms"`
	// This field is from variant [Quality].
	Views int64 `json:"views"`
	// This field is from variant [QualityBreakdownPage].
	GroupBy QualityBreakdownPageGroupBy `json:"group_by"`
	// This field is from variant [QualityBreakdownPage].
	Groups []QualityBreakdownPageGroup `json:"groups"`
	// This field is from variant [QualityBreakdownPage].
	NextCursor string `json:"next_cursor"`
	JSON       struct {
		Buffering       respjson.Field
		FatalErrors     respjson.Field
		PlaybackScore   respjson.Field
		StartupOutcomes respjson.Field
		StartupTimeMs   respjson.Field
		Views           respjson.Field
		GroupBy         respjson.Field
		Groups          respjson.Field
		NextCursor      respjson.Field
		raw             string
	} `json:"-"`
}

func (u QualityResponseUnion) AsQuality() (v Quality) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QualityResponseUnion) AsQualityBreakdownPage() (v QualityBreakdownPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QualityResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *QualityResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualityGetParams struct {
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
	Continent QualityGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType QualityGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
	// Optional on summary, quality and engagement. Group views by this field on their
	// first retained start. Missing, null and nonstring context forms a null group;
	// the literal device type unknown stays separate. Groups sort by views descending,
	// then value ascending with null first for ties. Pagination is available only with
	// group_by. Location uses the first retained start too. City groups use the stored
	// name; filter by country to avoid combining names shared by different countries.
	// Page groups use the first retained start even if playback continues after
	// navigation. Filter by page_origin when grouping page_path to distinguish
	// identical paths on different sites. video_title groups by the title on each
	// view's first retained start, so a renamed video counts each view under the title
	// it started with. user_id groups by the user ID on each view's first retained
	// start. Use custom.<key>, such as custom.plan, to group by a custom dimension;
	// views without the key form the null group.
	GroupBy QualityGetParamsGroupBy `query:"group_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QualityGetParams]'s query parameters as `url.Values`.
func (r QualityGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type QualityGetParamsContinent string

const (
	QualityGetParamsContinentAf QualityGetParamsContinent = "AF"
	QualityGetParamsContinentAn QualityGetParamsContinent = "AN"
	QualityGetParamsContinentAs QualityGetParamsContinent = "AS"
	QualityGetParamsContinentEu QualityGetParamsContinent = "EU"
	QualityGetParamsContinentNa QualityGetParamsContinent = "NA"
	QualityGetParamsContinentOc QualityGetParamsContinent = "OC"
	QualityGetParamsContinentSa QualityGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type QualityGetParamsDeviceType string

const (
	QualityGetParamsDeviceTypeDesktop QualityGetParamsDeviceType = "desktop"
	QualityGetParamsDeviceTypeMobile  QualityGetParamsDeviceType = "mobile"
	QualityGetParamsDeviceTypeTablet  QualityGetParamsDeviceType = "tablet"
	QualityGetParamsDeviceTypeUnknown QualityGetParamsDeviceType = "unknown"
)

// Optional on summary, quality and engagement. Group views by this field on their
// first retained start. Missing, null and nonstring context forms a null group;
// the literal device type unknown stays separate. Groups sort by views descending,
// then value ascending with null first for ties. Pagination is available only with
// group_by. Location uses the first retained start too. City groups use the stored
// name; filter by country to avoid combining names shared by different countries.
// Page groups use the first retained start even if playback continues after
// navigation. Filter by page_origin when grouping page_path to distinguish
// identical paths on different sites. video_title groups by the title on each
// view's first retained start, so a renamed video counts each view under the title
// it started with. user_id groups by the user ID on each view's first retained
// start. Use custom.<key>, such as custom.plan, to group by a custom dimension;
// views without the key form the null group.
type QualityGetParamsGroupBy string

const (
	QualityGetParamsGroupByBrowserName      QualityGetParamsGroupBy = "browser_name"
	QualityGetParamsGroupByOsName           QualityGetParamsGroupBy = "os_name"
	QualityGetParamsGroupByDeviceType       QualityGetParamsGroupBy = "device_type"
	QualityGetParamsGroupByContinent        QualityGetParamsGroupBy = "continent"
	QualityGetParamsGroupByCountry          QualityGetParamsGroupBy = "country"
	QualityGetParamsGroupByCity             QualityGetParamsGroupBy = "city"
	QualityGetParamsGroupByPageOrigin       QualityGetParamsGroupBy = "page_origin"
	QualityGetParamsGroupByPagePath         QualityGetParamsGroupBy = "page_path"
	QualityGetParamsGroupBySubtitleLanguage QualityGetParamsGroupBy = "subtitle_language"
	QualityGetParamsGroupByPlaybackRate     QualityGetParamsGroupBy = "playback_rate"
	QualityGetParamsGroupByVideoTitle       QualityGetParamsGroupBy = "video_title"
	QualityGetParamsGroupByUserID           QualityGetParamsGroupBy = "user_id"
)
