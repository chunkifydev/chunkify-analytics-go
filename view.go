// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"errors"
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
	"github.com/chunkifydev/chunkify-analytics-go/shared"
)

// Private retained-event and view-cohort queries.
//
// ViewService contains methods and other services that help with interacting with
// the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r ViewService) {
	r = ViewService{}
	r.Options = opts
	return
}

// Returns start context and selected measurements, plus a paginated history of all
// retained events for that view ID, in ingestion order. SDK sequence and
// occurrence times remain in each event. Follow next_cursor with the same view ID;
// the summary and event pages share the initial ingestion watermark. Orphaned
// events without a retained start remain available through /v1/events?view_id=...
// .
func (r *ViewService) Get(ctx context.Context, viewID string, query ViewGetParams, opts ...option.RequestOption) (res *ViewDetail, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/views/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// One row per first retained start, using the same final/latest measurement
// selection as metrics. Filter by identity, context, dates or has_fatal_error.
// Order by started_at descending by default (newest first), or ascending with
// order=asc. Equal start times use view_id in the same direction. Cursors bind
// filters, order, fatal-error selection, project and ingestion watermark. Missing
// end or error events are not proof of an ongoing or successful playback. Sort the
// entire matching cohort by sort_by before pagination, with missing values last.
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.ViewsPage[View], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/views"
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

// One row per first retained start, using the same final/latest measurement
// selection as metrics. Filter by identity, context, dates or has_fatal_error.
// Order by started_at descending by default (newest first), or ascending with
// order=asc. Equal start times use view_id in the same direction. Cursors bind
// filters, order, fatal-error selection, project and ingestion watermark. Missing
// end or error events are not proof of an ongoing or successful playback. Sort the
// entire matching cohort by sort_by before pagination, with missing values last.
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.ViewsPageAutoPager[View] {
	return pagination.NewViewsPageAutoPager(r.List(ctx, query, opts...))
}

type View struct {
	BrowserName        string               `json:"browser_name" api:"required"`
	DeviceType         string               `json:"device_type" api:"required"`
	HasEndedEvent      bool                 `json:"has_ended_event" api:"required"`
	HasFatalError      bool                 `json:"has_fatal_error" api:"required"`
	OsName             string               `json:"os_name" api:"required"`
	PlaybackScore      shared.PlaybackScore `json:"playback_score" api:"required"`
	PlayingTimeMs      float64              `json:"playing_time_ms" api:"required"`
	ReachedEnd         bool                 `json:"reached_end" api:"required"`
	RebufferDurationMs float64              `json:"rebuffer_duration_ms" api:"required"`
	StartedAt          time.Time            `json:"started_at" api:"required" format:"date-time"`
	StartupTimeMs      float64              `json:"startup_time_ms" api:"required"`
	UserID             string               `json:"user_id" api:"required"`
	VideoID            string               `json:"video_id" api:"required"`
	VideoTitle         string               `json:"video_title" api:"required"`
	ViewID             string               `json:"view_id" api:"required"`
	ViewerID           string               `json:"viewer_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserName        respjson.Field
		DeviceType         respjson.Field
		HasEndedEvent      respjson.Field
		HasFatalError      respjson.Field
		OsName             respjson.Field
		PlaybackScore      respjson.Field
		PlayingTimeMs      respjson.Field
		ReachedEnd         respjson.Field
		RebufferDurationMs respjson.Field
		StartedAt          respjson.Field
		StartupTimeMs      respjson.Field
		UserID             respjson.Field
		VideoID            respjson.Field
		VideoTitle         respjson.Field
		ViewID             respjson.Field
		ViewerID           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r View) RawJSON() string { return r.JSON.raw }
func (r *View) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewDetail struct {
	Events     []StoredEvent `json:"events" api:"required"`
	NextCursor string        `json:"next_cursor" api:"required"`
	View       View          `json:"view" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		NextCursor  respjson.Field
		View        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewDetail) RawJSON() string { return r.JSON.raw }
func (r *ViewDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewPage struct {
	NextCursor string `json:"next_cursor" api:"required"`
	Views      []View `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewPage) RawJSON() string { return r.JSON.raw }
func (r *ViewPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewGetParams struct {
	// Use next_cursor from the previous page. Keep the same filters, sort_by and
	// order. Video cursors from earlier versions are invalid; start a new pagination
	// sequence after upgrading.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum rows on this page. Events and videos always support pagination; summary,
	// quality and engagement support it only with group_by.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewListParams struct {
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
	// Select attempts with or without an observed fatal error. Absence of an observed
	// error does not prove successful playback.
	HasFatalError param.Opt[bool] `query:"has_fatal_error,omitzero" json:"-"`
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
	Continent ViewListParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType ViewListParamsDeviceType `query:"device_type,omitzero" json:"-"`
	// Sort by started_at and then view_id in this direction. Defaults to newest first.
	// Keep the same order with each cursor.
	//
	// Any of "asc", "desc".
	Order ViewListParamsOrder `query:"order,omitzero" json:"-"`
	// Sort all matching sessions before pagination. Missing values come last in both
	// directions. Ties use view_id in the requested direction. Keep sort_by, order and
	// filters unchanged with a cursor.
	//
	// Any of "started_at", "video_id", "view_id", "device_type", "playing_time_ms",
	// "startup_time_ms", "rebuffer_duration_ms", "playback_score".
	SortBy ViewListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type ViewListParamsContinent string

const (
	ViewListParamsContinentAf ViewListParamsContinent = "AF"
	ViewListParamsContinentAn ViewListParamsContinent = "AN"
	ViewListParamsContinentAs ViewListParamsContinent = "AS"
	ViewListParamsContinentEu ViewListParamsContinent = "EU"
	ViewListParamsContinentNa ViewListParamsContinent = "NA"
	ViewListParamsContinentOc ViewListParamsContinent = "OC"
	ViewListParamsContinentSa ViewListParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type ViewListParamsDeviceType string

const (
	ViewListParamsDeviceTypeDesktop ViewListParamsDeviceType = "desktop"
	ViewListParamsDeviceTypeMobile  ViewListParamsDeviceType = "mobile"
	ViewListParamsDeviceTypeTablet  ViewListParamsDeviceType = "tablet"
	ViewListParamsDeviceTypeUnknown ViewListParamsDeviceType = "unknown"
)

// Sort by started_at and then view_id in this direction. Defaults to newest first.
// Keep the same order with each cursor.
type ViewListParamsOrder string

const (
	ViewListParamsOrderAsc  ViewListParamsOrder = "asc"
	ViewListParamsOrderDesc ViewListParamsOrder = "desc"
)

// Sort all matching sessions before pagination. Missing values come last in both
// directions. Ties use view_id in the requested direction. Keep sort_by, order and
// filters unchanged with a cursor.
type ViewListParamsSortBy string

const (
	ViewListParamsSortByStartedAt          ViewListParamsSortBy = "started_at"
	ViewListParamsSortByVideoID            ViewListParamsSortBy = "video_id"
	ViewListParamsSortByViewID             ViewListParamsSortBy = "view_id"
	ViewListParamsSortByDeviceType         ViewListParamsSortBy = "device_type"
	ViewListParamsSortByPlayingTimeMs      ViewListParamsSortBy = "playing_time_ms"
	ViewListParamsSortByStartupTimeMs      ViewListParamsSortBy = "startup_time_ms"
	ViewListParamsSortByRebufferDurationMs ViewListParamsSortBy = "rebuffer_duration_ms"
	ViewListParamsSortByPlaybackScore      ViewListParamsSortBy = "playback_score"
)
