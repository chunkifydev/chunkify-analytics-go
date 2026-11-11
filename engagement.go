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
// EngagementService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEngagementService] method instead.
type EngagementService struct {
	Options []option.RequestOption
}

// NewEngagementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEngagementService(opts ...option.RequestOption) (r EngagementService) {
	r = EngagementService{}
	r.Options = opts
	return
}

// Get average watch time per attempt and the observed reached-end rate. Reaching
// the end does not prove the whole video was watched. Optionally group the full
// filtered view-start cohort with group_by=browser_name, os_name, device_type,
// continent, country or city. Without group_by the totals response is unchanged.
// With group_by, return paginated groups sorted by views descending. Equal view
// counts sort by value ascending, null first. Context comes from the first
// retained start; missing, null and nonstring values form a null group. Keep
// group_by and filters unchanged with each cursor. Unique viewers must not be
// summed across groups. Includes furthest-progress milestones and measured
// watch-time distribution. Progress can include seeking and is not retention.
func (r *EngagementService) Get(ctx context.Context, query EngagementGetParams, opts ...option.RequestOption) (res *EngagementResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/engagement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Observed watch time and end-of-media for the same start cohort as summary.
// Missing delivery can undercount both.
type Engagement struct {
	// Total playing time divided by all selected attempts. Null with no views; zero
	// when views exist with no measured time.
	AveragePlayingTimeMs float64 `json:"average_playing_time_ms" api:"required"`
	// Furthest progress on the selected final event or latest snapshot. Only views
	// with a positive known media duration and a progress ratio in [0,1] are samples.
	// Each milestone ratio uses sample_views. Seeking can count; this is not watched
	// coverage or audience retention.
	ProgressReached shared.ProgressReached `json:"progress_reached" api:"required"`
	// Reached-end views divided by all selected attempts. Null without views. Multiply
	// by 100 for percent. Does not imply the whole video was watched.
	ReachedEndRatio float64 `json:"reached_end_ratio" api:"required"`
	// Views whose winning final event has reason=ended. Counts each selected view at
	// most once.
	ReachedEndViews int64 `json:"reached_end_views" api:"required"`
	// Cumulative playing time summed once per selected view, as in summary.
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms" api:"required"`
	// Selected playback attempts with a retained start, including failed, ongoing and
	// start-only attempts.
	Views int64 `json:"views" api:"required"`
	// Playing time in milliseconds from one selected final event or latest snapshot
	// per view. Measured zero is included; missing measurements are excluded.
	// Percentiles use linear interpolation and are null without samples. This sample
	// denominator differs from average_playing_time_ms, which uses all attempts.
	WatchTimeMs shared.WatchTimeDistribution `json:"watch_time_ms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AveragePlayingTimeMs respjson.Field
		ProgressReached      respjson.Field
		ReachedEndRatio      respjson.Field
		ReachedEndViews      respjson.Field
		TotalPlayingTimeMs   respjson.Field
		Views                respjson.Field
		WatchTimeMs          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Engagement) RawJSON() string { return r.JSON.raw }
func (r *Engagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EngagementBreakdownPage struct {
	// A custom dimension key, such as custom.plan.
	GroupBy    EngagementBreakdownPageGroupBy `json:"group_by" api:"required"`
	Groups     []EngagementBreakdownPageGroup `json:"groups" api:"required"`
	NextCursor string                         `json:"next_cursor" api:"required"`
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
func (r EngagementBreakdownPage) RawJSON() string { return r.JSON.raw }
func (r *EngagementBreakdownPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EngagementBreakdownPageGroupBy string

const (
	EngagementBreakdownPageGroupByBrowserName      EngagementBreakdownPageGroupBy = "browser_name"
	EngagementBreakdownPageGroupByOsName           EngagementBreakdownPageGroupBy = "os_name"
	EngagementBreakdownPageGroupByDeviceType       EngagementBreakdownPageGroupBy = "device_type"
	EngagementBreakdownPageGroupByContinent        EngagementBreakdownPageGroupBy = "continent"
	EngagementBreakdownPageGroupByCountry          EngagementBreakdownPageGroupBy = "country"
	EngagementBreakdownPageGroupByCity             EngagementBreakdownPageGroupBy = "city"
	EngagementBreakdownPageGroupByPageOrigin       EngagementBreakdownPageGroupBy = "page_origin"
	EngagementBreakdownPageGroupByPagePath         EngagementBreakdownPageGroupBy = "page_path"
	EngagementBreakdownPageGroupBySubtitleLanguage EngagementBreakdownPageGroupBy = "subtitle_language"
	EngagementBreakdownPageGroupByPlaybackRate     EngagementBreakdownPageGroupBy = "playback_rate"
	EngagementBreakdownPageGroupByVideoTitle       EngagementBreakdownPageGroupBy = "video_title"
	EngagementBreakdownPageGroupByUserID           EngagementBreakdownPageGroupBy = "user_id"
)

type EngagementBreakdownPageGroup struct {
	// Observed watch time and end-of-media for the same start cohort as summary.
	// Missing delivery can undercount both.
	Data  Engagement `json:"data" api:"required"`
	Value string     `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EngagementBreakdownPageGroup) RawJSON() string { return r.JSON.raw }
func (r *EngagementBreakdownPageGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EngagementResponseUnion contains all possible properties and values from
// [Engagement], [EngagementBreakdownPage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EngagementResponseUnion struct {
	// This field is from variant [Engagement].
	AveragePlayingTimeMs float64 `json:"average_playing_time_ms"`
	// This field is from variant [Engagement].
	ProgressReached shared.ProgressReached `json:"progress_reached"`
	// This field is from variant [Engagement].
	ReachedEndRatio float64 `json:"reached_end_ratio"`
	// This field is from variant [Engagement].
	ReachedEndViews int64 `json:"reached_end_views"`
	// This field is from variant [Engagement].
	TotalPlayingTimeMs float64 `json:"total_playing_time_ms"`
	// This field is from variant [Engagement].
	Views int64 `json:"views"`
	// This field is from variant [Engagement].
	WatchTimeMs shared.WatchTimeDistribution `json:"watch_time_ms"`
	// This field is from variant [EngagementBreakdownPage].
	GroupBy EngagementBreakdownPageGroupBy `json:"group_by"`
	// This field is from variant [EngagementBreakdownPage].
	Groups []EngagementBreakdownPageGroup `json:"groups"`
	// This field is from variant [EngagementBreakdownPage].
	NextCursor string `json:"next_cursor"`
	JSON       struct {
		AveragePlayingTimeMs respjson.Field
		ProgressReached      respjson.Field
		ReachedEndRatio      respjson.Field
		ReachedEndViews      respjson.Field
		TotalPlayingTimeMs   respjson.Field
		Views                respjson.Field
		WatchTimeMs          respjson.Field
		GroupBy              respjson.Field
		Groups               respjson.Field
		NextCursor           respjson.Field
		raw                  string
	} `json:"-"`
}

func (u EngagementResponseUnion) AsEngagement() (v Engagement) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EngagementResponseUnion) AsEngagementBreakdownPage() (v EngagementBreakdownPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EngagementResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *EngagementResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EngagementGetParams struct {
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
	Continent EngagementGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType EngagementGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
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
	GroupBy EngagementGetParamsGroupBy `query:"group_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EngagementGetParams]'s query parameters as `url.Values`.
func (r EngagementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type EngagementGetParamsContinent string

const (
	EngagementGetParamsContinentAf EngagementGetParamsContinent = "AF"
	EngagementGetParamsContinentAn EngagementGetParamsContinent = "AN"
	EngagementGetParamsContinentAs EngagementGetParamsContinent = "AS"
	EngagementGetParamsContinentEu EngagementGetParamsContinent = "EU"
	EngagementGetParamsContinentNa EngagementGetParamsContinent = "NA"
	EngagementGetParamsContinentOc EngagementGetParamsContinent = "OC"
	EngagementGetParamsContinentSa EngagementGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type EngagementGetParamsDeviceType string

const (
	EngagementGetParamsDeviceTypeDesktop EngagementGetParamsDeviceType = "desktop"
	EngagementGetParamsDeviceTypeMobile  EngagementGetParamsDeviceType = "mobile"
	EngagementGetParamsDeviceTypeTablet  EngagementGetParamsDeviceType = "tablet"
	EngagementGetParamsDeviceTypeUnknown EngagementGetParamsDeviceType = "unknown"
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
type EngagementGetParamsGroupBy string

const (
	EngagementGetParamsGroupByBrowserName      EngagementGetParamsGroupBy = "browser_name"
	EngagementGetParamsGroupByOsName           EngagementGetParamsGroupBy = "os_name"
	EngagementGetParamsGroupByDeviceType       EngagementGetParamsGroupBy = "device_type"
	EngagementGetParamsGroupByContinent        EngagementGetParamsGroupBy = "continent"
	EngagementGetParamsGroupByCountry          EngagementGetParamsGroupBy = "country"
	EngagementGetParamsGroupByCity             EngagementGetParamsGroupBy = "city"
	EngagementGetParamsGroupByPageOrigin       EngagementGetParamsGroupBy = "page_origin"
	EngagementGetParamsGroupByPagePath         EngagementGetParamsGroupBy = "page_path"
	EngagementGetParamsGroupBySubtitleLanguage EngagementGetParamsGroupBy = "subtitle_language"
	EngagementGetParamsGroupByPlaybackRate     EngagementGetParamsGroupBy = "playback_rate"
	EngagementGetParamsGroupByVideoTitle       EngagementGetParamsGroupBy = "video_title"
	EngagementGetParamsGroupByUserID           EngagementGetParamsGroupBy = "user_id"
)
