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
	shimjson "github.com/chunkifydev/chunkify-analytics-go/internal/encoding/json"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
	"github.com/chunkifydev/chunkify-analytics-go/shared/constant"
)

// ExportService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExportService] method instead.
type ExportService struct {
	Options []option.RequestOption
}

// NewExportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExportService(opts ...option.RequestOption) (r ExportService) {
	r = ExportService{}
	r.Options = opts
	return
}

// Create a scoped 256-bit random download token for all matching rows. Requires a
// project private key; read-only keys are allowed. The link expires after 30
// minutes and is invalidated by issuer-key rotation/revocation, project deletion
// or collector restart. No CSV is generated or retained at creation. limit/cursor
// are rejected. Existing query endpoints continue returning JSON. Event CSV files
// contain documented scalar event fields and flat context.\* columns. Custom
// attributes, nested arrays and historical OTLP payloads are excluded. CSV cells
// never contain JSON objects or arrays; use GET /v1/events for original payloads.
func (r *ExportService) New(ctx context.Context, body ExportNewParams, opts ...option.RequestOption) (res *Export, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The token authorizes only its saved project and query. Requires no Authorization
// header and accepts no query parameters. Generate CSV on demand from one read
// snapshot on a dedicated connection; stream with bounded application buffers and
// backpressure. Dataset is selected when the link is opened. Rows are not capped
// by query pagination. Nested metric columns use slash-separated paths. Events
// have documented scalar columns and flat context.\* columns. Custom attributes,
// nested arrays and historical OTLP payloads are excluded. CSV cells never contain
// JSON objects or arrays. Nulls are empty cells; spreadsheet-formula strings are
// prefixed with an apostrophe. Interrupted or failed downloads terminate the
// stream; retry the full download. No files are retained. Defaults: 2 download
// starts/minute per project with burst capacity 2, one active CSV generation per
// project and four server-wide. Limits are shared across keys and links. Excess
// requests receive 429 with Retry-After instead of queueing. Override limits in
// config.toml [rate_limits]. Link expiry is checked before starting, not during an
// active download. S3 retention/deletion can remove source objects during a
// download.
//
// The download URL contains the credential needed for this request. Treat it as a
// secret until it expires. You can call this method without configuring an API key
// on the client.
func (r *ExportService) Download(ctx context.Context, token string, opts ...option.RequestOption) (res *string, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exports/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Discover column names before creating an export. Requires a project private key;
// read-only keys are allowed. kind applies only to timeseries. group_by applies to
// summary, quality and engagement. Date ranges and filters do not change the
// columns.
func (r *ExportService) ListColumns(ctx context.Context, query ExportListColumnsParams, opts ...option.RequestOption) (res *ExportColumns, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Export struct {
	ContentType constant.TextCsv `json:"content_type" default:"text/csv"`
	DownloadURL string           `json:"download_url" api:"required" format:"uri"`
	ExpiresAt   time.Time        `json:"expires_at" api:"required" format:"date-time"`
	Filename    string           `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		DownloadURL respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Export) RawJSON() string { return r.JSON.raw }
func (r *Export) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportColumns struct {
	// Available CSV column names in default output order.
	Columns []string `json:"columns" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportColumns) RawJSON() string { return r.JSON.raw }
func (r *ExportColumns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportQueryParam struct {
	BrowserName   param.Opt[string]    `json:"browser_name,omitzero"`
	City          param.Opt[string]    `json:"city,omitzero"`
	Country       param.Opt[string]    `json:"country,omitzero"`
	From          param.Opt[time.Time] `json:"from,omitzero" format:"date-time"`
	HasFatalError param.Opt[bool]      `json:"has_fatal_error,omitzero"`
	OsName        param.Opt[string]    `json:"os_name,omitzero"`
	// Exact, case-sensitive page_origin from the first retained view start, or each
	// event for raw-event exports.
	PageOrigin param.Opt[string] `json:"page_origin,omitzero"`
	// Exact, case-sensitive page_path from the first retained view start, or each
	// event for raw-event exports.
	PagePath param.Opt[string]    `json:"page_path,omitzero"`
	To       param.Opt[time.Time] `json:"to,omitzero" format:"date-time"`
	VideoID  param.Opt[string]    `json:"video_id,omitzero"`
	ViewID   param.Opt[string]    `json:"view_id,omitzero"`
	ViewerID param.Opt[string]    `json:"viewer_id,omitzero"`
	// Any of "AF", "AN", "AS", "EU", "NA", "OC", "SA".
	Continent ExportQueryContinent `json:"continent,omitzero"`
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType ExportQueryDeviceType `json:"device_type,omitzero"`
	// Any of "browser_name", "os_name", "device_type", "video_id", "viewer_id",
	// "continent", "country", "city", "page_origin", "page_path".
	Dimension ExportQueryDimension `json:"dimension,omitzero"`
	// Any of "chunkify.video.view_started", "chunkify.video.view_snapshot",
	// "chunkify.video.playback_error", "chunkify.video.view_ended",
	// "chunkify.video.playback_state_changed", "chunkify.video.rendition_changed",
	// "chunkify.video.seek_completed".
	EventName ExportQueryEventName `json:"event_name,omitzero"`
	// Any of "browser_name", "os_name", "device_type", "continent", "country", "city",
	// "page_origin", "page_path".
	GroupBy ExportQueryGroupBy `json:"group_by,omitzero"`
	// Any of "hour", "day".
	Interval ExportQueryInterval `json:"interval,omitzero"`
	// Any of "summary", "quality", "engagement".
	Kind ExportQueryKind `json:"kind,omitzero"`
	// Any of "asc", "desc".
	Order ExportQueryOrder `json:"order,omitzero"`
	// Videos or views only. Use a sort field supported by the selected resource.
	// Defaults to video_id for videos and started_at for views.
	//
	// Any of "video_id", "views", "unique_viewers", "total_playing_time_ms",
	// "playback_score", "started_at", "view_id", "device_type", "playing_time_ms",
	// "startup_time_ms", "rebuffer_duration_ms".
	SortBy ExportQuerySortBy `json:"sort_by,omitzero"`
	paramObj
}

func (r ExportQueryParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportQueryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportQueryContinent string

const (
	ExportQueryContinentAf ExportQueryContinent = "AF"
	ExportQueryContinentAn ExportQueryContinent = "AN"
	ExportQueryContinentAs ExportQueryContinent = "AS"
	ExportQueryContinentEu ExportQueryContinent = "EU"
	ExportQueryContinentNa ExportQueryContinent = "NA"
	ExportQueryContinentOc ExportQueryContinent = "OC"
	ExportQueryContinentSa ExportQueryContinent = "SA"
)

type ExportQueryDeviceType string

const (
	ExportQueryDeviceTypeDesktop ExportQueryDeviceType = "desktop"
	ExportQueryDeviceTypeMobile  ExportQueryDeviceType = "mobile"
	ExportQueryDeviceTypeTablet  ExportQueryDeviceType = "tablet"
	ExportQueryDeviceTypeUnknown ExportQueryDeviceType = "unknown"
)

type ExportQueryDimension string

const (
	ExportQueryDimensionBrowserName ExportQueryDimension = "browser_name"
	ExportQueryDimensionOsName      ExportQueryDimension = "os_name"
	ExportQueryDimensionDeviceType  ExportQueryDimension = "device_type"
	ExportQueryDimensionVideoID     ExportQueryDimension = "video_id"
	ExportQueryDimensionViewerID    ExportQueryDimension = "viewer_id"
	ExportQueryDimensionContinent   ExportQueryDimension = "continent"
	ExportQueryDimensionCountry     ExportQueryDimension = "country"
	ExportQueryDimensionCity        ExportQueryDimension = "city"
	ExportQueryDimensionPageOrigin  ExportQueryDimension = "page_origin"
	ExportQueryDimensionPagePath    ExportQueryDimension = "page_path"
)

type ExportQueryEventName string

const (
	ExportQueryEventNameChunkifyVideoViewStarted          ExportQueryEventName = "chunkify.video.view_started"
	ExportQueryEventNameChunkifyVideoViewSnapshot         ExportQueryEventName = "chunkify.video.view_snapshot"
	ExportQueryEventNameChunkifyVideoPlaybackError        ExportQueryEventName = "chunkify.video.playback_error"
	ExportQueryEventNameChunkifyVideoViewEnded            ExportQueryEventName = "chunkify.video.view_ended"
	ExportQueryEventNameChunkifyVideoPlaybackStateChanged ExportQueryEventName = "chunkify.video.playback_state_changed"
	ExportQueryEventNameChunkifyVideoRenditionChanged     ExportQueryEventName = "chunkify.video.rendition_changed"
	ExportQueryEventNameChunkifyVideoSeekCompleted        ExportQueryEventName = "chunkify.video.seek_completed"
)

type ExportQueryGroupBy string

const (
	ExportQueryGroupByBrowserName ExportQueryGroupBy = "browser_name"
	ExportQueryGroupByOsName      ExportQueryGroupBy = "os_name"
	ExportQueryGroupByDeviceType  ExportQueryGroupBy = "device_type"
	ExportQueryGroupByContinent   ExportQueryGroupBy = "continent"
	ExportQueryGroupByCountry     ExportQueryGroupBy = "country"
	ExportQueryGroupByCity        ExportQueryGroupBy = "city"
	ExportQueryGroupByPageOrigin  ExportQueryGroupBy = "page_origin"
	ExportQueryGroupByPagePath    ExportQueryGroupBy = "page_path"
)

type ExportQueryInterval string

const (
	ExportQueryIntervalHour ExportQueryInterval = "hour"
	ExportQueryIntervalDay  ExportQueryInterval = "day"
)

type ExportQueryKind string

const (
	ExportQueryKindSummary    ExportQueryKind = "summary"
	ExportQueryKindQuality    ExportQueryKind = "quality"
	ExportQueryKindEngagement ExportQueryKind = "engagement"
)

type ExportQueryOrder string

const (
	ExportQueryOrderAsc  ExportQueryOrder = "asc"
	ExportQueryOrderDesc ExportQueryOrder = "desc"
)

// Videos or views only. Use a sort field supported by the selected resource.
// Defaults to video_id for videos and started_at for views.
type ExportQuerySortBy string

const (
	ExportQuerySortByVideoID            ExportQuerySortBy = "video_id"
	ExportQuerySortByViews              ExportQuerySortBy = "views"
	ExportQuerySortByUniqueViewers      ExportQuerySortBy = "unique_viewers"
	ExportQuerySortByTotalPlayingTimeMs ExportQuerySortBy = "total_playing_time_ms"
	ExportQuerySortByPlaybackScore      ExportQuerySortBy = "playback_score"
	ExportQuerySortByStartedAt          ExportQuerySortBy = "started_at"
	ExportQuerySortByViewID             ExportQuerySortBy = "view_id"
	ExportQuerySortByDeviceType         ExportQuerySortBy = "device_type"
	ExportQuerySortByPlayingTimeMs      ExportQuerySortBy = "playing_time_ms"
	ExportQuerySortByStartupTimeMs      ExportQuerySortBy = "startup_time_ms"
	ExportQuerySortByRebufferDurationMs ExportQuerySortBy = "rebuffer_duration_ms"
)

// The property Resource is required.
type ExportRequestParam struct {
	// Any of "summary", "quality", "engagement", "timeseries", "videos", "events",
	// "views", "errors", "dimensions", "renditions", "retention".
	Resource ExportRequestResource `json:"resource,omitzero" api:"required"`
	// CSV column names in output order. Omit for all columns. Get the available names
	// from GET /v1/exports using the same resource, kind and group_by. Unknown names,
	// duplicates and empty selections are rejected.
	Columns []string         `json:"columns,omitzero"`
	Query   ExportQueryParam `json:"query,omitzero"`
	paramObj
}

func (r ExportRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportRequestResource string

const (
	ExportRequestResourceSummary    ExportRequestResource = "summary"
	ExportRequestResourceQuality    ExportRequestResource = "quality"
	ExportRequestResourceEngagement ExportRequestResource = "engagement"
	ExportRequestResourceTimeseries ExportRequestResource = "timeseries"
	ExportRequestResourceVideos     ExportRequestResource = "videos"
	ExportRequestResourceEvents     ExportRequestResource = "events"
	ExportRequestResourceViews      ExportRequestResource = "views"
	ExportRequestResourceErrors     ExportRequestResource = "errors"
	ExportRequestResourceDimensions ExportRequestResource = "dimensions"
	ExportRequestResourceRenditions ExportRequestResource = "renditions"
	ExportRequestResourceRetention  ExportRequestResource = "retention"
)

type ExportNewParams struct {
	ExportRequest ExportRequestParam
	paramObj
}

func (r ExportNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExportRequest)
}
func (r *ExportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportListColumnsParams struct {
	// Report whose CSV columns to list.
	//
	// Any of "summary", "quality", "engagement", "timeseries", "videos", "events",
	// "views", "errors", "dimensions", "renditions", "retention".
	Resource ExportListColumnsParamsResource `query:"resource,omitzero" api:"required" json:"-"`
	// Optional on summary, quality and engagement. Group views by this field on their
	// first retained start. Missing, null and nonstring context forms a null group;
	// the literal device type unknown stays separate. Groups sort by views descending,
	// then value ascending with null first for ties. Pagination is available only with
	// group_by. Location uses the first retained start too. City groups use the stored
	// name; filter by country to avoid combining names shared by different countries.
	// Page groups use the first retained start even if playback continues after
	// navigation. Filter by page_origin when grouping page_path to distinguish
	// identical paths on different sites.
	//
	// Any of "browser_name", "os_name", "device_type", "continent", "country", "city",
	// "page_origin", "page_path".
	GroupBy ExportListColumnsParamsGroupBy `query:"group_by,omitzero" json:"-"`
	// Metric set for every time bucket. Uses the same calculations and sample counts
	// as the corresponding totals route.
	//
	// Any of "summary", "quality", "engagement".
	Kind ExportListColumnsParamsKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExportListColumnsParams]'s query parameters as
// `url.Values`.
func (r ExportListColumnsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Report whose CSV columns to list.
type ExportListColumnsParamsResource string

const (
	ExportListColumnsParamsResourceSummary    ExportListColumnsParamsResource = "summary"
	ExportListColumnsParamsResourceQuality    ExportListColumnsParamsResource = "quality"
	ExportListColumnsParamsResourceEngagement ExportListColumnsParamsResource = "engagement"
	ExportListColumnsParamsResourceTimeseries ExportListColumnsParamsResource = "timeseries"
	ExportListColumnsParamsResourceVideos     ExportListColumnsParamsResource = "videos"
	ExportListColumnsParamsResourceEvents     ExportListColumnsParamsResource = "events"
	ExportListColumnsParamsResourceViews      ExportListColumnsParamsResource = "views"
	ExportListColumnsParamsResourceErrors     ExportListColumnsParamsResource = "errors"
	ExportListColumnsParamsResourceDimensions ExportListColumnsParamsResource = "dimensions"
	ExportListColumnsParamsResourceRenditions ExportListColumnsParamsResource = "renditions"
	ExportListColumnsParamsResourceRetention  ExportListColumnsParamsResource = "retention"
)

// Optional on summary, quality and engagement. Group views by this field on their
// first retained start. Missing, null and nonstring context forms a null group;
// the literal device type unknown stays separate. Groups sort by views descending,
// then value ascending with null first for ties. Pagination is available only with
// group_by. Location uses the first retained start too. City groups use the stored
// name; filter by country to avoid combining names shared by different countries.
// Page groups use the first retained start even if playback continues after
// navigation. Filter by page_origin when grouping page_path to distinguish
// identical paths on different sites.
type ExportListColumnsParamsGroupBy string

const (
	ExportListColumnsParamsGroupByBrowserName ExportListColumnsParamsGroupBy = "browser_name"
	ExportListColumnsParamsGroupByOsName      ExportListColumnsParamsGroupBy = "os_name"
	ExportListColumnsParamsGroupByDeviceType  ExportListColumnsParamsGroupBy = "device_type"
	ExportListColumnsParamsGroupByContinent   ExportListColumnsParamsGroupBy = "continent"
	ExportListColumnsParamsGroupByCountry     ExportListColumnsParamsGroupBy = "country"
	ExportListColumnsParamsGroupByCity        ExportListColumnsParamsGroupBy = "city"
	ExportListColumnsParamsGroupByPageOrigin  ExportListColumnsParamsGroupBy = "page_origin"
	ExportListColumnsParamsGroupByPagePath    ExportListColumnsParamsGroupBy = "page_path"
)

// Metric set for every time bucket. Uses the same calculations and sample counts
// as the corresponding totals route.
type ExportListColumnsParamsKind string

const (
	ExportListColumnsParamsKindSummary    ExportListColumnsParamsKind = "summary"
	ExportListColumnsParamsKindQuality    ExportListColumnsParamsKind = "quality"
	ExportListColumnsParamsKindEngagement ExportListColumnsParamsKind = "engagement"
)
