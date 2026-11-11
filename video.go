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
	"github.com/chunkifydev/chunkify-analytics-go/packages/pagination"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
	"github.com/chunkifydev/chunkify-analytics-go/shared"
)

// Private retained-event and view-cohort queries.
//
// VideoService contains methods and other services that help with interacting with
// the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoService] method instead.
type VideoService struct {
	Options []option.RequestOption
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r VideoService) {
	r = VideoService{}
	r.Options = opts
	return
}

// List per-video totals. Sort the entire filtered cohort by video_id, views,
// unique_viewers, total_playing_time_ms or playback_score before pagination.
// Defaults to video_id ascending. Missing scores come last; metric ties use
// video_id ascending in both directions. Unique viewer counts cannot be added
// across videos. Keep sort_by, order and filters unchanged with a cursor.
func (r *VideoService) List(ctx context.Context, query VideoListParams, opts ...option.RequestOption) (res *pagination.VideosPage[VideoMetrics], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/videos"
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

// List per-video totals. Sort the entire filtered cohort by video_id, views,
// unique_viewers, total_playing_time_ms or playback_score before pagination.
// Defaults to video_id ascending. Missing scores come last; metric ties use
// video_id ascending in both directions. Unique viewer counts cannot be added
// across videos. Keep sort_by, order and filters unchanged with a cursor.
func (r *VideoService) ListAutoPaging(ctx context.Context, query VideoListParams, opts ...option.RequestOption) *pagination.VideosPageAutoPager[VideoMetrics] {
	return pagination.NewVideosPageAutoPager(r.List(ctx, query, opts...))
}

type VideoMetrics struct {
	PlaybackScore      shared.PlaybackScore `json:"playback_score" api:"required"`
	TotalPlayingTimeMs float64              `json:"total_playing_time_ms" api:"required"`
	UniqueViewers      int64                `json:"unique_viewers" api:"required"`
	VideoID            string               `json:"video_id" api:"required"`
	// Title from the most recently received selected start that supplied one.
	VideoTitle string `json:"video_title" api:"required"`
	Views      int64  `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlaybackScore      respjson.Field
		TotalPlayingTimeMs respjson.Field
		UniqueViewers      respjson.Field
		VideoID            respjson.Field
		VideoTitle         respjson.Field
		Views              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoMetrics) RawJSON() string { return r.JSON.raw }
func (r *VideoMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoPage struct {
	NextCursor string         `json:"next_cursor" api:"required"`
	Videos     []VideoMetrics `json:"videos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		Videos      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoPage) RawJSON() string { return r.JSON.raw }
func (r *VideoPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoListParams struct {
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
	Continent VideoListParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType VideoListParamsDeviceType `query:"device_type,omitzero" json:"-"`
	// Sort direction. Events and videos default to asc; views default to desc. Events
	// sort by ingestion sequence. Videos and views sort by sort_by with missing values
	// last. View ties use view_id in the requested direction; video metric ties use
	// video_id ascending.
	//
	// Any of "asc", "desc".
	Order VideoListParamsOrder `query:"order,omitzero" json:"-"`
	// Rank all matching videos before pagination. Missing scores come last in both
	// directions. Metric ties use video_id ascending. Keep sort_by, order and filters
	// unchanged with a cursor.
	//
	// Any of "video_id", "views", "unique_viewers", "total_playing_time_ms",
	// "playback_score".
	SortBy VideoListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoListParams]'s query parameters as `url.Values`.
func (r VideoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type VideoListParamsContinent string

const (
	VideoListParamsContinentAf VideoListParamsContinent = "AF"
	VideoListParamsContinentAn VideoListParamsContinent = "AN"
	VideoListParamsContinentAs VideoListParamsContinent = "AS"
	VideoListParamsContinentEu VideoListParamsContinent = "EU"
	VideoListParamsContinentNa VideoListParamsContinent = "NA"
	VideoListParamsContinentOc VideoListParamsContinent = "OC"
	VideoListParamsContinentSa VideoListParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type VideoListParamsDeviceType string

const (
	VideoListParamsDeviceTypeDesktop VideoListParamsDeviceType = "desktop"
	VideoListParamsDeviceTypeMobile  VideoListParamsDeviceType = "mobile"
	VideoListParamsDeviceTypeTablet  VideoListParamsDeviceType = "tablet"
	VideoListParamsDeviceTypeUnknown VideoListParamsDeviceType = "unknown"
)

// Sort direction. Events and videos default to asc; views default to desc. Events
// sort by ingestion sequence. Videos and views sort by sort_by with missing values
// last. View ties use view_id in the requested direction; video metric ties use
// video_id ascending.
type VideoListParamsOrder string

const (
	VideoListParamsOrderAsc  VideoListParamsOrder = "asc"
	VideoListParamsOrderDesc VideoListParamsOrder = "desc"
)

// Rank all matching videos before pagination. Missing scores come last in both
// directions. Metric ties use video_id ascending. Keep sort_by, order and filters
// unchanged with a cursor.
type VideoListParamsSortBy string

const (
	VideoListParamsSortByVideoID            VideoListParamsSortBy = "video_id"
	VideoListParamsSortByViews              VideoListParamsSortBy = "views"
	VideoListParamsSortByUniqueViewers      VideoListParamsSortBy = "unique_viewers"
	VideoListParamsSortByTotalPlayingTimeMs VideoListParamsSortBy = "total_playing_time_ms"
	VideoListParamsSortByPlaybackScore      VideoListParamsSortBy = "playback_score"
)
