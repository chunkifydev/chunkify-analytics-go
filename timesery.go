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
	"github.com/chunkifydev/chunkify-analytics-go/shared/constant"
)

// Private retained-event and view-cohort queries.
//
// TimeseryService contains methods and other services that help with interacting
// with the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimeseryService] method instead.
type TimeseryService struct {
	Options []option.RequestOption
}

// NewTimeseryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimeseryService(opts ...option.RequestOption) (r TimeseryService) {
	r = TimeseryService{}
	r.Options = opts
	return
}

// Supply from, to and interval=hour|day. kind defaults to summary and also accepts
// quality or engagement. Every bucket selects first retained starts in its UTC
// interval and includes each selected view's whole retained measurement and
// errors, even after to. Empty buckets have zero counts and null unavailable
// measurements or ratios. At most 1000 buckets. Unique viewer counts are not
// additive.
func (r *TimeseryService) Get(ctx context.Context, query TimeseryGetParams, opts ...option.RequestOption) (res *TimeseriesResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EngagementTimeseries struct {
	Buckets []EngagementTimeseriesBucket `json:"buckets" api:"required"`
	// Inclusive query bound normalized to UTC milliseconds.
	From time.Time `json:"from" api:"required" format:"date-time"`
	// Any of "hour", "day".
	Interval EngagementTimeseriesInterval `json:"interval" api:"required"`
	Kind     constant.Engagement          `json:"kind" default:"engagement"`
	Timezone constant.Utc                 `json:"timezone" default:"UTC"`
	// Exclusive query bound normalized to UTC milliseconds.
	To time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		From        respjson.Field
		Interval    respjson.Field
		Kind        respjson.Field
		Timezone    respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EngagementTimeseries) RawJSON() string { return r.JSON.raw }
func (r *EngagementTimeseries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EngagementTimeseriesInterval string

const (
	EngagementTimeseriesIntervalHour EngagementTimeseriesInterval = "hour"
	EngagementTimeseriesIntervalDay  EngagementTimeseriesInterval = "day"
)

// Observed watch time and end-of-media for the same start cohort as summary.
// Missing delivery can undercount both.
type EngagementTimeseriesBucket struct {
	// Total playing time divided by all selected attempts. Null with no views; zero
	// when views exist with no measured time.
	AveragePlayingTimeMs float64   `json:"average_playing_time_ms" api:"required"`
	BucketStart          time.Time `json:"bucket_start" api:"required" format:"date-time"`
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
		BucketStart          respjson.Field
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
func (r EngagementTimeseriesBucket) RawJSON() string { return r.JSON.raw }
func (r *EngagementTimeseriesBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualityTimeseries struct {
	Buckets []QualityTimeseriesBucket `json:"buckets" api:"required"`
	// Inclusive query bound normalized to UTC milliseconds.
	From time.Time `json:"from" api:"required" format:"date-time"`
	// Any of "hour", "day".
	Interval QualityTimeseriesInterval `json:"interval" api:"required"`
	Kind     constant.Quality          `json:"kind" default:"quality"`
	Timezone constant.Utc              `json:"timezone" default:"UTC"`
	// Exclusive query bound normalized to UTC milliseconds.
	To time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		From        respjson.Field
		Interval    respjson.Field
		Kind        respjson.Field
		Timezone    respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualityTimeseries) RawJSON() string { return r.JSON.raw }
func (r *QualityTimeseries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualityTimeseriesInterval string

const (
	QualityTimeseriesIntervalHour QualityTimeseriesInterval = "hour"
	QualityTimeseriesIntervalDay  QualityTimeseriesInterval = "day"
)

// Playback quality for the same start cohort as summary. All rates are fractions.
// Counts expose measurement coverage.
type QualityTimeseriesBucket struct {
	BucketStart time.Time `json:"bucket_start" api:"required" format:"date-time"`
	// Duration-weighted buffering. Startup wait, pauses and seeking are excluded by
	// the SDK. Unknown durations are not treated as zero.
	Buffering QualityTimeseriesBucketBuffering `json:"buffering" api:"required"`
	// Observed fatal errors. Lost events can undercount this metric; absence of an
	// error is not proof of successful playback.
	FatalErrors   QualityTimeseriesBucketFatalErrors `json:"fatal_errors" api:"required"`
	PlaybackScore shared.PlaybackScore               `json:"playback_score" api:"required"`
	// Mutually exclusive outcomes. An explicit startup_outcome on the selected final
	// event takes precedence. Otherwise an available startup time, including zero, or
	// positive playing time on the selected measurement proves started. All other
	// views are unknown. sample_views excludes unknown; every ratio uses all selected
	// attempts, including unknown. Missing final delivery does not prove an exit or
	// failure.
	StartupOutcomes shared.StartupOutcomes `json:"startup_outcomes" api:"required"`
	// Time from attempt to first playing, measured once per eligible view. This is not
	// first-frame timing.
	StartupTimeMs QualityTimeseriesBucketStartupTimeMs `json:"startup_time_ms" api:"required"`
	// Selected playback attempts with a retained start, including failed attempts.
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketStart     respjson.Field
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
func (r QualityTimeseriesBucket) RawJSON() string { return r.JSON.raw }
func (r *QualityTimeseriesBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Duration-weighted buffering. Startup wait, pauses and seeking are excluded by
// the SDK. Unknown durations are not treated as zero.
type QualityTimeseriesBucketBuffering struct {
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
func (r QualityTimeseriesBucketBuffering) RawJSON() string { return r.JSON.raw }
func (r *QualityTimeseriesBucketBuffering) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Observed fatal errors. Lost events can undercount this metric; absence of an
// error is not proof of successful playback.
type QualityTimeseriesBucketFatalErrors struct {
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
func (r QualityTimeseriesBucketFatalErrors) RawJSON() string { return r.JSON.raw }
func (r *QualityTimeseriesBucketFatalErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time from attempt to first playing, measured once per eligible view. This is not
// first-frame timing.
type QualityTimeseriesBucketStartupTimeMs struct {
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
func (r QualityTimeseriesBucketStartupTimeMs) RawJSON() string { return r.JSON.raw }
func (r *QualityTimeseriesBucketStartupTimeMs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeBucket struct {
	// UTC interval start in RFC3339 with milliseconds and Z. May precede from in a
	// partial first bucket; only starts inside [from, to) contribute.
	BucketStart        time.Time `json:"bucket_start" api:"required" format:"date-time"`
	TotalPlayingTimeMs float64   `json:"total_playing_time_ms" api:"required"`
	UniqueViewers      int64     `json:"unique_viewers" api:"required"`
	Views              int64     `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketStart        respjson.Field
		TotalPlayingTimeMs respjson.Field
		UniqueViewers      respjson.Field
		Views              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeBucket) RawJSON() string { return r.JSON.raw }
func (r *TimeBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Timeseries struct {
	Buckets []TimeBucket `json:"buckets" api:"required"`
	// Inclusive query bound normalized to UTC milliseconds.
	From time.Time `json:"from" api:"required" format:"date-time"`
	// Any of "hour", "day".
	Interval TimeseriesInterval `json:"interval" api:"required"`
	Kind     constant.Summary   `json:"kind" default:"summary"`
	Timezone constant.Utc       `json:"timezone" default:"UTC"`
	// Exclusive query bound normalized to UTC milliseconds.
	To time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		From        respjson.Field
		Interval    respjson.Field
		Kind        respjson.Field
		Timezone    respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Timeseries) RawJSON() string { return r.JSON.raw }
func (r *Timeseries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeseriesInterval string

const (
	TimeseriesIntervalHour TimeseriesInterval = "hour"
	TimeseriesIntervalDay  TimeseriesInterval = "day"
)

// TimeseriesResponseUnion contains all possible properties and values from
// [Timeseries], [QualityTimeseries], [EngagementTimeseries].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TimeseriesResponseUnion struct {
	// This field is a union of [[]TimeBucket], [[]QualityTimeseriesBucket],
	// [[]EngagementTimeseriesBucket]
	Buckets  TimeseriesResponseUnionBuckets `json:"buckets"`
	From     time.Time                      `json:"from"`
	Interval string                         `json:"interval"`
	Kind     string                         `json:"kind"`
	// This field is from variant [Timeseries].
	Timezone constant.Utc `json:"timezone"`
	To       time.Time    `json:"to"`
	JSON     struct {
		Buckets  respjson.Field
		From     respjson.Field
		Interval respjson.Field
		Kind     respjson.Field
		Timezone respjson.Field
		To       respjson.Field
		raw      string
	} `json:"-"`
}

func (u TimeseriesResponseUnion) AsTimeseries() (v Timeseries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TimeseriesResponseUnion) AsQualityTimeseries() (v QualityTimeseries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TimeseriesResponseUnion) AsEngagementTimeseries() (v EngagementTimeseries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TimeseriesResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TimeseriesResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TimeseriesResponseUnionBuckets is an implicit subunion of
// [TimeseriesResponseUnion]. TimeseriesResponseUnionBuckets provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TimeseriesResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTimeBucketArray OfQualityTimeseriesBucketArray
// OfEngagementTimeseriesBucketArray]
type TimeseriesResponseUnionBuckets struct {
	// This field will be present if the value is a [[]TimeBucket] instead of an
	// object.
	OfTimeBucketArray []TimeBucket `json:",inline"`
	// This field will be present if the value is a [[]QualityTimeseriesBucket] instead
	// of an object.
	OfQualityTimeseriesBucketArray []QualityTimeseriesBucket `json:",inline"`
	// This field will be present if the value is a [[]EngagementTimeseriesBucket]
	// instead of an object.
	OfEngagementTimeseriesBucketArray []EngagementTimeseriesBucket `json:",inline"`
	JSON                              struct {
		OfTimeBucketArray                 respjson.Field
		OfQualityTimeseriesBucketArray    respjson.Field
		OfEngagementTimeseriesBucketArray respjson.Field
		raw                               string
	} `json:"-"`
}

func (r *TimeseriesResponseUnionBuckets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeseryGetParams struct {
	// Inclusive view-start time in RFC3339 format with a timezone.
	From time.Time `query:"from" api:"required" format:"date-time" json:"-"`
	// UTC hour or day. The requested range can contain at most 1000 buckets.
	//
	// Any of "hour", "day".
	Interval TimeseryGetParamsInterval `query:"interval,omitzero" api:"required" json:"-"`
	// Exclusive view-start time in RFC3339 format with a timezone. Must be later than
	// from.
	To time.Time `query:"to" api:"required" format:"date-time" json:"-"`
	// Exact browser name, such as Chrome. Missing values do not match.
	BrowserName param.Opt[string] `query:"browser_name,omitzero" json:"-"`
	// Exact, case-sensitive city name from the local lookup, with 1 to 256 Unicode
	// characters. Use country alongside city to distinguish places with the same name.
	// Missing locations do not match.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Exact uppercase ISO 3166-1 alpha-2 country code, such as GB. Missing locations
	// do not match.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
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
	// Exact video ID. Combined with all other filters using AND.
	VideoID param.Opt[string] `query:"video_id,omitzero" json:"-"`
	// Filter one playback attempt by view ID. Metrics use its first retained start.
	ViewID param.Opt[string] `query:"view_id,omitzero" json:"-"`
	// Exact supplied viewer ID. Missing/null viewer IDs do not match.
	ViewerID param.Opt[string] `query:"viewer_id,omitzero" json:"-"`
	// Exact continent code from the local lookup. Missing locations do not match.
	//
	// Any of "AF", "AN", "AS", "EU", "NA", "OC", "SA".
	Continent TimeseryGetParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType TimeseryGetParamsDeviceType `query:"device_type,omitzero" json:"-"`
	// Metric set for every time bucket. Uses the same calculations and sample counts
	// as the corresponding totals route.
	//
	// Any of "summary", "quality", "engagement".
	Kind TimeseryGetParamsKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TimeseryGetParams]'s query parameters as `url.Values`.
func (r TimeseryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// UTC hour or day. The requested range can contain at most 1000 buckets.
type TimeseryGetParamsInterval string

const (
	TimeseryGetParamsIntervalHour TimeseryGetParamsInterval = "hour"
	TimeseryGetParamsIntervalDay  TimeseryGetParamsInterval = "day"
)

// Exact continent code from the local lookup. Missing locations do not match.
type TimeseryGetParamsContinent string

const (
	TimeseryGetParamsContinentAf TimeseryGetParamsContinent = "AF"
	TimeseryGetParamsContinentAn TimeseryGetParamsContinent = "AN"
	TimeseryGetParamsContinentAs TimeseryGetParamsContinent = "AS"
	TimeseryGetParamsContinentEu TimeseryGetParamsContinent = "EU"
	TimeseryGetParamsContinentNa TimeseryGetParamsContinent = "NA"
	TimeseryGetParamsContinentOc TimeseryGetParamsContinent = "OC"
	TimeseryGetParamsContinentSa TimeseryGetParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type TimeseryGetParamsDeviceType string

const (
	TimeseryGetParamsDeviceTypeDesktop TimeseryGetParamsDeviceType = "desktop"
	TimeseryGetParamsDeviceTypeMobile  TimeseryGetParamsDeviceType = "mobile"
	TimeseryGetParamsDeviceTypeTablet  TimeseryGetParamsDeviceType = "tablet"
	TimeseryGetParamsDeviceTypeUnknown TimeseryGetParamsDeviceType = "unknown"
)

// Metric set for every time bucket. Uses the same calculations and sample counts
// as the corresponding totals route.
type TimeseryGetParamsKind string

const (
	TimeseryGetParamsKindSummary    TimeseryGetParamsKind = "summary"
	TimeseryGetParamsKindQuality    TimeseryGetParamsKind = "quality"
	TimeseryGetParamsKindEngagement TimeseryGetParamsKind = "engagement"
)
