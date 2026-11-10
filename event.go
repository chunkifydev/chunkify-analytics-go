// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"context"
	"encoding/json"
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
	"github.com/chunkifydev/chunkify-analytics-go/packages/pagination"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
	"github.com/chunkifydev/chunkify-analytics-go/shared"
	"github.com/chunkifydev/chunkify-analytics-go/shared/constant"
)

// EventService contains methods and other services that help with interacting with
// the Chunkify Analytics API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// List project events with filters and pagination. Use order=desc for the newest
// events.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.EventsPage[StoredEvent], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecretKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/events"
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

// List project events with filters and pagination. Use order=desc for the newest
// events.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.EventsPageAutoPager[StoredEvent] {
	return pagination.NewEventsPageAutoPager(r.List(ctx, query, opts...))
}

// Send one playback event using a public key. Success confirms storage completed.
//
// Configure the client with a public ingestion key for this method. The SDK sends
// the event you supply; it does not collect playback events automatically. For
// automatic browser player tracking, use `@chunkify/player-analytics`.
func (r *EventService) Ingest(ctx context.Context, params EventIngestParams, opts ...option.RequestOption) (res *EventIngestResponseUnion, err error) {
	if !param.IsOmitted(params.ContentEncoding) {
		opts = append(opts, option.WithHeader("Content-Encoding", fmt.Sprintf("%v", params.ContentEncoding)))
	}
	if !param.IsOmitted(params.Origin) {
		opts = append(opts, option.WithHeader("Origin", fmt.Sprintf("%v", params.Origin.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithPublicKeySecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// AnalyticsEventUnion contains all possible properties and values from
// [ViewStarted], [ViewSnapshot], [PlaybackError], [ViewEnded],
// [PlaybackStateChanged], [RenditionChanged], [SeekCompleted], [SubtitleChanged],
// [PlaybackRateChanged].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AnalyticsEventUnion struct {
	ConnectorName            string  `json:"connector_name"`
	ConnectorVersion         string  `json:"connector_version"`
	DeliveryHostname         string  `json:"delivery_hostname"`
	DeliveryRequestCount     float64 `json:"delivery_request_count"`
	DeliveryThroughputBps    float64 `json:"delivery_throughput_bps"`
	DeliveryTransferredBytes float64 `json:"delivery_transferred_bytes"`
	DroppedVideoFrames       float64 `json:"dropped_video_frames"`
	Environment              string  `json:"environment"`
	ErrorCategory            string  `json:"error_category"`
	// This field is a union of [ViewStartedErrorCodeUnion],
	// [ViewSnapshotErrorCodeUnion], [PlaybackErrorErrorCodeUnion],
	// [ViewEndedErrorCodeUnion], [PlaybackStateChangedErrorCodeUnion],
	// [RenditionChangedErrorCodeUnion], [SeekCompletedErrorCodeUnion],
	// [SubtitleChangedErrorCodeUnion], [PlaybackRateChangedErrorCodeUnion]
	ErrorCode                   AnalyticsEventUnionErrorCode  `json:"error_code"`
	ErrorFatal                  bool                          `json:"error_fatal"`
	ErrorPlayheadSeconds        float64                       `json:"error_playhead_seconds"`
	EventID                     string                        `json:"event_id"`
	EventName                   string                        `json:"event_name"`
	FirstVideoFrameTimeMs       float64                       `json:"first_video_frame_time_ms"`
	FurthestProgressRatio       float64                       `json:"furthest_progress_ratio"`
	MaximumPlayheadSeconds      float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds        float64                       `json:"media_duration_seconds"`
	MediaHeightPx               float64                       `json:"media_height_px"`
	MediaWidthPx                float64                       `json:"media_width_px"`
	OccurredAt                  time.Time                     `json:"occurred_at"`
	PlayerHeightPx              float64                       `json:"player_height_px"`
	PlayerName                  string                        `json:"player_name"`
	PlayerVersion               string                        `json:"player_version"`
	PlayerWidthPx               float64                       `json:"player_width_px"`
	PlayingTimeMs               float64                       `json:"playing_time_ms"`
	Reason                      string                        `json:"reason"`
	RebufferCount               float64                       `json:"rebuffer_count"`
	RebufferDurationMs          float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio               float64                       `json:"rebuffer_ratio"`
	RenditionPlayingTime        []shared.RenditionPlayingTime `json:"rendition_playing_time"`
	RenditionSwitchCount        float64                       `json:"rendition_switch_count"`
	SchemaVersion               float64                       `json:"schema_version"`
	SDKName                     string                        `json:"sdk_name"`
	SDKVersion                  string                        `json:"sdk_version"`
	SeekCount                   float64                       `json:"seek_count"`
	SeekDurationMs              float64                       `json:"seek_duration_ms"`
	SelectedRenditionBitrateBps float64                       `json:"selected_rendition_bitrate_bps"`
	SelectedRenditionCodec      string                        `json:"selected_rendition_codec"`
	SelectedRenditionHeight     float64                       `json:"selected_rendition_height"`
	SelectedRenditionWidth      float64                       `json:"selected_rendition_width"`
	Sequence                    int64                         `json:"sequence"`
	StartupOutcome              string                        `json:"startup_outcome"`
	StartupTimeMs               float64                       `json:"startup_time_ms"`
	TotalVideoFrames            float64                       `json:"total_video_frames"`
	VideoID                     string                        `json:"video_id"`
	ViewID                      string                        `json:"view_id"`
	ViewerID                    string                        `json:"viewer_id"`
	BrowserContext              any                           `json:"browser_context"`
	// This field is from variant [ViewStarted].
	Custom                  shared.CustomDimensions          `json:"custom"`
	PlaybackRate            float64                          `json:"playback_rate"`
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTime `json:"playback_rate_playing_time"`
	PlayedRanges            []shared.PlayedRange             `json:"played_ranges"`
	PlayheadSeconds         float64                          `json:"playhead_seconds"`
	SubtitleLanguage        string                           `json:"subtitle_language"`
	SubtitlePlayingTime     []shared.SubtitlePlayingTime     `json:"subtitle_playing_time"`
	UserID                  string                           `json:"user_id"`
	VideoTitle              string                           `json:"video_title"`
	VisibilityState         string                           `json:"visibility_state"`
	// This field is from variant [PlaybackStateChanged].
	FromState PlaybackStateChangedFromState `json:"from_state"`
	// This field is from variant [PlaybackStateChanged].
	PreviousStateDurationMs float64 `json:"previous_state_duration_ms"`
	// This field is from variant [PlaybackStateChanged].
	ToState PlaybackStateChangedToState `json:"to_state"`
	// This field is from variant [RenditionChanged].
	PreviousRenditionBitrateBps float64 `json:"previous_rendition_bitrate_bps"`
	// This field is from variant [RenditionChanged].
	PreviousRenditionCodec string `json:"previous_rendition_codec"`
	// This field is from variant [RenditionChanged].
	PreviousRenditionFrameRate float64 `json:"previous_rendition_frame_rate"`
	// This field is from variant [RenditionChanged].
	PreviousRenditionHeight float64 `json:"previous_rendition_height"`
	// This field is from variant [RenditionChanged].
	PreviousRenditionWidth float64 `json:"previous_rendition_width"`
	// This field is from variant [RenditionChanged].
	SelectedRenditionFrameRate float64 `json:"selected_rendition_frame_rate"`
	// This field is from variant [SeekCompleted].
	FromPlayheadSeconds float64 `json:"from_playhead_seconds"`
	// This field is from variant [SeekCompleted].
	LastSeekedAt time.Time `json:"last_seeked_at"`
	// This field is from variant [SeekCompleted].
	PlaybackResumedAt time.Time `json:"playback_resumed_at"`
	// This field is from variant [SeekCompleted].
	SeekInteractionDurationMs float64 `json:"seek_interaction_duration_ms"`
	// This field is from variant [SeekCompleted].
	SeekOutcome SeekCompletedSeekOutcome `json:"seek_outcome"`
	// This field is from variant [SeekCompleted].
	SeekResumeWaitMs float64 `json:"seek_resume_wait_ms"`
	// This field is from variant [SeekCompleted].
	SeekStartedAt time.Time `json:"seek_started_at"`
	// This field is from variant [SeekCompleted].
	StateAfterSeek SeekCompletedStateAfterSeek `json:"state_after_seek"`
	// This field is from variant [SeekCompleted].
	StateBeforeSeek SeekCompletedStateBeforeSeek `json:"state_before_seek"`
	// This field is from variant [SeekCompleted].
	ToPlayheadSeconds float64 `json:"to_playhead_seconds"`
	// This field is from variant [SubtitleChanged].
	PreviousSubtitleLanguage string `json:"previous_subtitle_language"`
	// This field is from variant [PlaybackRateChanged].
	PreviousPlaybackRate float64 `json:"previous_playback_rate"`
	JSON                 struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		Reason                      respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		Custom                      respjson.Field
		PlaybackRate                respjson.Field
		PlaybackRatePlayingTime     respjson.Field
		PlayedRanges                respjson.Field
		PlayheadSeconds             respjson.Field
		SubtitleLanguage            respjson.Field
		SubtitlePlayingTime         respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		FromState                   respjson.Field
		PreviousStateDurationMs     respjson.Field
		ToState                     respjson.Field
		PreviousRenditionBitrateBps respjson.Field
		PreviousRenditionCodec      respjson.Field
		PreviousRenditionFrameRate  respjson.Field
		PreviousRenditionHeight     respjson.Field
		PreviousRenditionWidth      respjson.Field
		SelectedRenditionFrameRate  respjson.Field
		FromPlayheadSeconds         respjson.Field
		LastSeekedAt                respjson.Field
		PlaybackResumedAt           respjson.Field
		SeekInteractionDurationMs   respjson.Field
		SeekOutcome                 respjson.Field
		SeekResumeWaitMs            respjson.Field
		SeekStartedAt               respjson.Field
		StateAfterSeek              respjson.Field
		StateBeforeSeek             respjson.Field
		ToPlayheadSeconds           respjson.Field
		PreviousSubtitleLanguage    respjson.Field
		PreviousPlaybackRate        respjson.Field
		raw                         string
	} `json:"-"`
}

func (u AnalyticsEventUnion) AsViewStarted() (v ViewStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsViewSnapshot() (v ViewSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsPlaybackError() (v PlaybackError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsViewEnded() (v ViewEnded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsPlaybackStateChanged() (v PlaybackStateChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsRenditionChanged() (v RenditionChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsSeekCompleted() (v SeekCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsSubtitleChanged() (v SubtitleChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnalyticsEventUnion) AsPlaybackRateChanged() (v PlaybackRateChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AnalyticsEventUnion) RawJSON() string { return u.JSON.raw }

func (r *AnalyticsEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AnalyticsEventUnionErrorCode is an implicit subunion of [AnalyticsEventUnion].
// AnalyticsEventUnionErrorCode provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [AnalyticsEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type AnalyticsEventUnionErrorCode struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (r *AnalyticsEventUnionErrorCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AnalyticsEventUnion to a AnalyticsEventUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AnalyticsEventUnionParam.Overrides()
func (r AnalyticsEventUnion) ToParam() AnalyticsEventUnionParam {
	return param.Override[AnalyticsEventUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AnalyticsEventUnionParam struct {
	OfViewStarted          *ViewStartedParam          `json:",omitzero,inline"`
	OfViewSnapshot         *ViewSnapshotParam         `json:",omitzero,inline"`
	OfPlaybackError        *PlaybackErrorParam        `json:",omitzero,inline"`
	OfViewEnded            *ViewEndedParam            `json:",omitzero,inline"`
	OfPlaybackStateChanged *PlaybackStateChangedParam `json:",omitzero,inline"`
	OfRenditionChanged     *RenditionChangedParam     `json:",omitzero,inline"`
	OfSeekCompleted        *SeekCompletedParam        `json:",omitzero,inline"`
	OfSubtitleChanged      *SubtitleChangedParam      `json:",omitzero,inline"`
	OfPlaybackRateChanged  *PlaybackRateChangedParam  `json:",omitzero,inline"`
	paramUnion
}

func (u AnalyticsEventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfViewStarted,
		u.OfViewSnapshot,
		u.OfPlaybackError,
		u.OfViewEnded,
		u.OfPlaybackStateChanged,
		u.OfRenditionChanged,
		u.OfSeekCompleted,
		u.OfSubtitleChanged,
		u.OfPlaybackRateChanged)
}
func (u *AnalyticsEventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AnalyticsEventUnionParam) asAny() any {
	if !param.IsOmitted(u.OfViewStarted) {
		return u.OfViewStarted
	} else if !param.IsOmitted(u.OfViewSnapshot) {
		return u.OfViewSnapshot
	} else if !param.IsOmitted(u.OfPlaybackError) {
		return u.OfPlaybackError
	} else if !param.IsOmitted(u.OfViewEnded) {
		return u.OfViewEnded
	} else if !param.IsOmitted(u.OfPlaybackStateChanged) {
		return u.OfPlaybackStateChanged
	} else if !param.IsOmitted(u.OfRenditionChanged) {
		return u.OfRenditionChanged
	} else if !param.IsOmitted(u.OfSeekCompleted) {
		return u.OfSeekCompleted
	} else if !param.IsOmitted(u.OfSubtitleChanged) {
		return u.OfSubtitleChanged
	} else if !param.IsOmitted(u.OfPlaybackRateChanged) {
		return u.OfPlaybackRateChanged
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetCustom() shared.CustomDimensionsParam {
	if vt := u.OfViewStarted; vt != nil {
		return vt.Custom
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetFromState() *string {
	if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.FromState)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousStateDurationMs() *float64 {
	if vt := u.OfPlaybackStateChanged; vt != nil {
		return &vt.PreviousStateDurationMs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetToState() *string {
	if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.ToState)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousRenditionBitrateBps() *float64 {
	if vt := u.OfRenditionChanged; vt != nil && vt.PreviousRenditionBitrateBps.Valid() {
		return &vt.PreviousRenditionBitrateBps.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousRenditionCodec() *string {
	if vt := u.OfRenditionChanged; vt != nil && vt.PreviousRenditionCodec.Valid() {
		return &vt.PreviousRenditionCodec.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousRenditionFrameRate() *float64 {
	if vt := u.OfRenditionChanged; vt != nil && vt.PreviousRenditionFrameRate.Valid() {
		return &vt.PreviousRenditionFrameRate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousRenditionHeight() *float64 {
	if vt := u.OfRenditionChanged; vt != nil && vt.PreviousRenditionHeight.Valid() {
		return &vt.PreviousRenditionHeight.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousRenditionWidth() *float64 {
	if vt := u.OfRenditionChanged; vt != nil && vt.PreviousRenditionWidth.Valid() {
		return &vt.PreviousRenditionWidth.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSelectedRenditionFrameRate() *float64 {
	if vt := u.OfRenditionChanged; vt != nil && vt.SelectedRenditionFrameRate.Valid() {
		return &vt.SelectedRenditionFrameRate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetFromPlayheadSeconds() *float64 {
	if vt := u.OfSeekCompleted; vt != nil && vt.FromPlayheadSeconds.Valid() {
		return &vt.FromPlayheadSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetLastSeekedAt() *time.Time {
	if vt := u.OfSeekCompleted; vt != nil && vt.LastSeekedAt.Valid() {
		return &vt.LastSeekedAt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlaybackResumedAt() *time.Time {
	if vt := u.OfSeekCompleted; vt != nil && vt.PlaybackResumedAt.Valid() {
		return &vt.PlaybackResumedAt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekInteractionDurationMs() *float64 {
	if vt := u.OfSeekCompleted; vt != nil {
		return &vt.SeekInteractionDurationMs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekOutcome() *string {
	if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.SeekOutcome)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekResumeWaitMs() *float64 {
	if vt := u.OfSeekCompleted; vt != nil && vt.SeekResumeWaitMs.Valid() {
		return &vt.SeekResumeWaitMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekStartedAt() *time.Time {
	if vt := u.OfSeekCompleted; vt != nil {
		return &vt.SeekStartedAt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetStateAfterSeek() *string {
	if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.StateAfterSeek)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetStateBeforeSeek() *string {
	if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.StateBeforeSeek)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetToPlayheadSeconds() *float64 {
	if vt := u.OfSeekCompleted; vt != nil && vt.ToPlayheadSeconds.Valid() {
		return &vt.ToPlayheadSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousSubtitleLanguage() *string {
	if vt := u.OfSubtitleChanged; vt != nil && vt.PreviousSubtitleLanguage.Valid() {
		return &vt.PreviousSubtitleLanguage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPreviousPlaybackRate() *float64 {
	if vt := u.OfPlaybackRateChanged; vt != nil && vt.PreviousPlaybackRate.Valid() {
		return &vt.PreviousPlaybackRate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetConnectorName() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.ConnectorName)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.ConnectorName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetConnectorVersion() *string {
	if vt := u.OfViewStarted; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.ConnectorVersion.Valid() {
		return &vt.ConnectorVersion.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetDeliveryHostname() *string {
	if vt := u.OfViewStarted; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.DeliveryHostname.Valid() {
		return &vt.DeliveryHostname.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetDeliveryRequestCount() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.DeliveryRequestCount.Valid() {
		return &vt.DeliveryRequestCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetDeliveryThroughputBps() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.DeliveryThroughputBps.Valid() {
		return &vt.DeliveryThroughputBps.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetDeliveryTransferredBytes() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.DeliveryTransferredBytes.Valid() {
		return &vt.DeliveryTransferredBytes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetDroppedVideoFrames() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.DroppedVideoFrames.Valid() {
		return &vt.DroppedVideoFrames.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetEnvironment() *string {
	if vt := u.OfViewStarted; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.Environment.Valid() {
		return &vt.Environment.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetErrorCategory() *string {
	if vt := u.OfViewStarted; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfViewSnapshot; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfPlaybackError; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfViewEnded; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfRenditionChanged; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfSeekCompleted; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.ErrorCategory.Valid() {
		return (*string)(&vt.ErrorCategory.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetErrorFatal() *bool {
	if vt := u.OfViewStarted; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.ErrorFatal.Valid() {
		return &vt.ErrorFatal.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetErrorPlayheadSeconds() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.ErrorPlayheadSeconds.Valid() {
		return &vt.ErrorPlayheadSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetEventID() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.EventID)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.EventID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetEventName() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.EventName)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.EventName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetFirstVideoFrameTimeMs() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.FirstVideoFrameTimeMs.Valid() {
		return &vt.FirstVideoFrameTimeMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetFurthestProgressRatio() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.FurthestProgressRatio.Valid() {
		return &vt.FurthestProgressRatio.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetMaximumPlayheadSeconds() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.MaximumPlayheadSeconds)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.MaximumPlayheadSeconds)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.MaximumPlayheadSeconds)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.MaximumPlayheadSeconds)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.MaximumPlayheadSeconds.Valid() {
		return &vt.MaximumPlayheadSeconds.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.MaximumPlayheadSeconds.Valid() {
		return &vt.MaximumPlayheadSeconds.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.MaximumPlayheadSeconds.Valid() {
		return &vt.MaximumPlayheadSeconds.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.MaximumPlayheadSeconds.Valid() {
		return &vt.MaximumPlayheadSeconds.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.MaximumPlayheadSeconds.Valid() {
		return &vt.MaximumPlayheadSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetMediaDurationSeconds() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.MediaDurationSeconds.Valid() {
		return &vt.MediaDurationSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetMediaHeightPx() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.MediaHeightPx.Valid() {
		return &vt.MediaHeightPx.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetMediaWidthPx() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.MediaWidthPx.Valid() {
		return &vt.MediaWidthPx.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayerHeightPx() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayerHeightPx.Valid() {
		return &vt.PlayerHeightPx.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayerName() *string {
	if vt := u.OfViewStarted; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayerName.Valid() {
		return &vt.PlayerName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayerVersion() *string {
	if vt := u.OfViewStarted; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayerVersion.Valid() {
		return &vt.PlayerVersion.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayerWidthPx() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayerWidthPx.Valid() {
		return &vt.PlayerWidthPx.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayingTimeMs() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.PlayingTimeMs)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.PlayingTimeMs)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.PlayingTimeMs)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.PlayingTimeMs)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayingTimeMs.Valid() {
		return &vt.PlayingTimeMs.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayingTimeMs.Valid() {
		return &vt.PlayingTimeMs.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.PlayingTimeMs.Valid() {
		return &vt.PlayingTimeMs.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayingTimeMs.Valid() {
		return &vt.PlayingTimeMs.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayingTimeMs.Valid() {
		return &vt.PlayingTimeMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetReason() *string {
	if vt := u.OfViewStarted; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfViewSnapshot; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfPlaybackError; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfViewEnded; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfRenditionChanged; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfSeekCompleted; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.Reason.Valid() {
		return (*string)(&vt.Reason.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetRebufferCount() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.RebufferCount)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.RebufferCount)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.RebufferCount)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.RebufferCount)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.RebufferCount.Valid() {
		return &vt.RebufferCount.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.RebufferCount.Valid() {
		return &vt.RebufferCount.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.RebufferCount.Valid() {
		return &vt.RebufferCount.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.RebufferCount.Valid() {
		return &vt.RebufferCount.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.RebufferCount.Valid() {
		return &vt.RebufferCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetRebufferDurationMs() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.RebufferDurationMs)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.RebufferDurationMs)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.RebufferDurationMs)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.RebufferDurationMs)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.RebufferDurationMs.Valid() {
		return &vt.RebufferDurationMs.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.RebufferDurationMs.Valid() {
		return &vt.RebufferDurationMs.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.RebufferDurationMs.Valid() {
		return &vt.RebufferDurationMs.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.RebufferDurationMs.Valid() {
		return &vt.RebufferDurationMs.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.RebufferDurationMs.Valid() {
		return &vt.RebufferDurationMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetRebufferRatio() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.RebufferRatio.Valid() {
		return &vt.RebufferRatio.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetRenditionSwitchCount() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.RenditionSwitchCount.Valid() {
		return &vt.RenditionSwitchCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSchemaVersion() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*float64)(&vt.SchemaVersion)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSDKName() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.SDKName)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.SDKName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSDKVersion() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.SDKVersion)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.SDKVersion)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekCount() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.SeekCount)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.SeekCount)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.SeekCount)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.SeekCount)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SeekCount.Valid() {
		return &vt.SeekCount.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SeekCount.Valid() {
		return &vt.SeekCount.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SeekCount.Valid() {
		return &vt.SeekCount.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SeekCount.Valid() {
		return &vt.SeekCount.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SeekCount.Valid() {
		return &vt.SeekCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSeekDurationMs() *float64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*float64)(&vt.SeekDurationMs)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*float64)(&vt.SeekDurationMs)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*float64)(&vt.SeekDurationMs)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*float64)(&vt.SeekDurationMs)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SeekDurationMs.Valid() {
		return &vt.SeekDurationMs.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SeekDurationMs.Valid() {
		return &vt.SeekDurationMs.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SeekDurationMs.Valid() {
		return &vt.SeekDurationMs.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SeekDurationMs.Valid() {
		return &vt.SeekDurationMs.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SeekDurationMs.Valid() {
		return &vt.SeekDurationMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSelectedRenditionBitrateBps() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SelectedRenditionBitrateBps.Valid() {
		return &vt.SelectedRenditionBitrateBps.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSelectedRenditionCodec() *string {
	if vt := u.OfViewStarted; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SelectedRenditionCodec.Valid() {
		return &vt.SelectedRenditionCodec.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSelectedRenditionHeight() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SelectedRenditionHeight.Valid() {
		return &vt.SelectedRenditionHeight.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSelectedRenditionWidth() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.SelectedRenditionWidth.Valid() {
		return &vt.SelectedRenditionWidth.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSequence() *int64 {
	if vt := u.OfViewStarted; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*int64)(&vt.Sequence)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*int64)(&vt.Sequence)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetStartupOutcome() *string {
	if vt := u.OfViewStarted; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfViewSnapshot; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfPlaybackError; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfViewEnded; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfRenditionChanged; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfSeekCompleted; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.StartupOutcome.Valid() {
		return (*string)(&vt.StartupOutcome.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetStartupTimeMs() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.StartupTimeMs.Valid() {
		return &vt.StartupTimeMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetTotalVideoFrames() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.TotalVideoFrames.Valid() {
		return &vt.TotalVideoFrames.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetVideoID() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.VideoID)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.VideoID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetViewID() *string {
	if vt := u.OfViewStarted; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfViewSnapshot; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfPlaybackError; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfViewEnded; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfRenditionChanged; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfSeekCompleted; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return (*string)(&vt.ViewID)
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return (*string)(&vt.ViewID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetViewerID() *string {
	if vt := u.OfViewStarted; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.ViewerID.Valid() {
		return &vt.ViewerID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlaybackRate() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.PlaybackRate.Valid() {
		return &vt.PlaybackRate.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlaybackRate.Valid() {
		return &vt.PlaybackRate.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlaybackRate.Valid() {
		return &vt.PlaybackRate.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlaybackRate.Valid() {
		return &vt.PlaybackRate.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlaybackRate.Valid() {
		return &vt.PlaybackRate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetPlayheadSeconds() *float64 {
	if vt := u.OfViewStarted; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.PlayheadSeconds.Valid() {
		return &vt.PlayheadSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetSubtitleLanguage() *string {
	if vt := u.OfViewStarted; vt != nil && vt.SubtitleLanguage.Valid() {
		return &vt.SubtitleLanguage.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.SubtitleLanguage.Valid() {
		return &vt.SubtitleLanguage.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.SubtitleLanguage.Valid() {
		return &vt.SubtitleLanguage.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.SubtitleLanguage.Valid() {
		return &vt.SubtitleLanguage.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.SubtitleLanguage.Valid() {
		return &vt.SubtitleLanguage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetUserID() *string {
	if vt := u.OfViewStarted; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetVideoTitle() *string {
	if vt := u.OfViewStarted; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfViewSnapshot; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfPlaybackError; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfViewEnded; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfRenditionChanged; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfSeekCompleted; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.VideoTitle.Valid() {
		return &vt.VideoTitle.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnalyticsEventUnionParam) GetVisibilityState() *string {
	if vt := u.OfViewStarted; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfViewSnapshot; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfPlaybackError; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfViewEnded; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfPlaybackStateChanged; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfRenditionChanged; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfSeekCompleted; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfSubtitleChanged; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	} else if vt := u.OfPlaybackRateChanged; vt != nil && vt.VisibilityState.Valid() {
		return (*string)(&vt.VisibilityState.Value)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u AnalyticsEventUnionParam) GetErrorCode() (res analyticsEventUnionParamErrorCode) {
	if vt := u.OfViewStarted; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfViewSnapshot; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfPlaybackError; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfViewEnded; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfRenditionChanged; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfSeekCompleted; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfSubtitleChanged; vt != nil {
		res.any = vt.ErrorCode.asAny()
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		res.any = vt.ErrorCode.asAny()
	}
	return
}

// Can have the runtime types [*string], [*float64]
type analyticsEventUnionParamErrorCode struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u analyticsEventUnionParamErrorCode) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's OccurredAt property, if present.
func (u AnalyticsEventUnionParam) GetOccurredAt() *time.Time {
	if vt := u.OfViewStarted; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfViewSnapshot; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfPlaybackError; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfViewEnded; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfRenditionChanged; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfSeekCompleted; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return &vt.OccurredAt
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return &vt.OccurredAt
	}
	return nil
}

// Returns a pointer to the underlying variant's RenditionPlayingTime property, if
// present.
func (u AnalyticsEventUnionParam) GetRenditionPlayingTime() []shared.RenditionPlayingTimeParam {
	if vt := u.OfViewStarted; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfViewSnapshot; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfPlaybackError; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfViewEnded; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfRenditionChanged; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfSeekCompleted; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return vt.RenditionPlayingTime
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return vt.RenditionPlayingTime
	}
	return nil
}

// Returns a pointer to the underlying variant's BrowserContext property, if
// present.
func (u AnalyticsEventUnionParam) GetBrowserContext() *any {
	if vt := u.OfViewStarted; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfViewSnapshot; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfPlaybackError; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfViewEnded; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfPlaybackStateChanged; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfRenditionChanged; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfSeekCompleted; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfSubtitleChanged; vt != nil {
		return &vt.BrowserContext
	} else if vt := u.OfPlaybackRateChanged; vt != nil {
		return &vt.BrowserContext
	}
	return nil
}

// Returns a pointer to the underlying variant's PlaybackRatePlayingTime property,
// if present.
func (u AnalyticsEventUnionParam) GetPlaybackRatePlayingTime() []shared.PlaybackRatePlayingTimeParam {
	if vt := u.OfViewStarted; vt != nil {
		return vt.PlaybackRatePlayingTime
	} else if vt := u.OfViewSnapshot; vt != nil {
		return vt.PlaybackRatePlayingTime
	} else if vt := u.OfPlaybackError; vt != nil {
		return vt.PlaybackRatePlayingTime
	} else if vt := u.OfViewEnded; vt != nil {
		return vt.PlaybackRatePlayingTime
	}
	return nil
}

// Returns a pointer to the underlying variant's PlayedRanges property, if present.
func (u AnalyticsEventUnionParam) GetPlayedRanges() []shared.PlayedRangeParam {
	if vt := u.OfViewStarted; vt != nil {
		return vt.PlayedRanges
	} else if vt := u.OfViewSnapshot; vt != nil {
		return vt.PlayedRanges
	} else if vt := u.OfPlaybackError; vt != nil {
		return vt.PlayedRanges
	} else if vt := u.OfViewEnded; vt != nil {
		return vt.PlayedRanges
	}
	return nil
}

// Returns a pointer to the underlying variant's SubtitlePlayingTime property, if
// present.
func (u AnalyticsEventUnionParam) GetSubtitlePlayingTime() []shared.SubtitlePlayingTimeParam {
	if vt := u.OfViewStarted; vt != nil {
		return vt.SubtitlePlayingTime
	} else if vt := u.OfViewSnapshot; vt != nil {
		return vt.SubtitlePlayingTime
	} else if vt := u.OfPlaybackError; vt != nil {
		return vt.SubtitlePlayingTime
	} else if vt := u.OfViewEnded; vt != nil {
		return vt.SubtitlePlayingTime
	}
	return nil
}

type EventPage struct {
	Events     []StoredEvent `json:"events" api:"required"`
	NextCursor string        `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventPage) RawJSON() string { return r.JSON.raw }
func (r *EventPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NativeCreated struct {
	Accepted  bool `json:"accepted" api:"required"`
	Duplicate bool `json:"duplicate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Duplicate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NativeCreated) RawJSON() string { return r.JSON.raw }
func (r *NativeCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NativeDuplicate struct {
	Accepted  bool `json:"accepted" api:"required"`
	Duplicate bool `json:"duplicate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Duplicate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NativeDuplicate) RawJSON() string { return r.JSON.raw }
func (r *NativeDuplicate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackError struct {
	ConnectorName            string  `json:"connector_name" api:"required"`
	ConnectorVersion         string  `json:"connector_version" api:"required"`
	DeliveryHostname         string  `json:"delivery_hostname" api:"required"`
	DeliveryRequestCount     float64 `json:"delivery_request_count" api:"required"`
	DeliveryThroughputBps    float64 `json:"delivery_throughput_bps" api:"required"`
	DeliveryTransferredBytes float64 `json:"delivery_transferred_bytes" api:"required"`
	DroppedVideoFrames       float64 `json:"dropped_video_frames" api:"required"`
	Environment              string  `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory          PlaybackErrorErrorCategory          `json:"error_category" api:"required"`
	ErrorCode              PlaybackErrorErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal             bool                                `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds   float64                             `json:"error_playhead_seconds" api:"required"`
	EventID                string                              `json:"event_id" api:"required" format:"uuid"`
	EventName              constant.ChunkifyVideoPlaybackError `json:"event_name" default:"chunkify.video.playback_error"`
	FirstVideoFrameTimeMs  float64                             `json:"first_video_frame_time_ms" api:"required"`
	FurthestProgressRatio  float64                             `json:"furthest_progress_ratio" api:"required"`
	MaximumPlayheadSeconds float64                             `json:"maximum_playhead_seconds" api:"required"`
	MediaDurationSeconds   float64                             `json:"media_duration_seconds" api:"required"`
	MediaHeightPx          float64                             `json:"media_height_px" api:"required"`
	MediaWidthPx           float64                             `json:"media_width_px" api:"required"`
	OccurredAt             time.Time                           `json:"occurred_at" api:"required" format:"date-time"`
	PlayerHeightPx         float64                             `json:"player_height_px" api:"required"`
	PlayerName             string                              `json:"player_name" api:"required"`
	PlayerVersion          string                              `json:"player_version" api:"required"`
	PlayerWidthPx          float64                             `json:"player_width_px" api:"required"`
	PlayingTimeMs          float64                             `json:"playing_time_ms" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason               PlaybackErrorReason           `json:"reason" api:"required"`
	RebufferCount        float64                       `json:"rebuffer_count" api:"required"`
	RebufferDurationMs   float64                       `json:"rebuffer_duration_ms" api:"required"`
	RebufferRatio        float64                       `json:"rebuffer_ratio" api:"required"`
	RenditionPlayingTime []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"required"`
	RenditionSwitchCount float64                       `json:"rendition_switch_count" api:"required"`
	SchemaVersion        float64                       `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string  `json:"sdk_name" api:"required"`
	SDKVersion                  string  `json:"sdk_version" api:"required"`
	SeekCount                   float64 `json:"seek_count" api:"required"`
	SeekDurationMs              float64 `json:"seek_duration_ms" api:"required"`
	SelectedRenditionBitrateBps float64 `json:"selected_rendition_bitrate_bps" api:"required"`
	SelectedRenditionCodec      string  `json:"selected_rendition_codec" api:"required"`
	SelectedRenditionHeight     float64 `json:"selected_rendition_height" api:"required"`
	SelectedRenditionWidth      float64 `json:"selected_rendition_width" api:"required"`
	Sequence                    int64   `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   PlaybackErrorStartupOutcome `json:"startup_outcome" api:"required"`
	StartupTimeMs    float64                     `json:"startup_time_ms" api:"required"`
	TotalVideoFrames float64                     `json:"total_video_frames" api:"required"`
	VideoID          string                      `json:"video_id" api:"required"`
	ViewID           string                      `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate float64 `json:"playback_rate" api:"nullable"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTime `json:"playback_rate_playing_time" api:"nullable"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRange `json:"played_ranges" api:"nullable"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds float64 `json:"playhead_seconds" api:"nullable"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage string `json:"subtitle_language" api:"nullable"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTime `json:"subtitle_playing_time" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState PlaybackErrorVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		Reason                      respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		PlaybackRate                respjson.Field
		PlaybackRatePlayingTime     respjson.Field
		PlayedRanges                respjson.Field
		PlayheadSeconds             respjson.Field
		SubtitleLanguage            respjson.Field
		SubtitlePlayingTime         respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackError) RawJSON() string { return r.JSON.raw }
func (r *PlaybackError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlaybackError to a PlaybackErrorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlaybackErrorParam.Overrides()
func (r PlaybackError) ToParam() PlaybackErrorParam {
	return param.Override[PlaybackErrorParam](json.RawMessage(r.RawJSON()))
}

type PlaybackErrorErrorCategory string

const (
	PlaybackErrorErrorCategoryAborted     PlaybackErrorErrorCategory = "aborted"
	PlaybackErrorErrorCategoryNetwork     PlaybackErrorErrorCategory = "network"
	PlaybackErrorErrorCategoryDecode      PlaybackErrorErrorCategory = "decode"
	PlaybackErrorErrorCategoryUnsupported PlaybackErrorErrorCategory = "unsupported"
	PlaybackErrorErrorCategoryUnknown     PlaybackErrorErrorCategory = "unknown"
)

// PlaybackErrorErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type PlaybackErrorErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u PlaybackErrorErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlaybackErrorErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlaybackErrorErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *PlaybackErrorErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackErrorReason string

const (
	PlaybackErrorReasonEnded         PlaybackErrorReason = "ended"
	PlaybackErrorReasonDestroyed     PlaybackErrorReason = "destroyed"
	PlaybackErrorReasonSourceChanged PlaybackErrorReason = "source_changed"
	PlaybackErrorReasonPageExit      PlaybackErrorReason = "page_exit"
)

type PlaybackErrorStartupOutcome string

const (
	PlaybackErrorStartupOutcomeStarted           PlaybackErrorStartupOutcome = "started"
	PlaybackErrorStartupOutcomeFailedBeforeStart PlaybackErrorStartupOutcome = "failed_before_start"
	PlaybackErrorStartupOutcomeEndedBeforeStart  PlaybackErrorStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type PlaybackErrorVisibilityState string

const (
	PlaybackErrorVisibilityStateVisible PlaybackErrorVisibilityState = "visible"
	PlaybackErrorVisibilityStateHidden  PlaybackErrorVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname,
// DeliveryRequestCount, DeliveryThroughputBps, DeliveryTransferredBytes,
// DroppedVideoFrames, Environment, ErrorCategory, ErrorCode, ErrorFatal,
// ErrorPlayheadSeconds, EventID, EventName, FirstVideoFrameTimeMs,
// FurthestProgressRatio, MaximumPlayheadSeconds, MediaDurationSeconds,
// MediaHeightPx, MediaWidthPx, OccurredAt, PlayerHeightPx, PlayerName,
// PlayerVersion, PlayerWidthPx, PlayingTimeMs, Reason, RebufferCount,
// RebufferDurationMs, RebufferRatio, RenditionPlayingTime, RenditionSwitchCount,
// SchemaVersion, SDKName, SDKVersion, SeekCount, SeekDurationMs,
// SelectedRenditionBitrateBps, SelectedRenditionCodec, SelectedRenditionHeight,
// SelectedRenditionWidth, Sequence, StartupOutcome, StartupTimeMs,
// TotalVideoFrames, VideoID, ViewID, ViewerID are required.
type PlaybackErrorParam struct {
	ConnectorVersion         param.Opt[string]  `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname         param.Opt[string]  `json:"delivery_hostname,omitzero" api:"required"`
	DeliveryRequestCount     param.Opt[float64] `json:"delivery_request_count,omitzero" api:"required"`
	DeliveryThroughputBps    param.Opt[float64] `json:"delivery_throughput_bps,omitzero" api:"required"`
	DeliveryTransferredBytes param.Opt[float64] `json:"delivery_transferred_bytes,omitzero" api:"required"`
	DroppedVideoFrames       param.Opt[float64] `json:"dropped_video_frames,omitzero" api:"required"`
	Environment              param.Opt[string]  `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory         param.Opt[PlaybackErrorErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal            param.Opt[bool]                       `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds  param.Opt[float64]                    `json:"error_playhead_seconds,omitzero" api:"required"`
	FirstVideoFrameTimeMs param.Opt[float64]                    `json:"first_video_frame_time_ms,omitzero" api:"required"`
	FurthestProgressRatio param.Opt[float64]                    `json:"furthest_progress_ratio,omitzero" api:"required"`
	MediaDurationSeconds  param.Opt[float64]                    `json:"media_duration_seconds,omitzero" api:"required"`
	MediaHeightPx         param.Opt[float64]                    `json:"media_height_px,omitzero" api:"required"`
	MediaWidthPx          param.Opt[float64]                    `json:"media_width_px,omitzero" api:"required"`
	PlayerHeightPx        param.Opt[float64]                    `json:"player_height_px,omitzero" api:"required"`
	PlayerName            param.Opt[string]                     `json:"player_name,omitzero" api:"required"`
	PlayerVersion         param.Opt[string]                     `json:"player_version,omitzero" api:"required"`
	PlayerWidthPx         param.Opt[float64]                    `json:"player_width_px,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason                      param.Opt[PlaybackErrorReason] `json:"reason,omitzero" api:"required"`
	RebufferRatio               param.Opt[float64]             `json:"rebuffer_ratio,omitzero" api:"required"`
	RenditionSwitchCount        param.Opt[float64]             `json:"rendition_switch_count,omitzero" api:"required"`
	SelectedRenditionBitrateBps param.Opt[float64]             `json:"selected_rendition_bitrate_bps,omitzero" api:"required"`
	SelectedRenditionCodec      param.Opt[string]              `json:"selected_rendition_codec,omitzero" api:"required"`
	SelectedRenditionHeight     param.Opt[float64]             `json:"selected_rendition_height,omitzero" api:"required"`
	SelectedRenditionWidth      param.Opt[float64]             `json:"selected_rendition_width,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   param.Opt[PlaybackErrorStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	StartupTimeMs    param.Opt[float64]                     `json:"startup_time_ms,omitzero" api:"required"`
	TotalVideoFrames param.Opt[float64]                     `json:"total_video_frames,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID               param.Opt[string]                  `json:"viewer_id,omitzero" api:"required"`
	ErrorCode              PlaybackErrorErrorCodeUnionParam   `json:"error_code,omitzero" api:"required"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam `json:"rendition_playing_time,omitzero" api:"required"`
	ConnectorName          string                             `json:"connector_name" api:"required"`
	EventID                string                             `json:"event_id" api:"required" format:"uuid"`
	MaximumPlayheadSeconds float64                            `json:"maximum_playhead_seconds" api:"required"`
	OccurredAt             time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	PlayingTimeMs          float64                            `json:"playing_time_ms" api:"required"`
	RebufferCount          float64                            `json:"rebuffer_count" api:"required"`
	RebufferDurationMs     float64                            `json:"rebuffer_duration_ms" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName        string  `json:"sdk_name" api:"required"`
	SDKVersion     string  `json:"sdk_version" api:"required"`
	SeekCount      float64 `json:"seek_count" api:"required"`
	SeekDurationMs float64 `json:"seek_duration_ms" api:"required"`
	Sequence       int64   `json:"sequence" api:"required"`
	VideoID        string  `json:"video_id" api:"required"`
	ViewID         string  `json:"view_id" api:"required" format:"uuid"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate param.Opt[float64] `json:"playback_rate,omitzero"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds param.Opt[float64] `json:"playhead_seconds,omitzero"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage param.Opt[string] `json:"subtitle_language,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState param.Opt[PlaybackErrorVisibilityState] `json:"visibility_state,omitzero"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTimeParam `json:"playback_rate_playing_time,omitzero"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRangeParam `json:"played_ranges,omitzero"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTimeParam `json:"subtitle_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.playback_error".
	EventName constant.ChunkifyVideoPlaybackError `json:"event_name" default:"chunkify.video.playback_error"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r PlaybackErrorParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaybackErrorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaybackErrorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlaybackErrorParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PlaybackErrorErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u PlaybackErrorErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *PlaybackErrorErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PlaybackErrorErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type PlaybackRateChanged struct {
	ConnectorName    string `json:"connector_name" api:"required"`
	ConnectorVersion string `json:"connector_version" api:"required"`
	DeliveryHostname string `json:"delivery_hostname" api:"required"`
	Environment      string `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        PlaybackRateChangedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode            PlaybackRateChangedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal           bool                                      `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds float64                                   `json:"error_playhead_seconds" api:"required"`
	EventID              string                                    `json:"event_id" api:"required" format:"uuid"`
	EventName            constant.ChunkifyVideoPlaybackRateChanged `json:"event_name" default:"chunkify.video.playback_rate_changed"`
	OccurredAt           time.Time                                 `json:"occurred_at" api:"required" format:"date-time"`
	// Rate after the change.
	PlaybackRate    float64 `json:"playback_rate" api:"required"`
	PlayerName      string  `json:"player_name" api:"required"`
	PlayerVersion   string  `json:"player_version" api:"required"`
	PlayheadSeconds float64 `json:"playhead_seconds" api:"required"`
	// Rate before the change.
	PreviousPlaybackRate float64 `json:"previous_playback_rate" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason        PlaybackRateChangedReason `json:"reason" api:"required"`
	SchemaVersion float64                   `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName    string `json:"sdk_name" api:"required"`
	SDKVersion string `json:"sdk_version" api:"required"`
	Sequence   int64  `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome PlaybackRateChangedStartupOutcome `json:"startup_outcome" api:"required"`
	VideoID        string                            `json:"video_id" api:"required"`
	ViewID         string                            `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext              any                           `json:"browser_context"`
	DeliveryRequestCount        float64                       `json:"delivery_request_count" api:"nullable"`
	DeliveryThroughputBps       float64                       `json:"delivery_throughput_bps" api:"nullable"`
	DeliveryTransferredBytes    float64                       `json:"delivery_transferred_bytes" api:"nullable"`
	DroppedVideoFrames          float64                       `json:"dropped_video_frames" api:"nullable"`
	FirstVideoFrameTimeMs       float64                       `json:"first_video_frame_time_ms" api:"nullable"`
	FurthestProgressRatio       float64                       `json:"furthest_progress_ratio" api:"nullable"`
	MaximumPlayheadSeconds      float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds        float64                       `json:"media_duration_seconds" api:"nullable"`
	MediaHeightPx               float64                       `json:"media_height_px" api:"nullable"`
	MediaWidthPx                float64                       `json:"media_width_px" api:"nullable"`
	PlayerHeightPx              float64                       `json:"player_height_px" api:"nullable"`
	PlayerWidthPx               float64                       `json:"player_width_px" api:"nullable"`
	PlayingTimeMs               float64                       `json:"playing_time_ms"`
	RebufferCount               float64                       `json:"rebuffer_count"`
	RebufferDurationMs          float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio               float64                       `json:"rebuffer_ratio" api:"nullable"`
	RenditionPlayingTime        []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"nullable"`
	RenditionSwitchCount        float64                       `json:"rendition_switch_count" api:"nullable"`
	SeekCount                   float64                       `json:"seek_count"`
	SeekDurationMs              float64                       `json:"seek_duration_ms"`
	SelectedRenditionBitrateBps float64                       `json:"selected_rendition_bitrate_bps" api:"nullable"`
	SelectedRenditionCodec      string                        `json:"selected_rendition_codec" api:"nullable"`
	SelectedRenditionHeight     float64                       `json:"selected_rendition_height" api:"nullable"`
	SelectedRenditionWidth      float64                       `json:"selected_rendition_width" api:"nullable"`
	StartupTimeMs               float64                       `json:"startup_time_ms" api:"nullable"`
	TotalVideoFrames            float64                       `json:"total_video_frames" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState PlaybackRateChangedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		OccurredAt                  respjson.Field
		PlaybackRate                respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayheadSeconds             respjson.Field
		PreviousPlaybackRate        respjson.Field
		Reason                      respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackRateChanged) RawJSON() string { return r.JSON.raw }
func (r *PlaybackRateChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlaybackRateChanged to a PlaybackRateChangedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlaybackRateChangedParam.Overrides()
func (r PlaybackRateChanged) ToParam() PlaybackRateChangedParam {
	return param.Override[PlaybackRateChangedParam](json.RawMessage(r.RawJSON()))
}

type PlaybackRateChangedErrorCategory string

const (
	PlaybackRateChangedErrorCategoryAborted     PlaybackRateChangedErrorCategory = "aborted"
	PlaybackRateChangedErrorCategoryNetwork     PlaybackRateChangedErrorCategory = "network"
	PlaybackRateChangedErrorCategoryDecode      PlaybackRateChangedErrorCategory = "decode"
	PlaybackRateChangedErrorCategoryUnsupported PlaybackRateChangedErrorCategory = "unsupported"
	PlaybackRateChangedErrorCategoryUnknown     PlaybackRateChangedErrorCategory = "unknown"
)

// PlaybackRateChangedErrorCodeUnion contains all possible properties and values
// from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type PlaybackRateChangedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u PlaybackRateChangedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlaybackRateChangedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlaybackRateChangedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *PlaybackRateChangedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackRateChangedReason string

const (
	PlaybackRateChangedReasonEnded         PlaybackRateChangedReason = "ended"
	PlaybackRateChangedReasonDestroyed     PlaybackRateChangedReason = "destroyed"
	PlaybackRateChangedReasonSourceChanged PlaybackRateChangedReason = "source_changed"
	PlaybackRateChangedReasonPageExit      PlaybackRateChangedReason = "page_exit"
)

type PlaybackRateChangedStartupOutcome string

const (
	PlaybackRateChangedStartupOutcomeStarted           PlaybackRateChangedStartupOutcome = "started"
	PlaybackRateChangedStartupOutcomeFailedBeforeStart PlaybackRateChangedStartupOutcome = "failed_before_start"
	PlaybackRateChangedStartupOutcomeEndedBeforeStart  PlaybackRateChangedStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type PlaybackRateChangedVisibilityState string

const (
	PlaybackRateChangedVisibilityStateVisible PlaybackRateChangedVisibilityState = "visible"
	PlaybackRateChangedVisibilityStateHidden  PlaybackRateChangedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname, Environment,
// ErrorCategory, ErrorCode, ErrorFatal, ErrorPlayheadSeconds, EventID, EventName,
// OccurredAt, PlaybackRate, PlayerName, PlayerVersion, PlayheadSeconds,
// PreviousPlaybackRate, Reason, SchemaVersion, SDKName, SDKVersion, Sequence,
// StartupOutcome, VideoID, ViewID, ViewerID are required.
type PlaybackRateChangedParam struct {
	ConnectorVersion param.Opt[string] `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname param.Opt[string] `json:"delivery_hostname,omitzero" api:"required"`
	Environment      param.Opt[string] `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        param.Opt[PlaybackRateChangedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal           param.Opt[bool]                             `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds param.Opt[float64]                          `json:"error_playhead_seconds,omitzero" api:"required"`
	// Rate after the change.
	PlaybackRate    param.Opt[float64] `json:"playback_rate,omitzero" api:"required"`
	PlayerName      param.Opt[string]  `json:"player_name,omitzero" api:"required"`
	PlayerVersion   param.Opt[string]  `json:"player_version,omitzero" api:"required"`
	PlayheadSeconds param.Opt[float64] `json:"playhead_seconds,omitzero" api:"required"`
	// Rate before the change.
	PreviousPlaybackRate param.Opt[float64] `json:"previous_playback_rate,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason param.Opt[PlaybackRateChangedReason] `json:"reason,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome param.Opt[PlaybackRateChangedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID      param.Opt[string]                      `json:"viewer_id,omitzero" api:"required"`
	ErrorCode     PlaybackRateChangedErrorCodeUnionParam `json:"error_code,omitzero" api:"required"`
	ConnectorName string                                 `json:"connector_name" api:"required"`
	EventID       string                                 `json:"event_id" api:"required" format:"uuid"`
	OccurredAt    time.Time                              `json:"occurred_at" api:"required" format:"date-time"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string             `json:"sdk_name" api:"required"`
	SDKVersion                  string             `json:"sdk_version" api:"required"`
	Sequence                    int64              `json:"sequence" api:"required"`
	VideoID                     string             `json:"video_id" api:"required"`
	ViewID                      string             `json:"view_id" api:"required" format:"uuid"`
	DeliveryRequestCount        param.Opt[float64] `json:"delivery_request_count,omitzero"`
	DeliveryThroughputBps       param.Opt[float64] `json:"delivery_throughput_bps,omitzero"`
	DeliveryTransferredBytes    param.Opt[float64] `json:"delivery_transferred_bytes,omitzero"`
	DroppedVideoFrames          param.Opt[float64] `json:"dropped_video_frames,omitzero"`
	FirstVideoFrameTimeMs       param.Opt[float64] `json:"first_video_frame_time_ms,omitzero"`
	FurthestProgressRatio       param.Opt[float64] `json:"furthest_progress_ratio,omitzero"`
	MediaDurationSeconds        param.Opt[float64] `json:"media_duration_seconds,omitzero"`
	MediaHeightPx               param.Opt[float64] `json:"media_height_px,omitzero"`
	MediaWidthPx                param.Opt[float64] `json:"media_width_px,omitzero"`
	PlayerHeightPx              param.Opt[float64] `json:"player_height_px,omitzero"`
	PlayerWidthPx               param.Opt[float64] `json:"player_width_px,omitzero"`
	RebufferRatio               param.Opt[float64] `json:"rebuffer_ratio,omitzero"`
	RenditionSwitchCount        param.Opt[float64] `json:"rendition_switch_count,omitzero"`
	SelectedRenditionBitrateBps param.Opt[float64] `json:"selected_rendition_bitrate_bps,omitzero"`
	SelectedRenditionCodec      param.Opt[string]  `json:"selected_rendition_codec,omitzero"`
	SelectedRenditionHeight     param.Opt[float64] `json:"selected_rendition_height,omitzero"`
	SelectedRenditionWidth      param.Opt[float64] `json:"selected_rendition_width,omitzero"`
	StartupTimeMs               param.Opt[float64] `json:"startup_time_ms,omitzero"`
	TotalVideoFrames            param.Opt[float64] `json:"total_video_frames,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState        param.Opt[PlaybackRateChangedVisibilityState] `json:"visibility_state,omitzero"`
	MaximumPlayheadSeconds param.Opt[float64]                            `json:"maximum_playhead_seconds,omitzero"`
	PlayingTimeMs          param.Opt[float64]                            `json:"playing_time_ms,omitzero"`
	RebufferCount          param.Opt[float64]                            `json:"rebuffer_count,omitzero"`
	RebufferDurationMs     param.Opt[float64]                            `json:"rebuffer_duration_ms,omitzero"`
	SeekCount              param.Opt[float64]                            `json:"seek_count,omitzero"`
	SeekDurationMs         param.Opt[float64]                            `json:"seek_duration_ms,omitzero"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam            `json:"rendition_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.playback_rate_changed".
	EventName constant.ChunkifyVideoPlaybackRateChanged `json:"event_name" default:"chunkify.video.playback_rate_changed"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r PlaybackRateChangedParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaybackRateChangedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaybackRateChangedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlaybackRateChangedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PlaybackRateChangedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u PlaybackRateChangedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *PlaybackRateChangedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PlaybackRateChangedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type PlaybackStateChanged struct {
	ConnectorName    string `json:"connector_name" api:"required"`
	ConnectorVersion string `json:"connector_version" api:"required"`
	DeliveryHostname string `json:"delivery_hostname" api:"required"`
	Environment      string `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        PlaybackStateChangedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode            PlaybackStateChangedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal           bool                                       `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds float64                                    `json:"error_playhead_seconds" api:"required"`
	EventID              string                                     `json:"event_id" api:"required" format:"uuid"`
	EventName            constant.ChunkifyVideoPlaybackStateChanged `json:"event_name" default:"chunkify.video.playback_state_changed"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	FromState               PlaybackStateChangedFromState `json:"from_state" api:"required"`
	OccurredAt              time.Time                     `json:"occurred_at" api:"required" format:"date-time"`
	PlayerName              string                        `json:"player_name" api:"required"`
	PlayerVersion           string                        `json:"player_version" api:"required"`
	PlayheadSeconds         float64                       `json:"playhead_seconds" api:"required"`
	PreviousStateDurationMs float64                       `json:"previous_state_duration_ms" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason        PlaybackStateChangedReason `json:"reason" api:"required"`
	SchemaVersion float64                    `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName    string `json:"sdk_name" api:"required"`
	SDKVersion string `json:"sdk_version" api:"required"`
	Sequence   int64  `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome PlaybackStateChangedStartupOutcome `json:"startup_outcome" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	ToState PlaybackStateChangedToState `json:"to_state" api:"required"`
	VideoID string                      `json:"video_id" api:"required"`
	ViewID  string                      `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext              any                           `json:"browser_context"`
	DeliveryRequestCount        float64                       `json:"delivery_request_count" api:"nullable"`
	DeliveryThroughputBps       float64                       `json:"delivery_throughput_bps" api:"nullable"`
	DeliveryTransferredBytes    float64                       `json:"delivery_transferred_bytes" api:"nullable"`
	DroppedVideoFrames          float64                       `json:"dropped_video_frames" api:"nullable"`
	FirstVideoFrameTimeMs       float64                       `json:"first_video_frame_time_ms" api:"nullable"`
	FurthestProgressRatio       float64                       `json:"furthest_progress_ratio" api:"nullable"`
	MaximumPlayheadSeconds      float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds        float64                       `json:"media_duration_seconds" api:"nullable"`
	MediaHeightPx               float64                       `json:"media_height_px" api:"nullable"`
	MediaWidthPx                float64                       `json:"media_width_px" api:"nullable"`
	PlayerHeightPx              float64                       `json:"player_height_px" api:"nullable"`
	PlayerWidthPx               float64                       `json:"player_width_px" api:"nullable"`
	PlayingTimeMs               float64                       `json:"playing_time_ms"`
	RebufferCount               float64                       `json:"rebuffer_count"`
	RebufferDurationMs          float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio               float64                       `json:"rebuffer_ratio" api:"nullable"`
	RenditionPlayingTime        []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"nullable"`
	RenditionSwitchCount        float64                       `json:"rendition_switch_count" api:"nullable"`
	SeekCount                   float64                       `json:"seek_count"`
	SeekDurationMs              float64                       `json:"seek_duration_ms"`
	SelectedRenditionBitrateBps float64                       `json:"selected_rendition_bitrate_bps" api:"nullable"`
	SelectedRenditionCodec      string                        `json:"selected_rendition_codec" api:"nullable"`
	SelectedRenditionHeight     float64                       `json:"selected_rendition_height" api:"nullable"`
	SelectedRenditionWidth      float64                       `json:"selected_rendition_width" api:"nullable"`
	StartupTimeMs               float64                       `json:"startup_time_ms" api:"nullable"`
	TotalVideoFrames            float64                       `json:"total_video_frames" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState PlaybackStateChangedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FromState                   respjson.Field
		OccurredAt                  respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayheadSeconds             respjson.Field
		PreviousStateDurationMs     respjson.Field
		Reason                      respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		ToState                     respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackStateChanged) RawJSON() string { return r.JSON.raw }
func (r *PlaybackStateChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlaybackStateChanged to a PlaybackStateChangedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlaybackStateChangedParam.Overrides()
func (r PlaybackStateChanged) ToParam() PlaybackStateChangedParam {
	return param.Override[PlaybackStateChangedParam](json.RawMessage(r.RawJSON()))
}

type PlaybackStateChangedErrorCategory string

const (
	PlaybackStateChangedErrorCategoryAborted     PlaybackStateChangedErrorCategory = "aborted"
	PlaybackStateChangedErrorCategoryNetwork     PlaybackStateChangedErrorCategory = "network"
	PlaybackStateChangedErrorCategoryDecode      PlaybackStateChangedErrorCategory = "decode"
	PlaybackStateChangedErrorCategoryUnsupported PlaybackStateChangedErrorCategory = "unsupported"
	PlaybackStateChangedErrorCategoryUnknown     PlaybackStateChangedErrorCategory = "unknown"
)

// PlaybackStateChangedErrorCodeUnion contains all possible properties and values
// from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type PlaybackStateChangedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u PlaybackStateChangedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlaybackStateChangedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlaybackStateChangedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *PlaybackStateChangedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackStateChangedFromState string

const (
	PlaybackStateChangedFromStateStartup   PlaybackStateChangedFromState = "startup"
	PlaybackStateChangedFromStateResuming  PlaybackStateChangedFromState = "resuming"
	PlaybackStateChangedFromStatePlaying   PlaybackStateChangedFromState = "playing"
	PlaybackStateChangedFromStatePaused    PlaybackStateChangedFromState = "paused"
	PlaybackStateChangedFromStateBuffering PlaybackStateChangedFromState = "buffering"
	PlaybackStateChangedFromStateSeeking   PlaybackStateChangedFromState = "seeking"
	PlaybackStateChangedFromStateError     PlaybackStateChangedFromState = "error"
	PlaybackStateChangedFromStateEnded     PlaybackStateChangedFromState = "ended"
	PlaybackStateChangedFromStateStopped   PlaybackStateChangedFromState = "stopped"
)

type PlaybackStateChangedReason string

const (
	PlaybackStateChangedReasonEnded         PlaybackStateChangedReason = "ended"
	PlaybackStateChangedReasonDestroyed     PlaybackStateChangedReason = "destroyed"
	PlaybackStateChangedReasonSourceChanged PlaybackStateChangedReason = "source_changed"
	PlaybackStateChangedReasonPageExit      PlaybackStateChangedReason = "page_exit"
)

type PlaybackStateChangedStartupOutcome string

const (
	PlaybackStateChangedStartupOutcomeStarted           PlaybackStateChangedStartupOutcome = "started"
	PlaybackStateChangedStartupOutcomeFailedBeforeStart PlaybackStateChangedStartupOutcome = "failed_before_start"
	PlaybackStateChangedStartupOutcomeEndedBeforeStart  PlaybackStateChangedStartupOutcome = "ended_before_start"
)

type PlaybackStateChangedToState string

const (
	PlaybackStateChangedToStateStartup   PlaybackStateChangedToState = "startup"
	PlaybackStateChangedToStateResuming  PlaybackStateChangedToState = "resuming"
	PlaybackStateChangedToStatePlaying   PlaybackStateChangedToState = "playing"
	PlaybackStateChangedToStatePaused    PlaybackStateChangedToState = "paused"
	PlaybackStateChangedToStateBuffering PlaybackStateChangedToState = "buffering"
	PlaybackStateChangedToStateSeeking   PlaybackStateChangedToState = "seeking"
	PlaybackStateChangedToStateError     PlaybackStateChangedToState = "error"
	PlaybackStateChangedToStateEnded     PlaybackStateChangedToState = "ended"
	PlaybackStateChangedToStateStopped   PlaybackStateChangedToState = "stopped"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type PlaybackStateChangedVisibilityState string

const (
	PlaybackStateChangedVisibilityStateVisible PlaybackStateChangedVisibilityState = "visible"
	PlaybackStateChangedVisibilityStateHidden  PlaybackStateChangedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname, Environment,
// ErrorCategory, ErrorCode, ErrorFatal, ErrorPlayheadSeconds, EventID, EventName,
// FromState, OccurredAt, PlayerName, PlayerVersion, PlayheadSeconds,
// PreviousStateDurationMs, Reason, SchemaVersion, SDKName, SDKVersion, Sequence,
// StartupOutcome, ToState, VideoID, ViewID, ViewerID are required.
type PlaybackStateChangedParam struct {
	ConnectorVersion param.Opt[string] `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname param.Opt[string] `json:"delivery_hostname,omitzero" api:"required"`
	Environment      param.Opt[string] `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        param.Opt[PlaybackStateChangedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal           param.Opt[bool]                              `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds param.Opt[float64]                           `json:"error_playhead_seconds,omitzero" api:"required"`
	PlayerName           param.Opt[string]                            `json:"player_name,omitzero" api:"required"`
	PlayerVersion        param.Opt[string]                            `json:"player_version,omitzero" api:"required"`
	PlayheadSeconds      param.Opt[float64]                           `json:"playhead_seconds,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason param.Opt[PlaybackStateChangedReason] `json:"reason,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome param.Opt[PlaybackStateChangedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID      param.Opt[string]                       `json:"viewer_id,omitzero" api:"required"`
	ErrorCode     PlaybackStateChangedErrorCodeUnionParam `json:"error_code,omitzero" api:"required"`
	ConnectorName string                                  `json:"connector_name" api:"required"`
	EventID       string                                  `json:"event_id" api:"required" format:"uuid"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	FromState               PlaybackStateChangedFromState `json:"from_state,omitzero" api:"required"`
	OccurredAt              time.Time                     `json:"occurred_at" api:"required" format:"date-time"`
	PreviousStateDurationMs float64                       `json:"previous_state_duration_ms" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName    string `json:"sdk_name" api:"required"`
	SDKVersion string `json:"sdk_version" api:"required"`
	Sequence   int64  `json:"sequence" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	ToState                     PlaybackStateChangedToState `json:"to_state,omitzero" api:"required"`
	VideoID                     string                      `json:"video_id" api:"required"`
	ViewID                      string                      `json:"view_id" api:"required" format:"uuid"`
	DeliveryRequestCount        param.Opt[float64]          `json:"delivery_request_count,omitzero"`
	DeliveryThroughputBps       param.Opt[float64]          `json:"delivery_throughput_bps,omitzero"`
	DeliveryTransferredBytes    param.Opt[float64]          `json:"delivery_transferred_bytes,omitzero"`
	DroppedVideoFrames          param.Opt[float64]          `json:"dropped_video_frames,omitzero"`
	FirstVideoFrameTimeMs       param.Opt[float64]          `json:"first_video_frame_time_ms,omitzero"`
	FurthestProgressRatio       param.Opt[float64]          `json:"furthest_progress_ratio,omitzero"`
	MediaDurationSeconds        param.Opt[float64]          `json:"media_duration_seconds,omitzero"`
	MediaHeightPx               param.Opt[float64]          `json:"media_height_px,omitzero"`
	MediaWidthPx                param.Opt[float64]          `json:"media_width_px,omitzero"`
	PlayerHeightPx              param.Opt[float64]          `json:"player_height_px,omitzero"`
	PlayerWidthPx               param.Opt[float64]          `json:"player_width_px,omitzero"`
	RebufferRatio               param.Opt[float64]          `json:"rebuffer_ratio,omitzero"`
	RenditionSwitchCount        param.Opt[float64]          `json:"rendition_switch_count,omitzero"`
	SelectedRenditionBitrateBps param.Opt[float64]          `json:"selected_rendition_bitrate_bps,omitzero"`
	SelectedRenditionCodec      param.Opt[string]           `json:"selected_rendition_codec,omitzero"`
	SelectedRenditionHeight     param.Opt[float64]          `json:"selected_rendition_height,omitzero"`
	SelectedRenditionWidth      param.Opt[float64]          `json:"selected_rendition_width,omitzero"`
	StartupTimeMs               param.Opt[float64]          `json:"startup_time_ms,omitzero"`
	TotalVideoFrames            param.Opt[float64]          `json:"total_video_frames,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState        param.Opt[PlaybackStateChangedVisibilityState] `json:"visibility_state,omitzero"`
	MaximumPlayheadSeconds param.Opt[float64]                             `json:"maximum_playhead_seconds,omitzero"`
	PlayingTimeMs          param.Opt[float64]                             `json:"playing_time_ms,omitzero"`
	RebufferCount          param.Opt[float64]                             `json:"rebuffer_count,omitzero"`
	RebufferDurationMs     param.Opt[float64]                             `json:"rebuffer_duration_ms,omitzero"`
	SeekCount              param.Opt[float64]                             `json:"seek_count,omitzero"`
	SeekDurationMs         param.Opt[float64]                             `json:"seek_duration_ms,omitzero"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam             `json:"rendition_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.playback_state_changed".
	EventName constant.ChunkifyVideoPlaybackStateChanged `json:"event_name" default:"chunkify.video.playback_state_changed"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r PlaybackStateChangedParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaybackStateChangedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaybackStateChangedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlaybackStateChangedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PlaybackStateChangedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u PlaybackStateChangedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *PlaybackStateChangedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PlaybackStateChangedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type RenditionChanged struct {
	ConnectorName    string `json:"connector_name" api:"required"`
	ConnectorVersion string `json:"connector_version" api:"required"`
	DeliveryHostname string `json:"delivery_hostname" api:"required"`
	Environment      string `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory               RenditionChangedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode                   RenditionChangedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal                  bool                                   `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds        float64                                `json:"error_playhead_seconds" api:"required"`
	EventID                     string                                 `json:"event_id" api:"required" format:"uuid"`
	EventName                   constant.ChunkifyVideoRenditionChanged `json:"event_name" default:"chunkify.video.rendition_changed"`
	OccurredAt                  time.Time                              `json:"occurred_at" api:"required" format:"date-time"`
	PlayerHeightPx              float64                                `json:"player_height_px" api:"required"`
	PlayerName                  string                                 `json:"player_name" api:"required"`
	PlayerVersion               string                                 `json:"player_version" api:"required"`
	PlayerWidthPx               float64                                `json:"player_width_px" api:"required"`
	PlayheadSeconds             float64                                `json:"playhead_seconds" api:"required"`
	PreviousRenditionBitrateBps float64                                `json:"previous_rendition_bitrate_bps" api:"required"`
	PreviousRenditionCodec      string                                 `json:"previous_rendition_codec" api:"required"`
	PreviousRenditionFrameRate  float64                                `json:"previous_rendition_frame_rate" api:"required"`
	PreviousRenditionHeight     float64                                `json:"previous_rendition_height" api:"required"`
	PreviousRenditionWidth      float64                                `json:"previous_rendition_width" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason        RenditionChangedReason `json:"reason" api:"required"`
	SchemaVersion float64                `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string  `json:"sdk_name" api:"required"`
	SDKVersion                  string  `json:"sdk_version" api:"required"`
	SelectedRenditionBitrateBps float64 `json:"selected_rendition_bitrate_bps" api:"required"`
	SelectedRenditionCodec      string  `json:"selected_rendition_codec" api:"required"`
	SelectedRenditionFrameRate  float64 `json:"selected_rendition_frame_rate" api:"required"`
	SelectedRenditionHeight     float64 `json:"selected_rendition_height" api:"required"`
	SelectedRenditionWidth      float64 `json:"selected_rendition_width" api:"required"`
	Sequence                    int64   `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome RenditionChangedStartupOutcome `json:"startup_outcome" api:"required"`
	VideoID        string                         `json:"video_id" api:"required"`
	ViewID         string                         `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext           any                           `json:"browser_context"`
	DeliveryRequestCount     float64                       `json:"delivery_request_count" api:"nullable"`
	DeliveryThroughputBps    float64                       `json:"delivery_throughput_bps" api:"nullable"`
	DeliveryTransferredBytes float64                       `json:"delivery_transferred_bytes" api:"nullable"`
	DroppedVideoFrames       float64                       `json:"dropped_video_frames" api:"nullable"`
	FirstVideoFrameTimeMs    float64                       `json:"first_video_frame_time_ms" api:"nullable"`
	FurthestProgressRatio    float64                       `json:"furthest_progress_ratio" api:"nullable"`
	MaximumPlayheadSeconds   float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds     float64                       `json:"media_duration_seconds" api:"nullable"`
	MediaHeightPx            float64                       `json:"media_height_px" api:"nullable"`
	MediaWidthPx             float64                       `json:"media_width_px" api:"nullable"`
	PlayingTimeMs            float64                       `json:"playing_time_ms"`
	RebufferCount            float64                       `json:"rebuffer_count"`
	RebufferDurationMs       float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio            float64                       `json:"rebuffer_ratio" api:"nullable"`
	RenditionPlayingTime     []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"nullable"`
	RenditionSwitchCount     float64                       `json:"rendition_switch_count" api:"nullable"`
	SeekCount                float64                       `json:"seek_count"`
	SeekDurationMs           float64                       `json:"seek_duration_ms"`
	StartupTimeMs            float64                       `json:"startup_time_ms" api:"nullable"`
	TotalVideoFrames         float64                       `json:"total_video_frames" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState RenditionChangedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayheadSeconds             respjson.Field
		PreviousRenditionBitrateBps respjson.Field
		PreviousRenditionCodec      respjson.Field
		PreviousRenditionFrameRate  respjson.Field
		PreviousRenditionHeight     respjson.Field
		PreviousRenditionWidth      respjson.Field
		Reason                      respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionFrameRate  respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		PlayingTimeMs               respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenditionChanged) RawJSON() string { return r.JSON.raw }
func (r *RenditionChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RenditionChanged to a RenditionChangedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RenditionChangedParam.Overrides()
func (r RenditionChanged) ToParam() RenditionChangedParam {
	return param.Override[RenditionChangedParam](json.RawMessage(r.RawJSON()))
}

type RenditionChangedErrorCategory string

const (
	RenditionChangedErrorCategoryAborted     RenditionChangedErrorCategory = "aborted"
	RenditionChangedErrorCategoryNetwork     RenditionChangedErrorCategory = "network"
	RenditionChangedErrorCategoryDecode      RenditionChangedErrorCategory = "decode"
	RenditionChangedErrorCategoryUnsupported RenditionChangedErrorCategory = "unsupported"
	RenditionChangedErrorCategoryUnknown     RenditionChangedErrorCategory = "unknown"
)

// RenditionChangedErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type RenditionChangedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u RenditionChangedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RenditionChangedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RenditionChangedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *RenditionChangedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RenditionChangedReason string

const (
	RenditionChangedReasonEnded         RenditionChangedReason = "ended"
	RenditionChangedReasonDestroyed     RenditionChangedReason = "destroyed"
	RenditionChangedReasonSourceChanged RenditionChangedReason = "source_changed"
	RenditionChangedReasonPageExit      RenditionChangedReason = "page_exit"
)

type RenditionChangedStartupOutcome string

const (
	RenditionChangedStartupOutcomeStarted           RenditionChangedStartupOutcome = "started"
	RenditionChangedStartupOutcomeFailedBeforeStart RenditionChangedStartupOutcome = "failed_before_start"
	RenditionChangedStartupOutcomeEndedBeforeStart  RenditionChangedStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type RenditionChangedVisibilityState string

const (
	RenditionChangedVisibilityStateVisible RenditionChangedVisibilityState = "visible"
	RenditionChangedVisibilityStateHidden  RenditionChangedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname, Environment,
// ErrorCategory, ErrorCode, ErrorFatal, ErrorPlayheadSeconds, EventID, EventName,
// OccurredAt, PlayerHeightPx, PlayerName, PlayerVersion, PlayerWidthPx,
// PlayheadSeconds, PreviousRenditionBitrateBps, PreviousRenditionCodec,
// PreviousRenditionFrameRate, PreviousRenditionHeight, PreviousRenditionWidth,
// Reason, SchemaVersion, SDKName, SDKVersion, SelectedRenditionBitrateBps,
// SelectedRenditionCodec, SelectedRenditionFrameRate, SelectedRenditionHeight,
// SelectedRenditionWidth, Sequence, StartupOutcome, VideoID, ViewID, ViewerID are
// required.
type RenditionChangedParam struct {
	ConnectorVersion param.Opt[string] `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname param.Opt[string] `json:"delivery_hostname,omitzero" api:"required"`
	Environment      param.Opt[string] `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory               param.Opt[RenditionChangedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal                  param.Opt[bool]                          `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds        param.Opt[float64]                       `json:"error_playhead_seconds,omitzero" api:"required"`
	PlayerHeightPx              param.Opt[float64]                       `json:"player_height_px,omitzero" api:"required"`
	PlayerName                  param.Opt[string]                        `json:"player_name,omitzero" api:"required"`
	PlayerVersion               param.Opt[string]                        `json:"player_version,omitzero" api:"required"`
	PlayerWidthPx               param.Opt[float64]                       `json:"player_width_px,omitzero" api:"required"`
	PlayheadSeconds             param.Opt[float64]                       `json:"playhead_seconds,omitzero" api:"required"`
	PreviousRenditionBitrateBps param.Opt[float64]                       `json:"previous_rendition_bitrate_bps,omitzero" api:"required"`
	PreviousRenditionCodec      param.Opt[string]                        `json:"previous_rendition_codec,omitzero" api:"required"`
	PreviousRenditionFrameRate  param.Opt[float64]                       `json:"previous_rendition_frame_rate,omitzero" api:"required"`
	PreviousRenditionHeight     param.Opt[float64]                       `json:"previous_rendition_height,omitzero" api:"required"`
	PreviousRenditionWidth      param.Opt[float64]                       `json:"previous_rendition_width,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason                      param.Opt[RenditionChangedReason] `json:"reason,omitzero" api:"required"`
	SelectedRenditionBitrateBps param.Opt[float64]                `json:"selected_rendition_bitrate_bps,omitzero" api:"required"`
	SelectedRenditionCodec      param.Opt[string]                 `json:"selected_rendition_codec,omitzero" api:"required"`
	SelectedRenditionFrameRate  param.Opt[float64]                `json:"selected_rendition_frame_rate,omitzero" api:"required"`
	SelectedRenditionHeight     param.Opt[float64]                `json:"selected_rendition_height,omitzero" api:"required"`
	SelectedRenditionWidth      param.Opt[float64]                `json:"selected_rendition_width,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome param.Opt[RenditionChangedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID      param.Opt[string]                   `json:"viewer_id,omitzero" api:"required"`
	ErrorCode     RenditionChangedErrorCodeUnionParam `json:"error_code,omitzero" api:"required"`
	ConnectorName string                              `json:"connector_name" api:"required"`
	EventID       string                              `json:"event_id" api:"required" format:"uuid"`
	OccurredAt    time.Time                           `json:"occurred_at" api:"required" format:"date-time"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                  string             `json:"sdk_name" api:"required"`
	SDKVersion               string             `json:"sdk_version" api:"required"`
	Sequence                 int64              `json:"sequence" api:"required"`
	VideoID                  string             `json:"video_id" api:"required"`
	ViewID                   string             `json:"view_id" api:"required" format:"uuid"`
	DeliveryRequestCount     param.Opt[float64] `json:"delivery_request_count,omitzero"`
	DeliveryThroughputBps    param.Opt[float64] `json:"delivery_throughput_bps,omitzero"`
	DeliveryTransferredBytes param.Opt[float64] `json:"delivery_transferred_bytes,omitzero"`
	DroppedVideoFrames       param.Opt[float64] `json:"dropped_video_frames,omitzero"`
	FirstVideoFrameTimeMs    param.Opt[float64] `json:"first_video_frame_time_ms,omitzero"`
	FurthestProgressRatio    param.Opt[float64] `json:"furthest_progress_ratio,omitzero"`
	MediaDurationSeconds     param.Opt[float64] `json:"media_duration_seconds,omitzero"`
	MediaHeightPx            param.Opt[float64] `json:"media_height_px,omitzero"`
	MediaWidthPx             param.Opt[float64] `json:"media_width_px,omitzero"`
	RebufferRatio            param.Opt[float64] `json:"rebuffer_ratio,omitzero"`
	RenditionSwitchCount     param.Opt[float64] `json:"rendition_switch_count,omitzero"`
	StartupTimeMs            param.Opt[float64] `json:"startup_time_ms,omitzero"`
	TotalVideoFrames         param.Opt[float64] `json:"total_video_frames,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState        param.Opt[RenditionChangedVisibilityState] `json:"visibility_state,omitzero"`
	MaximumPlayheadSeconds param.Opt[float64]                         `json:"maximum_playhead_seconds,omitzero"`
	PlayingTimeMs          param.Opt[float64]                         `json:"playing_time_ms,omitzero"`
	RebufferCount          param.Opt[float64]                         `json:"rebuffer_count,omitzero"`
	RebufferDurationMs     param.Opt[float64]                         `json:"rebuffer_duration_ms,omitzero"`
	SeekCount              param.Opt[float64]                         `json:"seek_count,omitzero"`
	SeekDurationMs         param.Opt[float64]                         `json:"seek_duration_ms,omitzero"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam         `json:"rendition_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.rendition_changed".
	EventName constant.ChunkifyVideoRenditionChanged `json:"event_name" default:"chunkify.video.rendition_changed"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r RenditionChangedParam) MarshalJSON() (data []byte, err error) {
	type shadow RenditionChangedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RenditionChangedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RenditionChangedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RenditionChangedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u RenditionChangedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *RenditionChangedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RenditionChangedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type SeekCompleted struct {
	ConnectorName    string `json:"connector_name" api:"required"`
	ConnectorVersion string `json:"connector_version" api:"required"`
	DeliveryHostname string `json:"delivery_hostname" api:"required"`
	Environment      string `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        SeekCompletedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode            SeekCompletedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal           bool                                `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds float64                             `json:"error_playhead_seconds" api:"required"`
	EventID              string                              `json:"event_id" api:"required" format:"uuid"`
	EventName            constant.ChunkifyVideoSeekCompleted `json:"event_name" default:"chunkify.video.seek_completed"`
	FromPlayheadSeconds  float64                             `json:"from_playhead_seconds" api:"required"`
	LastSeekedAt         time.Time                           `json:"last_seeked_at" api:"required" format:"date-time"`
	OccurredAt           time.Time                           `json:"occurred_at" api:"required" format:"date-time"`
	PlaybackResumedAt    time.Time                           `json:"playback_resumed_at" api:"required" format:"date-time"`
	PlayerName           string                              `json:"player_name" api:"required"`
	PlayerVersion        string                              `json:"player_version" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason        SeekCompletedReason `json:"reason" api:"required"`
	SchemaVersion float64             `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                   string  `json:"sdk_name" api:"required"`
	SDKVersion                string  `json:"sdk_version" api:"required"`
	SeekInteractionDurationMs float64 `json:"seek_interaction_duration_ms" api:"required"`
	// Any of "resumed", "paused", "interrupted".
	SeekOutcome      SeekCompletedSeekOutcome `json:"seek_outcome" api:"required"`
	SeekResumeWaitMs float64                  `json:"seek_resume_wait_ms" api:"required"`
	SeekStartedAt    time.Time                `json:"seek_started_at" api:"required" format:"date-time"`
	Sequence         int64                    `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome SeekCompletedStartupOutcome `json:"startup_outcome" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	StateAfterSeek SeekCompletedStateAfterSeek `json:"state_after_seek" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	StateBeforeSeek   SeekCompletedStateBeforeSeek `json:"state_before_seek" api:"required"`
	ToPlayheadSeconds float64                      `json:"to_playhead_seconds" api:"required"`
	VideoID           string                       `json:"video_id" api:"required"`
	ViewID            string                       `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext              any                           `json:"browser_context"`
	DeliveryRequestCount        float64                       `json:"delivery_request_count" api:"nullable"`
	DeliveryThroughputBps       float64                       `json:"delivery_throughput_bps" api:"nullable"`
	DeliveryTransferredBytes    float64                       `json:"delivery_transferred_bytes" api:"nullable"`
	DroppedVideoFrames          float64                       `json:"dropped_video_frames" api:"nullable"`
	FirstVideoFrameTimeMs       float64                       `json:"first_video_frame_time_ms" api:"nullable"`
	FurthestProgressRatio       float64                       `json:"furthest_progress_ratio" api:"nullable"`
	MaximumPlayheadSeconds      float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds        float64                       `json:"media_duration_seconds" api:"nullable"`
	MediaHeightPx               float64                       `json:"media_height_px" api:"nullable"`
	MediaWidthPx                float64                       `json:"media_width_px" api:"nullable"`
	PlayerHeightPx              float64                       `json:"player_height_px" api:"nullable"`
	PlayerWidthPx               float64                       `json:"player_width_px" api:"nullable"`
	PlayingTimeMs               float64                       `json:"playing_time_ms"`
	RebufferCount               float64                       `json:"rebuffer_count"`
	RebufferDurationMs          float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio               float64                       `json:"rebuffer_ratio" api:"nullable"`
	RenditionPlayingTime        []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"nullable"`
	RenditionSwitchCount        float64                       `json:"rendition_switch_count" api:"nullable"`
	SeekCount                   float64                       `json:"seek_count"`
	SeekDurationMs              float64                       `json:"seek_duration_ms"`
	SelectedRenditionBitrateBps float64                       `json:"selected_rendition_bitrate_bps" api:"nullable"`
	SelectedRenditionCodec      string                        `json:"selected_rendition_codec" api:"nullable"`
	SelectedRenditionHeight     float64                       `json:"selected_rendition_height" api:"nullable"`
	SelectedRenditionWidth      float64                       `json:"selected_rendition_width" api:"nullable"`
	StartupTimeMs               float64                       `json:"startup_time_ms" api:"nullable"`
	TotalVideoFrames            float64                       `json:"total_video_frames" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState SeekCompletedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FromPlayheadSeconds         respjson.Field
		LastSeekedAt                respjson.Field
		OccurredAt                  respjson.Field
		PlaybackResumedAt           respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		Reason                      respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekInteractionDurationMs   respjson.Field
		SeekOutcome                 respjson.Field
		SeekResumeWaitMs            respjson.Field
		SeekStartedAt               respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StateAfterSeek              respjson.Field
		StateBeforeSeek             respjson.Field
		ToPlayheadSeconds           respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeekCompleted) RawJSON() string { return r.JSON.raw }
func (r *SeekCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SeekCompleted to a SeekCompletedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SeekCompletedParam.Overrides()
func (r SeekCompleted) ToParam() SeekCompletedParam {
	return param.Override[SeekCompletedParam](json.RawMessage(r.RawJSON()))
}

type SeekCompletedErrorCategory string

const (
	SeekCompletedErrorCategoryAborted     SeekCompletedErrorCategory = "aborted"
	SeekCompletedErrorCategoryNetwork     SeekCompletedErrorCategory = "network"
	SeekCompletedErrorCategoryDecode      SeekCompletedErrorCategory = "decode"
	SeekCompletedErrorCategoryUnsupported SeekCompletedErrorCategory = "unsupported"
	SeekCompletedErrorCategoryUnknown     SeekCompletedErrorCategory = "unknown"
)

// SeekCompletedErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type SeekCompletedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u SeekCompletedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SeekCompletedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SeekCompletedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *SeekCompletedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SeekCompletedReason string

const (
	SeekCompletedReasonEnded         SeekCompletedReason = "ended"
	SeekCompletedReasonDestroyed     SeekCompletedReason = "destroyed"
	SeekCompletedReasonSourceChanged SeekCompletedReason = "source_changed"
	SeekCompletedReasonPageExit      SeekCompletedReason = "page_exit"
)

type SeekCompletedSeekOutcome string

const (
	SeekCompletedSeekOutcomeResumed     SeekCompletedSeekOutcome = "resumed"
	SeekCompletedSeekOutcomePaused      SeekCompletedSeekOutcome = "paused"
	SeekCompletedSeekOutcomeInterrupted SeekCompletedSeekOutcome = "interrupted"
)

type SeekCompletedStartupOutcome string

const (
	SeekCompletedStartupOutcomeStarted           SeekCompletedStartupOutcome = "started"
	SeekCompletedStartupOutcomeFailedBeforeStart SeekCompletedStartupOutcome = "failed_before_start"
	SeekCompletedStartupOutcomeEndedBeforeStart  SeekCompletedStartupOutcome = "ended_before_start"
)

type SeekCompletedStateAfterSeek string

const (
	SeekCompletedStateAfterSeekStartup   SeekCompletedStateAfterSeek = "startup"
	SeekCompletedStateAfterSeekResuming  SeekCompletedStateAfterSeek = "resuming"
	SeekCompletedStateAfterSeekPlaying   SeekCompletedStateAfterSeek = "playing"
	SeekCompletedStateAfterSeekPaused    SeekCompletedStateAfterSeek = "paused"
	SeekCompletedStateAfterSeekBuffering SeekCompletedStateAfterSeek = "buffering"
	SeekCompletedStateAfterSeekSeeking   SeekCompletedStateAfterSeek = "seeking"
	SeekCompletedStateAfterSeekError     SeekCompletedStateAfterSeek = "error"
	SeekCompletedStateAfterSeekEnded     SeekCompletedStateAfterSeek = "ended"
	SeekCompletedStateAfterSeekStopped   SeekCompletedStateAfterSeek = "stopped"
)

type SeekCompletedStateBeforeSeek string

const (
	SeekCompletedStateBeforeSeekStartup   SeekCompletedStateBeforeSeek = "startup"
	SeekCompletedStateBeforeSeekResuming  SeekCompletedStateBeforeSeek = "resuming"
	SeekCompletedStateBeforeSeekPlaying   SeekCompletedStateBeforeSeek = "playing"
	SeekCompletedStateBeforeSeekPaused    SeekCompletedStateBeforeSeek = "paused"
	SeekCompletedStateBeforeSeekBuffering SeekCompletedStateBeforeSeek = "buffering"
	SeekCompletedStateBeforeSeekSeeking   SeekCompletedStateBeforeSeek = "seeking"
	SeekCompletedStateBeforeSeekError     SeekCompletedStateBeforeSeek = "error"
	SeekCompletedStateBeforeSeekEnded     SeekCompletedStateBeforeSeek = "ended"
	SeekCompletedStateBeforeSeekStopped   SeekCompletedStateBeforeSeek = "stopped"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type SeekCompletedVisibilityState string

const (
	SeekCompletedVisibilityStateVisible SeekCompletedVisibilityState = "visible"
	SeekCompletedVisibilityStateHidden  SeekCompletedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname, Environment,
// ErrorCategory, ErrorCode, ErrorFatal, ErrorPlayheadSeconds, EventID, EventName,
// FromPlayheadSeconds, LastSeekedAt, OccurredAt, PlaybackResumedAt, PlayerName,
// PlayerVersion, Reason, SchemaVersion, SDKName, SDKVersion,
// SeekInteractionDurationMs, SeekOutcome, SeekResumeWaitMs, SeekStartedAt,
// Sequence, StartupOutcome, StateAfterSeek, StateBeforeSeek, ToPlayheadSeconds,
// VideoID, ViewID, ViewerID are required.
type SeekCompletedParam struct {
	ConnectorVersion param.Opt[string] `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname param.Opt[string] `json:"delivery_hostname,omitzero" api:"required"`
	Environment      param.Opt[string] `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        param.Opt[SeekCompletedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal           param.Opt[bool]                       `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds param.Opt[float64]                    `json:"error_playhead_seconds,omitzero" api:"required"`
	FromPlayheadSeconds  param.Opt[float64]                    `json:"from_playhead_seconds,omitzero" api:"required"`
	LastSeekedAt         param.Opt[time.Time]                  `json:"last_seeked_at,omitzero" api:"required" format:"date-time"`
	PlaybackResumedAt    param.Opt[time.Time]                  `json:"playback_resumed_at,omitzero" api:"required" format:"date-time"`
	PlayerName           param.Opt[string]                     `json:"player_name,omitzero" api:"required"`
	PlayerVersion        param.Opt[string]                     `json:"player_version,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason           param.Opt[SeekCompletedReason] `json:"reason,omitzero" api:"required"`
	SeekResumeWaitMs param.Opt[float64]             `json:"seek_resume_wait_ms,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome    param.Opt[SeekCompletedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	ToPlayheadSeconds param.Opt[float64]                     `json:"to_playhead_seconds,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID      param.Opt[string]                `json:"viewer_id,omitzero" api:"required"`
	ErrorCode     SeekCompletedErrorCodeUnionParam `json:"error_code,omitzero" api:"required"`
	ConnectorName string                           `json:"connector_name" api:"required"`
	EventID       string                           `json:"event_id" api:"required" format:"uuid"`
	OccurredAt    time.Time                        `json:"occurred_at" api:"required" format:"date-time"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                   string  `json:"sdk_name" api:"required"`
	SDKVersion                string  `json:"sdk_version" api:"required"`
	SeekInteractionDurationMs float64 `json:"seek_interaction_duration_ms" api:"required"`
	// Any of "resumed", "paused", "interrupted".
	SeekOutcome   SeekCompletedSeekOutcome `json:"seek_outcome,omitzero" api:"required"`
	SeekStartedAt time.Time                `json:"seek_started_at" api:"required" format:"date-time"`
	Sequence      int64                    `json:"sequence" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	StateAfterSeek SeekCompletedStateAfterSeek `json:"state_after_seek,omitzero" api:"required"`
	// Any of "startup", "resuming", "playing", "paused", "buffering", "seeking",
	// "error", "ended", "stopped".
	StateBeforeSeek             SeekCompletedStateBeforeSeek `json:"state_before_seek,omitzero" api:"required"`
	VideoID                     string                       `json:"video_id" api:"required"`
	ViewID                      string                       `json:"view_id" api:"required" format:"uuid"`
	DeliveryRequestCount        param.Opt[float64]           `json:"delivery_request_count,omitzero"`
	DeliveryThroughputBps       param.Opt[float64]           `json:"delivery_throughput_bps,omitzero"`
	DeliveryTransferredBytes    param.Opt[float64]           `json:"delivery_transferred_bytes,omitzero"`
	DroppedVideoFrames          param.Opt[float64]           `json:"dropped_video_frames,omitzero"`
	FirstVideoFrameTimeMs       param.Opt[float64]           `json:"first_video_frame_time_ms,omitzero"`
	FurthestProgressRatio       param.Opt[float64]           `json:"furthest_progress_ratio,omitzero"`
	MediaDurationSeconds        param.Opt[float64]           `json:"media_duration_seconds,omitzero"`
	MediaHeightPx               param.Opt[float64]           `json:"media_height_px,omitzero"`
	MediaWidthPx                param.Opt[float64]           `json:"media_width_px,omitzero"`
	PlayerHeightPx              param.Opt[float64]           `json:"player_height_px,omitzero"`
	PlayerWidthPx               param.Opt[float64]           `json:"player_width_px,omitzero"`
	RebufferRatio               param.Opt[float64]           `json:"rebuffer_ratio,omitzero"`
	RenditionSwitchCount        param.Opt[float64]           `json:"rendition_switch_count,omitzero"`
	SelectedRenditionBitrateBps param.Opt[float64]           `json:"selected_rendition_bitrate_bps,omitzero"`
	SelectedRenditionCodec      param.Opt[string]            `json:"selected_rendition_codec,omitzero"`
	SelectedRenditionHeight     param.Opt[float64]           `json:"selected_rendition_height,omitzero"`
	SelectedRenditionWidth      param.Opt[float64]           `json:"selected_rendition_width,omitzero"`
	StartupTimeMs               param.Opt[float64]           `json:"startup_time_ms,omitzero"`
	TotalVideoFrames            param.Opt[float64]           `json:"total_video_frames,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState        param.Opt[SeekCompletedVisibilityState] `json:"visibility_state,omitzero"`
	MaximumPlayheadSeconds param.Opt[float64]                      `json:"maximum_playhead_seconds,omitzero"`
	PlayingTimeMs          param.Opt[float64]                      `json:"playing_time_ms,omitzero"`
	RebufferCount          param.Opt[float64]                      `json:"rebuffer_count,omitzero"`
	RebufferDurationMs     param.Opt[float64]                      `json:"rebuffer_duration_ms,omitzero"`
	SeekCount              param.Opt[float64]                      `json:"seek_count,omitzero"`
	SeekDurationMs         param.Opt[float64]                      `json:"seek_duration_ms,omitzero"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam      `json:"rendition_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.seek_completed".
	EventName constant.ChunkifyVideoSeekCompleted `json:"event_name" default:"chunkify.video.seek_completed"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r SeekCompletedParam) MarshalJSON() (data []byte, err error) {
	type shadow SeekCompletedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeekCompletedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SeekCompletedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SeekCompletedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u SeekCompletedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *SeekCompletedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SeekCompletedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type StoredEvent struct {
	// A schema version 1 playback event. Extra properties are preserved.
	Event AnalyticsEventUnion `json:"event" api:"required"`
	// ID of the public key used to receive this event. Null when unavailable.
	IngestKeyID string `json:"ingest_key_id" api:"required"`
	// Trusted project identifier resolved from the public key. Historical events
	// migrate to default.
	ProjectID  string    `json:"project_id" api:"required"`
	ReceivedAt time.Time `json:"received_at" api:"required" format:"date-time"`
	// Browser context and locally derived continent/country/city. Request IP addresses
	// are never added to events or stored context.
	EventContext shared.EventContext `json:"event_context"`
	// Legacy context present for an event received through OTLP by an earlier
	// collector version. Existing context remains readable; new ingestion accepts
	// native JSON events only.
	Otlp shared.StoredOtlpContext `json:"otlp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event        respjson.Field
		IngestKeyID  respjson.Field
		ProjectID    respjson.Field
		ReceivedAt   respjson.Field
		EventContext respjson.Field
		Otlp         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoredEvent) RawJSON() string { return r.JSON.raw }
func (r *StoredEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtitleChanged struct {
	ConnectorName    string `json:"connector_name" api:"required"`
	ConnectorVersion string `json:"connector_version" api:"required"`
	DeliveryHostname string `json:"delivery_hostname" api:"required"`
	Environment      string `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        SubtitleChangedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode            SubtitleChangedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal           bool                                  `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds float64                               `json:"error_playhead_seconds" api:"required"`
	EventID              string                                `json:"event_id" api:"required" format:"uuid"`
	EventName            constant.ChunkifyVideoSubtitleChanged `json:"event_name" default:"chunkify.video.subtitle_changed"`
	OccurredAt           time.Time                             `json:"occurred_at" api:"required" format:"date-time"`
	PlayerName           string                                `json:"player_name" api:"required"`
	PlayerVersion        string                                `json:"player_version" api:"required"`
	PlayheadSeconds      float64                               `json:"playhead_seconds" api:"required"`
	// Language before the change; null when subtitles were off.
	PreviousSubtitleLanguage string `json:"previous_subtitle_language" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason        SubtitleChangedReason `json:"reason" api:"required"`
	SchemaVersion float64               `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName    string `json:"sdk_name" api:"required"`
	SDKVersion string `json:"sdk_version" api:"required"`
	Sequence   int64  `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome SubtitleChangedStartupOutcome `json:"startup_outcome" api:"required"`
	// Language after the change; null when subtitles are now off.
	SubtitleLanguage string `json:"subtitle_language" api:"required"`
	VideoID          string `json:"video_id" api:"required"`
	ViewID           string `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext              any                           `json:"browser_context"`
	DeliveryRequestCount        float64                       `json:"delivery_request_count" api:"nullable"`
	DeliveryThroughputBps       float64                       `json:"delivery_throughput_bps" api:"nullable"`
	DeliveryTransferredBytes    float64                       `json:"delivery_transferred_bytes" api:"nullable"`
	DroppedVideoFrames          float64                       `json:"dropped_video_frames" api:"nullable"`
	FirstVideoFrameTimeMs       float64                       `json:"first_video_frame_time_ms" api:"nullable"`
	FurthestProgressRatio       float64                       `json:"furthest_progress_ratio" api:"nullable"`
	MaximumPlayheadSeconds      float64                       `json:"maximum_playhead_seconds"`
	MediaDurationSeconds        float64                       `json:"media_duration_seconds" api:"nullable"`
	MediaHeightPx               float64                       `json:"media_height_px" api:"nullable"`
	MediaWidthPx                float64                       `json:"media_width_px" api:"nullable"`
	PlayerHeightPx              float64                       `json:"player_height_px" api:"nullable"`
	PlayerWidthPx               float64                       `json:"player_width_px" api:"nullable"`
	PlayingTimeMs               float64                       `json:"playing_time_ms"`
	RebufferCount               float64                       `json:"rebuffer_count"`
	RebufferDurationMs          float64                       `json:"rebuffer_duration_ms"`
	RebufferRatio               float64                       `json:"rebuffer_ratio" api:"nullable"`
	RenditionPlayingTime        []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"nullable"`
	RenditionSwitchCount        float64                       `json:"rendition_switch_count" api:"nullable"`
	SeekCount                   float64                       `json:"seek_count"`
	SeekDurationMs              float64                       `json:"seek_duration_ms"`
	SelectedRenditionBitrateBps float64                       `json:"selected_rendition_bitrate_bps" api:"nullable"`
	SelectedRenditionCodec      string                        `json:"selected_rendition_codec" api:"nullable"`
	SelectedRenditionHeight     float64                       `json:"selected_rendition_height" api:"nullable"`
	SelectedRenditionWidth      float64                       `json:"selected_rendition_width" api:"nullable"`
	StartupTimeMs               float64                       `json:"startup_time_ms" api:"nullable"`
	TotalVideoFrames            float64                       `json:"total_video_frames" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState SubtitleChangedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		OccurredAt                  respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayheadSeconds             respjson.Field
		PreviousSubtitleLanguage    respjson.Field
		Reason                      respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		SubtitleLanguage            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtitleChanged) RawJSON() string { return r.JSON.raw }
func (r *SubtitleChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubtitleChanged to a SubtitleChangedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubtitleChangedParam.Overrides()
func (r SubtitleChanged) ToParam() SubtitleChangedParam {
	return param.Override[SubtitleChangedParam](json.RawMessage(r.RawJSON()))
}

type SubtitleChangedErrorCategory string

const (
	SubtitleChangedErrorCategoryAborted     SubtitleChangedErrorCategory = "aborted"
	SubtitleChangedErrorCategoryNetwork     SubtitleChangedErrorCategory = "network"
	SubtitleChangedErrorCategoryDecode      SubtitleChangedErrorCategory = "decode"
	SubtitleChangedErrorCategoryUnsupported SubtitleChangedErrorCategory = "unsupported"
	SubtitleChangedErrorCategoryUnknown     SubtitleChangedErrorCategory = "unknown"
)

// SubtitleChangedErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type SubtitleChangedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u SubtitleChangedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SubtitleChangedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SubtitleChangedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *SubtitleChangedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtitleChangedReason string

const (
	SubtitleChangedReasonEnded         SubtitleChangedReason = "ended"
	SubtitleChangedReasonDestroyed     SubtitleChangedReason = "destroyed"
	SubtitleChangedReasonSourceChanged SubtitleChangedReason = "source_changed"
	SubtitleChangedReasonPageExit      SubtitleChangedReason = "page_exit"
)

type SubtitleChangedStartupOutcome string

const (
	SubtitleChangedStartupOutcomeStarted           SubtitleChangedStartupOutcome = "started"
	SubtitleChangedStartupOutcomeFailedBeforeStart SubtitleChangedStartupOutcome = "failed_before_start"
	SubtitleChangedStartupOutcomeEndedBeforeStart  SubtitleChangedStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type SubtitleChangedVisibilityState string

const (
	SubtitleChangedVisibilityStateVisible SubtitleChangedVisibilityState = "visible"
	SubtitleChangedVisibilityStateHidden  SubtitleChangedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname, Environment,
// ErrorCategory, ErrorCode, ErrorFatal, ErrorPlayheadSeconds, EventID, EventName,
// OccurredAt, PlayerName, PlayerVersion, PlayheadSeconds,
// PreviousSubtitleLanguage, Reason, SchemaVersion, SDKName, SDKVersion, Sequence,
// StartupOutcome, SubtitleLanguage, VideoID, ViewID, ViewerID are required.
type SubtitleChangedParam struct {
	ConnectorVersion param.Opt[string] `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname param.Opt[string] `json:"delivery_hostname,omitzero" api:"required"`
	Environment      param.Opt[string] `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory        param.Opt[SubtitleChangedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal           param.Opt[bool]                         `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds param.Opt[float64]                      `json:"error_playhead_seconds,omitzero" api:"required"`
	PlayerName           param.Opt[string]                       `json:"player_name,omitzero" api:"required"`
	PlayerVersion        param.Opt[string]                       `json:"player_version,omitzero" api:"required"`
	PlayheadSeconds      param.Opt[float64]                      `json:"playhead_seconds,omitzero" api:"required"`
	// Language before the change; null when subtitles were off.
	PreviousSubtitleLanguage param.Opt[string] `json:"previous_subtitle_language,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason param.Opt[SubtitleChangedReason] `json:"reason,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome param.Opt[SubtitleChangedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	// Language after the change; null when subtitles are now off.
	SubtitleLanguage param.Opt[string] `json:"subtitle_language,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID      param.Opt[string]                  `json:"viewer_id,omitzero" api:"required"`
	ErrorCode     SubtitleChangedErrorCodeUnionParam `json:"error_code,omitzero" api:"required"`
	ConnectorName string                             `json:"connector_name" api:"required"`
	EventID       string                             `json:"event_id" api:"required" format:"uuid"`
	OccurredAt    time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string             `json:"sdk_name" api:"required"`
	SDKVersion                  string             `json:"sdk_version" api:"required"`
	Sequence                    int64              `json:"sequence" api:"required"`
	VideoID                     string             `json:"video_id" api:"required"`
	ViewID                      string             `json:"view_id" api:"required" format:"uuid"`
	DeliveryRequestCount        param.Opt[float64] `json:"delivery_request_count,omitzero"`
	DeliveryThroughputBps       param.Opt[float64] `json:"delivery_throughput_bps,omitzero"`
	DeliveryTransferredBytes    param.Opt[float64] `json:"delivery_transferred_bytes,omitzero"`
	DroppedVideoFrames          param.Opt[float64] `json:"dropped_video_frames,omitzero"`
	FirstVideoFrameTimeMs       param.Opt[float64] `json:"first_video_frame_time_ms,omitzero"`
	FurthestProgressRatio       param.Opt[float64] `json:"furthest_progress_ratio,omitzero"`
	MediaDurationSeconds        param.Opt[float64] `json:"media_duration_seconds,omitzero"`
	MediaHeightPx               param.Opt[float64] `json:"media_height_px,omitzero"`
	MediaWidthPx                param.Opt[float64] `json:"media_width_px,omitzero"`
	PlayerHeightPx              param.Opt[float64] `json:"player_height_px,omitzero"`
	PlayerWidthPx               param.Opt[float64] `json:"player_width_px,omitzero"`
	RebufferRatio               param.Opt[float64] `json:"rebuffer_ratio,omitzero"`
	RenditionSwitchCount        param.Opt[float64] `json:"rendition_switch_count,omitzero"`
	SelectedRenditionBitrateBps param.Opt[float64] `json:"selected_rendition_bitrate_bps,omitzero"`
	SelectedRenditionCodec      param.Opt[string]  `json:"selected_rendition_codec,omitzero"`
	SelectedRenditionHeight     param.Opt[float64] `json:"selected_rendition_height,omitzero"`
	SelectedRenditionWidth      param.Opt[float64] `json:"selected_rendition_width,omitzero"`
	StartupTimeMs               param.Opt[float64] `json:"startup_time_ms,omitzero"`
	TotalVideoFrames            param.Opt[float64] `json:"total_video_frames,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState        param.Opt[SubtitleChangedVisibilityState] `json:"visibility_state,omitzero"`
	MaximumPlayheadSeconds param.Opt[float64]                        `json:"maximum_playhead_seconds,omitzero"`
	PlayingTimeMs          param.Opt[float64]                        `json:"playing_time_ms,omitzero"`
	RebufferCount          param.Opt[float64]                        `json:"rebuffer_count,omitzero"`
	RebufferDurationMs     param.Opt[float64]                        `json:"rebuffer_duration_ms,omitzero"`
	SeekCount              param.Opt[float64]                        `json:"seek_count,omitzero"`
	SeekDurationMs         param.Opt[float64]                        `json:"seek_duration_ms,omitzero"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam        `json:"rendition_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.subtitle_changed".
	EventName constant.ChunkifyVideoSubtitleChanged `json:"event_name" default:"chunkify.video.subtitle_changed"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r SubtitleChangedParam) MarshalJSON() (data []byte, err error) {
	type shadow SubtitleChangedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubtitleChangedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SubtitleChangedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SubtitleChangedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u SubtitleChangedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *SubtitleChangedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SubtitleChangedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ViewEnded struct {
	ConnectorName            string  `json:"connector_name" api:"required"`
	ConnectorVersion         string  `json:"connector_version" api:"required"`
	DeliveryHostname         string  `json:"delivery_hostname" api:"required"`
	DeliveryRequestCount     float64 `json:"delivery_request_count" api:"required"`
	DeliveryThroughputBps    float64 `json:"delivery_throughput_bps" api:"required"`
	DeliveryTransferredBytes float64 `json:"delivery_transferred_bytes" api:"required"`
	DroppedVideoFrames       float64 `json:"dropped_video_frames" api:"required"`
	Environment              string  `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory          ViewEndedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode              ViewEndedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal             bool                            `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds   float64                         `json:"error_playhead_seconds" api:"required"`
	EventID                string                          `json:"event_id" api:"required" format:"uuid"`
	EventName              constant.ChunkifyVideoViewEnded `json:"event_name" default:"chunkify.video.view_ended"`
	FirstVideoFrameTimeMs  float64                         `json:"first_video_frame_time_ms" api:"required"`
	FurthestProgressRatio  float64                         `json:"furthest_progress_ratio" api:"required"`
	MaximumPlayheadSeconds float64                         `json:"maximum_playhead_seconds" api:"required"`
	MediaDurationSeconds   float64                         `json:"media_duration_seconds" api:"required"`
	MediaHeightPx          float64                         `json:"media_height_px" api:"required"`
	MediaWidthPx           float64                         `json:"media_width_px" api:"required"`
	OccurredAt             time.Time                       `json:"occurred_at" api:"required" format:"date-time"`
	PlayerHeightPx         float64                         `json:"player_height_px" api:"required"`
	PlayerName             string                          `json:"player_name" api:"required"`
	PlayerVersion          string                          `json:"player_version" api:"required"`
	PlayerWidthPx          float64                         `json:"player_width_px" api:"required"`
	PlayingTimeMs          float64                         `json:"playing_time_ms" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason               ViewEndedReason               `json:"reason" api:"required"`
	RebufferCount        float64                       `json:"rebuffer_count" api:"required"`
	RebufferDurationMs   float64                       `json:"rebuffer_duration_ms" api:"required"`
	RebufferRatio        float64                       `json:"rebuffer_ratio" api:"required"`
	RenditionPlayingTime []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"required"`
	RenditionSwitchCount float64                       `json:"rendition_switch_count" api:"required"`
	SchemaVersion        float64                       `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string  `json:"sdk_name" api:"required"`
	SDKVersion                  string  `json:"sdk_version" api:"required"`
	SeekCount                   float64 `json:"seek_count" api:"required"`
	SeekDurationMs              float64 `json:"seek_duration_ms" api:"required"`
	SelectedRenditionBitrateBps float64 `json:"selected_rendition_bitrate_bps" api:"required"`
	SelectedRenditionCodec      string  `json:"selected_rendition_codec" api:"required"`
	SelectedRenditionHeight     float64 `json:"selected_rendition_height" api:"required"`
	SelectedRenditionWidth      float64 `json:"selected_rendition_width" api:"required"`
	Sequence                    int64   `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   ViewEndedStartupOutcome `json:"startup_outcome" api:"required"`
	StartupTimeMs    float64                 `json:"startup_time_ms" api:"required"`
	TotalVideoFrames float64                 `json:"total_video_frames" api:"required"`
	VideoID          string                  `json:"video_id" api:"required"`
	ViewID           string                  `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate float64 `json:"playback_rate" api:"nullable"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTime `json:"playback_rate_playing_time" api:"nullable"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRange `json:"played_ranges" api:"nullable"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds float64 `json:"playhead_seconds" api:"nullable"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage string `json:"subtitle_language" api:"nullable"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTime `json:"subtitle_playing_time" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState ViewEndedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		Reason                      respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		PlaybackRate                respjson.Field
		PlaybackRatePlayingTime     respjson.Field
		PlayedRanges                respjson.Field
		PlayheadSeconds             respjson.Field
		SubtitleLanguage            respjson.Field
		SubtitlePlayingTime         respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewEnded) RawJSON() string { return r.JSON.raw }
func (r *ViewEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewEnded to a ViewEndedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewEndedParam.Overrides()
func (r ViewEnded) ToParam() ViewEndedParam {
	return param.Override[ViewEndedParam](json.RawMessage(r.RawJSON()))
}

type ViewEndedErrorCategory string

const (
	ViewEndedErrorCategoryAborted     ViewEndedErrorCategory = "aborted"
	ViewEndedErrorCategoryNetwork     ViewEndedErrorCategory = "network"
	ViewEndedErrorCategoryDecode      ViewEndedErrorCategory = "decode"
	ViewEndedErrorCategoryUnsupported ViewEndedErrorCategory = "unsupported"
	ViewEndedErrorCategoryUnknown     ViewEndedErrorCategory = "unknown"
)

// ViewEndedErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type ViewEndedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ViewEndedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewEndedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewEndedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewEndedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewEndedReason string

const (
	ViewEndedReasonEnded         ViewEndedReason = "ended"
	ViewEndedReasonDestroyed     ViewEndedReason = "destroyed"
	ViewEndedReasonSourceChanged ViewEndedReason = "source_changed"
	ViewEndedReasonPageExit      ViewEndedReason = "page_exit"
)

type ViewEndedStartupOutcome string

const (
	ViewEndedStartupOutcomeStarted           ViewEndedStartupOutcome = "started"
	ViewEndedStartupOutcomeFailedBeforeStart ViewEndedStartupOutcome = "failed_before_start"
	ViewEndedStartupOutcomeEndedBeforeStart  ViewEndedStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type ViewEndedVisibilityState string

const (
	ViewEndedVisibilityStateVisible ViewEndedVisibilityState = "visible"
	ViewEndedVisibilityStateHidden  ViewEndedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname,
// DeliveryRequestCount, DeliveryThroughputBps, DeliveryTransferredBytes,
// DroppedVideoFrames, Environment, ErrorCategory, ErrorCode, ErrorFatal,
// ErrorPlayheadSeconds, EventID, EventName, FirstVideoFrameTimeMs,
// FurthestProgressRatio, MaximumPlayheadSeconds, MediaDurationSeconds,
// MediaHeightPx, MediaWidthPx, OccurredAt, PlayerHeightPx, PlayerName,
// PlayerVersion, PlayerWidthPx, PlayingTimeMs, Reason, RebufferCount,
// RebufferDurationMs, RebufferRatio, RenditionPlayingTime, RenditionSwitchCount,
// SchemaVersion, SDKName, SDKVersion, SeekCount, SeekDurationMs,
// SelectedRenditionBitrateBps, SelectedRenditionCodec, SelectedRenditionHeight,
// SelectedRenditionWidth, Sequence, StartupOutcome, StartupTimeMs,
// TotalVideoFrames, VideoID, ViewID, ViewerID are required.
type ViewEndedParam struct {
	ConnectorVersion         param.Opt[string]  `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname         param.Opt[string]  `json:"delivery_hostname,omitzero" api:"required"`
	DeliveryRequestCount     param.Opt[float64] `json:"delivery_request_count,omitzero" api:"required"`
	DeliveryThroughputBps    param.Opt[float64] `json:"delivery_throughput_bps,omitzero" api:"required"`
	DeliveryTransferredBytes param.Opt[float64] `json:"delivery_transferred_bytes,omitzero" api:"required"`
	DroppedVideoFrames       param.Opt[float64] `json:"dropped_video_frames,omitzero" api:"required"`
	Environment              param.Opt[string]  `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory         param.Opt[ViewEndedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal            param.Opt[bool]                   `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds  param.Opt[float64]                `json:"error_playhead_seconds,omitzero" api:"required"`
	FirstVideoFrameTimeMs param.Opt[float64]                `json:"first_video_frame_time_ms,omitzero" api:"required"`
	FurthestProgressRatio param.Opt[float64]                `json:"furthest_progress_ratio,omitzero" api:"required"`
	MediaDurationSeconds  param.Opt[float64]                `json:"media_duration_seconds,omitzero" api:"required"`
	MediaHeightPx         param.Opt[float64]                `json:"media_height_px,omitzero" api:"required"`
	MediaWidthPx          param.Opt[float64]                `json:"media_width_px,omitzero" api:"required"`
	PlayerHeightPx        param.Opt[float64]                `json:"player_height_px,omitzero" api:"required"`
	PlayerName            param.Opt[string]                 `json:"player_name,omitzero" api:"required"`
	PlayerVersion         param.Opt[string]                 `json:"player_version,omitzero" api:"required"`
	PlayerWidthPx         param.Opt[float64]                `json:"player_width_px,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason                      param.Opt[ViewEndedReason] `json:"reason,omitzero" api:"required"`
	RebufferRatio               param.Opt[float64]         `json:"rebuffer_ratio,omitzero" api:"required"`
	RenditionSwitchCount        param.Opt[float64]         `json:"rendition_switch_count,omitzero" api:"required"`
	SelectedRenditionBitrateBps param.Opt[float64]         `json:"selected_rendition_bitrate_bps,omitzero" api:"required"`
	SelectedRenditionCodec      param.Opt[string]          `json:"selected_rendition_codec,omitzero" api:"required"`
	SelectedRenditionHeight     param.Opt[float64]         `json:"selected_rendition_height,omitzero" api:"required"`
	SelectedRenditionWidth      param.Opt[float64]         `json:"selected_rendition_width,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   param.Opt[ViewEndedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	StartupTimeMs    param.Opt[float64]                 `json:"startup_time_ms,omitzero" api:"required"`
	TotalVideoFrames param.Opt[float64]                 `json:"total_video_frames,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID               param.Opt[string]                  `json:"viewer_id,omitzero" api:"required"`
	ErrorCode              ViewEndedErrorCodeUnionParam       `json:"error_code,omitzero" api:"required"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam `json:"rendition_playing_time,omitzero" api:"required"`
	ConnectorName          string                             `json:"connector_name" api:"required"`
	EventID                string                             `json:"event_id" api:"required" format:"uuid"`
	MaximumPlayheadSeconds float64                            `json:"maximum_playhead_seconds" api:"required"`
	OccurredAt             time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	PlayingTimeMs          float64                            `json:"playing_time_ms" api:"required"`
	RebufferCount          float64                            `json:"rebuffer_count" api:"required"`
	RebufferDurationMs     float64                            `json:"rebuffer_duration_ms" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName        string  `json:"sdk_name" api:"required"`
	SDKVersion     string  `json:"sdk_version" api:"required"`
	SeekCount      float64 `json:"seek_count" api:"required"`
	SeekDurationMs float64 `json:"seek_duration_ms" api:"required"`
	Sequence       int64   `json:"sequence" api:"required"`
	VideoID        string  `json:"video_id" api:"required"`
	ViewID         string  `json:"view_id" api:"required" format:"uuid"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate param.Opt[float64] `json:"playback_rate,omitzero"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds param.Opt[float64] `json:"playhead_seconds,omitzero"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage param.Opt[string] `json:"subtitle_language,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState param.Opt[ViewEndedVisibilityState] `json:"visibility_state,omitzero"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTimeParam `json:"playback_rate_playing_time,omitzero"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRangeParam `json:"played_ranges,omitzero"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTimeParam `json:"subtitle_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.view_ended".
	EventName constant.ChunkifyVideoViewEnded `json:"event_name" default:"chunkify.video.view_ended"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r ViewEndedParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewEndedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewEndedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewEndedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewEndedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u ViewEndedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *ViewEndedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ViewEndedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ViewSnapshot struct {
	ConnectorName            string  `json:"connector_name" api:"required"`
	ConnectorVersion         string  `json:"connector_version" api:"required"`
	DeliveryHostname         string  `json:"delivery_hostname" api:"required"`
	DeliveryRequestCount     float64 `json:"delivery_request_count" api:"required"`
	DeliveryThroughputBps    float64 `json:"delivery_throughput_bps" api:"required"`
	DeliveryTransferredBytes float64 `json:"delivery_transferred_bytes" api:"required"`
	DroppedVideoFrames       float64 `json:"dropped_video_frames" api:"required"`
	Environment              string  `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory          ViewSnapshotErrorCategory          `json:"error_category" api:"required"`
	ErrorCode              ViewSnapshotErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal             bool                               `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds   float64                            `json:"error_playhead_seconds" api:"required"`
	EventID                string                             `json:"event_id" api:"required" format:"uuid"`
	EventName              constant.ChunkifyVideoViewSnapshot `json:"event_name" default:"chunkify.video.view_snapshot"`
	FirstVideoFrameTimeMs  float64                            `json:"first_video_frame_time_ms" api:"required"`
	FurthestProgressRatio  float64                            `json:"furthest_progress_ratio" api:"required"`
	MaximumPlayheadSeconds float64                            `json:"maximum_playhead_seconds" api:"required"`
	MediaDurationSeconds   float64                            `json:"media_duration_seconds" api:"required"`
	MediaHeightPx          float64                            `json:"media_height_px" api:"required"`
	MediaWidthPx           float64                            `json:"media_width_px" api:"required"`
	OccurredAt             time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	PlayerHeightPx         float64                            `json:"player_height_px" api:"required"`
	PlayerName             string                             `json:"player_name" api:"required"`
	PlayerVersion          string                             `json:"player_version" api:"required"`
	PlayerWidthPx          float64                            `json:"player_width_px" api:"required"`
	PlayingTimeMs          float64                            `json:"playing_time_ms" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason               ViewSnapshotReason            `json:"reason" api:"required"`
	RebufferCount        float64                       `json:"rebuffer_count" api:"required"`
	RebufferDurationMs   float64                       `json:"rebuffer_duration_ms" api:"required"`
	RebufferRatio        float64                       `json:"rebuffer_ratio" api:"required"`
	RenditionPlayingTime []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"required"`
	RenditionSwitchCount float64                       `json:"rendition_switch_count" api:"required"`
	SchemaVersion        float64                       `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string  `json:"sdk_name" api:"required"`
	SDKVersion                  string  `json:"sdk_version" api:"required"`
	SeekCount                   float64 `json:"seek_count" api:"required"`
	SeekDurationMs              float64 `json:"seek_duration_ms" api:"required"`
	SelectedRenditionBitrateBps float64 `json:"selected_rendition_bitrate_bps" api:"required"`
	SelectedRenditionCodec      string  `json:"selected_rendition_codec" api:"required"`
	SelectedRenditionHeight     float64 `json:"selected_rendition_height" api:"required"`
	SelectedRenditionWidth      float64 `json:"selected_rendition_width" api:"required"`
	Sequence                    int64   `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   ViewSnapshotStartupOutcome `json:"startup_outcome" api:"required"`
	StartupTimeMs    float64                    `json:"startup_time_ms" api:"required"`
	TotalVideoFrames float64                    `json:"total_video_frames" api:"required"`
	VideoID          string                     `json:"video_id" api:"required"`
	ViewID           string                     `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate float64 `json:"playback_rate" api:"nullable"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTime `json:"playback_rate_playing_time" api:"nullable"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRange `json:"played_ranges" api:"nullable"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds float64 `json:"playhead_seconds" api:"nullable"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage string `json:"subtitle_language" api:"nullable"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTime `json:"subtitle_playing_time" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState ViewSnapshotVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		Reason                      respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		PlaybackRate                respjson.Field
		PlaybackRatePlayingTime     respjson.Field
		PlayedRanges                respjson.Field
		PlayheadSeconds             respjson.Field
		SubtitleLanguage            respjson.Field
		SubtitlePlayingTime         respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewSnapshot) RawJSON() string { return r.JSON.raw }
func (r *ViewSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewSnapshot to a ViewSnapshotParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewSnapshotParam.Overrides()
func (r ViewSnapshot) ToParam() ViewSnapshotParam {
	return param.Override[ViewSnapshotParam](json.RawMessage(r.RawJSON()))
}

type ViewSnapshotErrorCategory string

const (
	ViewSnapshotErrorCategoryAborted     ViewSnapshotErrorCategory = "aborted"
	ViewSnapshotErrorCategoryNetwork     ViewSnapshotErrorCategory = "network"
	ViewSnapshotErrorCategoryDecode      ViewSnapshotErrorCategory = "decode"
	ViewSnapshotErrorCategoryUnsupported ViewSnapshotErrorCategory = "unsupported"
	ViewSnapshotErrorCategoryUnknown     ViewSnapshotErrorCategory = "unknown"
)

// ViewSnapshotErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type ViewSnapshotErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ViewSnapshotErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewSnapshotErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewSnapshotErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewSnapshotErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewSnapshotReason string

const (
	ViewSnapshotReasonEnded         ViewSnapshotReason = "ended"
	ViewSnapshotReasonDestroyed     ViewSnapshotReason = "destroyed"
	ViewSnapshotReasonSourceChanged ViewSnapshotReason = "source_changed"
	ViewSnapshotReasonPageExit      ViewSnapshotReason = "page_exit"
)

type ViewSnapshotStartupOutcome string

const (
	ViewSnapshotStartupOutcomeStarted           ViewSnapshotStartupOutcome = "started"
	ViewSnapshotStartupOutcomeFailedBeforeStart ViewSnapshotStartupOutcome = "failed_before_start"
	ViewSnapshotStartupOutcomeEndedBeforeStart  ViewSnapshotStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type ViewSnapshotVisibilityState string

const (
	ViewSnapshotVisibilityStateVisible ViewSnapshotVisibilityState = "visible"
	ViewSnapshotVisibilityStateHidden  ViewSnapshotVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname,
// DeliveryRequestCount, DeliveryThroughputBps, DeliveryTransferredBytes,
// DroppedVideoFrames, Environment, ErrorCategory, ErrorCode, ErrorFatal,
// ErrorPlayheadSeconds, EventID, EventName, FirstVideoFrameTimeMs,
// FurthestProgressRatio, MaximumPlayheadSeconds, MediaDurationSeconds,
// MediaHeightPx, MediaWidthPx, OccurredAt, PlayerHeightPx, PlayerName,
// PlayerVersion, PlayerWidthPx, PlayingTimeMs, Reason, RebufferCount,
// RebufferDurationMs, RebufferRatio, RenditionPlayingTime, RenditionSwitchCount,
// SchemaVersion, SDKName, SDKVersion, SeekCount, SeekDurationMs,
// SelectedRenditionBitrateBps, SelectedRenditionCodec, SelectedRenditionHeight,
// SelectedRenditionWidth, Sequence, StartupOutcome, StartupTimeMs,
// TotalVideoFrames, VideoID, ViewID, ViewerID are required.
type ViewSnapshotParam struct {
	ConnectorVersion         param.Opt[string]  `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname         param.Opt[string]  `json:"delivery_hostname,omitzero" api:"required"`
	DeliveryRequestCount     param.Opt[float64] `json:"delivery_request_count,omitzero" api:"required"`
	DeliveryThroughputBps    param.Opt[float64] `json:"delivery_throughput_bps,omitzero" api:"required"`
	DeliveryTransferredBytes param.Opt[float64] `json:"delivery_transferred_bytes,omitzero" api:"required"`
	DroppedVideoFrames       param.Opt[float64] `json:"dropped_video_frames,omitzero" api:"required"`
	Environment              param.Opt[string]  `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory         param.Opt[ViewSnapshotErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal            param.Opt[bool]                      `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds  param.Opt[float64]                   `json:"error_playhead_seconds,omitzero" api:"required"`
	FirstVideoFrameTimeMs param.Opt[float64]                   `json:"first_video_frame_time_ms,omitzero" api:"required"`
	FurthestProgressRatio param.Opt[float64]                   `json:"furthest_progress_ratio,omitzero" api:"required"`
	MediaDurationSeconds  param.Opt[float64]                   `json:"media_duration_seconds,omitzero" api:"required"`
	MediaHeightPx         param.Opt[float64]                   `json:"media_height_px,omitzero" api:"required"`
	MediaWidthPx          param.Opt[float64]                   `json:"media_width_px,omitzero" api:"required"`
	PlayerHeightPx        param.Opt[float64]                   `json:"player_height_px,omitzero" api:"required"`
	PlayerName            param.Opt[string]                    `json:"player_name,omitzero" api:"required"`
	PlayerVersion         param.Opt[string]                    `json:"player_version,omitzero" api:"required"`
	PlayerWidthPx         param.Opt[float64]                   `json:"player_width_px,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason                      param.Opt[ViewSnapshotReason] `json:"reason,omitzero" api:"required"`
	RebufferRatio               param.Opt[float64]            `json:"rebuffer_ratio,omitzero" api:"required"`
	RenditionSwitchCount        param.Opt[float64]            `json:"rendition_switch_count,omitzero" api:"required"`
	SelectedRenditionBitrateBps param.Opt[float64]            `json:"selected_rendition_bitrate_bps,omitzero" api:"required"`
	SelectedRenditionCodec      param.Opt[string]             `json:"selected_rendition_codec,omitzero" api:"required"`
	SelectedRenditionHeight     param.Opt[float64]            `json:"selected_rendition_height,omitzero" api:"required"`
	SelectedRenditionWidth      param.Opt[float64]            `json:"selected_rendition_width,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   param.Opt[ViewSnapshotStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	StartupTimeMs    param.Opt[float64]                    `json:"startup_time_ms,omitzero" api:"required"`
	TotalVideoFrames param.Opt[float64]                    `json:"total_video_frames,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID               param.Opt[string]                  `json:"viewer_id,omitzero" api:"required"`
	ErrorCode              ViewSnapshotErrorCodeUnionParam    `json:"error_code,omitzero" api:"required"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam `json:"rendition_playing_time,omitzero" api:"required"`
	ConnectorName          string                             `json:"connector_name" api:"required"`
	EventID                string                             `json:"event_id" api:"required" format:"uuid"`
	MaximumPlayheadSeconds float64                            `json:"maximum_playhead_seconds" api:"required"`
	OccurredAt             time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	PlayingTimeMs          float64                            `json:"playing_time_ms" api:"required"`
	RebufferCount          float64                            `json:"rebuffer_count" api:"required"`
	RebufferDurationMs     float64                            `json:"rebuffer_duration_ms" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName        string  `json:"sdk_name" api:"required"`
	SDKVersion     string  `json:"sdk_version" api:"required"`
	SeekCount      float64 `json:"seek_count" api:"required"`
	SeekDurationMs float64 `json:"seek_duration_ms" api:"required"`
	Sequence       int64   `json:"sequence" api:"required"`
	VideoID        string  `json:"video_id" api:"required"`
	ViewID         string  `json:"view_id" api:"required" format:"uuid"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate param.Opt[float64] `json:"playback_rate,omitzero"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds param.Opt[float64] `json:"playhead_seconds,omitzero"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage param.Opt[string] `json:"subtitle_language,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState param.Opt[ViewSnapshotVisibilityState] `json:"visibility_state,omitzero"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTimeParam `json:"playback_rate_playing_time,omitzero"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRangeParam `json:"played_ranges,omitzero"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTimeParam `json:"subtitle_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.view_snapshot".
	EventName constant.ChunkifyVideoViewSnapshot `json:"event_name" default:"chunkify.video.view_snapshot"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r ViewSnapshotParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewSnapshotParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewSnapshotParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewSnapshotParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewSnapshotErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u ViewSnapshotErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *ViewSnapshotErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ViewSnapshotErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ViewStarted struct {
	ConnectorName            string  `json:"connector_name" api:"required"`
	ConnectorVersion         string  `json:"connector_version" api:"required"`
	DeliveryHostname         string  `json:"delivery_hostname" api:"required"`
	DeliveryRequestCount     float64 `json:"delivery_request_count" api:"required"`
	DeliveryThroughputBps    float64 `json:"delivery_throughput_bps" api:"required"`
	DeliveryTransferredBytes float64 `json:"delivery_transferred_bytes" api:"required"`
	DroppedVideoFrames       float64 `json:"dropped_video_frames" api:"required"`
	Environment              string  `json:"environment" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory          ViewStartedErrorCategory          `json:"error_category" api:"required"`
	ErrorCode              ViewStartedErrorCodeUnion         `json:"error_code" api:"required"`
	ErrorFatal             bool                              `json:"error_fatal" api:"required"`
	ErrorPlayheadSeconds   float64                           `json:"error_playhead_seconds" api:"required"`
	EventID                string                            `json:"event_id" api:"required" format:"uuid"`
	EventName              constant.ChunkifyVideoViewStarted `json:"event_name" default:"chunkify.video.view_started"`
	FirstVideoFrameTimeMs  float64                           `json:"first_video_frame_time_ms" api:"required"`
	FurthestProgressRatio  float64                           `json:"furthest_progress_ratio" api:"required"`
	MaximumPlayheadSeconds float64                           `json:"maximum_playhead_seconds" api:"required"`
	MediaDurationSeconds   float64                           `json:"media_duration_seconds" api:"required"`
	MediaHeightPx          float64                           `json:"media_height_px" api:"required"`
	MediaWidthPx           float64                           `json:"media_width_px" api:"required"`
	OccurredAt             time.Time                         `json:"occurred_at" api:"required" format:"date-time"`
	PlayerHeightPx         float64                           `json:"player_height_px" api:"required"`
	PlayerName             string                            `json:"player_name" api:"required"`
	PlayerVersion          string                            `json:"player_version" api:"required"`
	PlayerWidthPx          float64                           `json:"player_width_px" api:"required"`
	PlayingTimeMs          float64                           `json:"playing_time_ms" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason               ViewStartedReason             `json:"reason" api:"required"`
	RebufferCount        float64                       `json:"rebuffer_count" api:"required"`
	RebufferDurationMs   float64                       `json:"rebuffer_duration_ms" api:"required"`
	RebufferRatio        float64                       `json:"rebuffer_ratio" api:"required"`
	RenditionPlayingTime []shared.RenditionPlayingTime `json:"rendition_playing_time" api:"required"`
	RenditionSwitchCount float64                       `json:"rendition_switch_count" api:"required"`
	SchemaVersion        float64                       `json:"schema_version" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName                     string  `json:"sdk_name" api:"required"`
	SDKVersion                  string  `json:"sdk_version" api:"required"`
	SeekCount                   float64 `json:"seek_count" api:"required"`
	SeekDurationMs              float64 `json:"seek_duration_ms" api:"required"`
	SelectedRenditionBitrateBps float64 `json:"selected_rendition_bitrate_bps" api:"required"`
	SelectedRenditionCodec      string  `json:"selected_rendition_codec" api:"required"`
	SelectedRenditionHeight     float64 `json:"selected_rendition_height" api:"required"`
	SelectedRenditionWidth      float64 `json:"selected_rendition_width" api:"required"`
	Sequence                    int64   `json:"sequence" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   ViewStartedStartupOutcome `json:"startup_outcome" api:"required"`
	StartupTimeMs    float64                   `json:"startup_time_ms" api:"required"`
	TotalVideoFrames float64                   `json:"total_video_frames" api:"required"`
	VideoID          string                    `json:"video_id" api:"required"`
	ViewID           string                    `json:"view_id" api:"required" format:"uuid"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID string `json:"viewer_id" api:"required"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context"`
	// Integrator-defined dimensions set once when the view starts: at most 10 keys of
	// 1 to 40 characters a-z, 0-9 and \_, each with text of 1 to 256 characters, a
	// number or a boolean. Reports compare values as text, exactly and
	// case-sensitively: numbers as their JSON text (5, 5.5; send whole numbers without
	// a decimal part, as JavaScript does), booleans as true or false. Only
	// view_started accepts custom. Never put personal data here; use user_id for user
	// identity.
	Custom shared.CustomDimensions `json:"custom"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate float64 `json:"playback_rate" api:"nullable"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTime `json:"playback_rate_playing_time" api:"nullable"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRange `json:"played_ranges" api:"nullable"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds float64 `json:"playhead_seconds" api:"nullable"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage string `json:"subtitle_language" api:"nullable"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTime `json:"subtitle_playing_time" api:"nullable"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID string `json:"user_id" api:"nullable"`
	// Optional human-readable video title for reports.
	VideoTitle string `json:"video_title" api:"nullable"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState ViewStartedVisibilityState `json:"visibility_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName               respjson.Field
		ConnectorVersion            respjson.Field
		DeliveryHostname            respjson.Field
		DeliveryRequestCount        respjson.Field
		DeliveryThroughputBps       respjson.Field
		DeliveryTransferredBytes    respjson.Field
		DroppedVideoFrames          respjson.Field
		Environment                 respjson.Field
		ErrorCategory               respjson.Field
		ErrorCode                   respjson.Field
		ErrorFatal                  respjson.Field
		ErrorPlayheadSeconds        respjson.Field
		EventID                     respjson.Field
		EventName                   respjson.Field
		FirstVideoFrameTimeMs       respjson.Field
		FurthestProgressRatio       respjson.Field
		MaximumPlayheadSeconds      respjson.Field
		MediaDurationSeconds        respjson.Field
		MediaHeightPx               respjson.Field
		MediaWidthPx                respjson.Field
		OccurredAt                  respjson.Field
		PlayerHeightPx              respjson.Field
		PlayerName                  respjson.Field
		PlayerVersion               respjson.Field
		PlayerWidthPx               respjson.Field
		PlayingTimeMs               respjson.Field
		Reason                      respjson.Field
		RebufferCount               respjson.Field
		RebufferDurationMs          respjson.Field
		RebufferRatio               respjson.Field
		RenditionPlayingTime        respjson.Field
		RenditionSwitchCount        respjson.Field
		SchemaVersion               respjson.Field
		SDKName                     respjson.Field
		SDKVersion                  respjson.Field
		SeekCount                   respjson.Field
		SeekDurationMs              respjson.Field
		SelectedRenditionBitrateBps respjson.Field
		SelectedRenditionCodec      respjson.Field
		SelectedRenditionHeight     respjson.Field
		SelectedRenditionWidth      respjson.Field
		Sequence                    respjson.Field
		StartupOutcome              respjson.Field
		StartupTimeMs               respjson.Field
		TotalVideoFrames            respjson.Field
		VideoID                     respjson.Field
		ViewID                      respjson.Field
		ViewerID                    respjson.Field
		BrowserContext              respjson.Field
		Custom                      respjson.Field
		PlaybackRate                respjson.Field
		PlaybackRatePlayingTime     respjson.Field
		PlayedRanges                respjson.Field
		PlayheadSeconds             respjson.Field
		SubtitleLanguage            respjson.Field
		SubtitlePlayingTime         respjson.Field
		UserID                      respjson.Field
		VideoTitle                  respjson.Field
		VisibilityState             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewStarted) RawJSON() string { return r.JSON.raw }
func (r *ViewStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewStarted to a ViewStartedParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewStartedParam.Overrides()
func (r ViewStarted) ToParam() ViewStartedParam {
	return param.Override[ViewStartedParam](json.RawMessage(r.RawJSON()))
}

type ViewStartedErrorCategory string

const (
	ViewStartedErrorCategoryAborted     ViewStartedErrorCategory = "aborted"
	ViewStartedErrorCategoryNetwork     ViewStartedErrorCategory = "network"
	ViewStartedErrorCategoryDecode      ViewStartedErrorCategory = "decode"
	ViewStartedErrorCategoryUnsupported ViewStartedErrorCategory = "unsupported"
	ViewStartedErrorCategoryUnknown     ViewStartedErrorCategory = "unknown"
)

// ViewStartedErrorCodeUnion contains all possible properties and values from
// [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type ViewStartedErrorCodeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ViewStartedErrorCodeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewStartedErrorCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewStartedErrorCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewStartedErrorCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewStartedReason string

const (
	ViewStartedReasonEnded         ViewStartedReason = "ended"
	ViewStartedReasonDestroyed     ViewStartedReason = "destroyed"
	ViewStartedReasonSourceChanged ViewStartedReason = "source_changed"
	ViewStartedReasonPageExit      ViewStartedReason = "page_exit"
)

type ViewStartedStartupOutcome string

const (
	ViewStartedStartupOutcomeStarted           ViewStartedStartupOutcome = "started"
	ViewStartedStartupOutcomeFailedBeforeStart ViewStartedStartupOutcome = "failed_before_start"
	ViewStartedStartupOutcomeEndedBeforeStart  ViewStartedStartupOutcome = "ended_before_start"
)

// Tab visibility when the event was recorded. Hidden does not imply paused
// playback or identify the cause of a pause. Missing on older events; null when
// unavailable.
type ViewStartedVisibilityState string

const (
	ViewStartedVisibilityStateVisible ViewStartedVisibilityState = "visible"
	ViewStartedVisibilityStateHidden  ViewStartedVisibilityState = "hidden"
)

// The properties ConnectorName, ConnectorVersion, DeliveryHostname,
// DeliveryRequestCount, DeliveryThroughputBps, DeliveryTransferredBytes,
// DroppedVideoFrames, Environment, ErrorCategory, ErrorCode, ErrorFatal,
// ErrorPlayheadSeconds, EventID, EventName, FirstVideoFrameTimeMs,
// FurthestProgressRatio, MaximumPlayheadSeconds, MediaDurationSeconds,
// MediaHeightPx, MediaWidthPx, OccurredAt, PlayerHeightPx, PlayerName,
// PlayerVersion, PlayerWidthPx, PlayingTimeMs, Reason, RebufferCount,
// RebufferDurationMs, RebufferRatio, RenditionPlayingTime, RenditionSwitchCount,
// SchemaVersion, SDKName, SDKVersion, SeekCount, SeekDurationMs,
// SelectedRenditionBitrateBps, SelectedRenditionCodec, SelectedRenditionHeight,
// SelectedRenditionWidth, Sequence, StartupOutcome, StartupTimeMs,
// TotalVideoFrames, VideoID, ViewID, ViewerID are required.
type ViewStartedParam struct {
	ConnectorVersion         param.Opt[string]  `json:"connector_version,omitzero" api:"required"`
	DeliveryHostname         param.Opt[string]  `json:"delivery_hostname,omitzero" api:"required"`
	DeliveryRequestCount     param.Opt[float64] `json:"delivery_request_count,omitzero" api:"required"`
	DeliveryThroughputBps    param.Opt[float64] `json:"delivery_throughput_bps,omitzero" api:"required"`
	DeliveryTransferredBytes param.Opt[float64] `json:"delivery_transferred_bytes,omitzero" api:"required"`
	DroppedVideoFrames       param.Opt[float64] `json:"dropped_video_frames,omitzero" api:"required"`
	Environment              param.Opt[string]  `json:"environment,omitzero" api:"required"`
	// Any of "aborted", "network", "decode", "unsupported", "unknown".
	ErrorCategory         param.Opt[ViewStartedErrorCategory] `json:"error_category,omitzero" api:"required"`
	ErrorFatal            param.Opt[bool]                     `json:"error_fatal,omitzero" api:"required"`
	ErrorPlayheadSeconds  param.Opt[float64]                  `json:"error_playhead_seconds,omitzero" api:"required"`
	FirstVideoFrameTimeMs param.Opt[float64]                  `json:"first_video_frame_time_ms,omitzero" api:"required"`
	FurthestProgressRatio param.Opt[float64]                  `json:"furthest_progress_ratio,omitzero" api:"required"`
	MediaDurationSeconds  param.Opt[float64]                  `json:"media_duration_seconds,omitzero" api:"required"`
	MediaHeightPx         param.Opt[float64]                  `json:"media_height_px,omitzero" api:"required"`
	MediaWidthPx          param.Opt[float64]                  `json:"media_width_px,omitzero" api:"required"`
	PlayerHeightPx        param.Opt[float64]                  `json:"player_height_px,omitzero" api:"required"`
	PlayerName            param.Opt[string]                   `json:"player_name,omitzero" api:"required"`
	PlayerVersion         param.Opt[string]                   `json:"player_version,omitzero" api:"required"`
	PlayerWidthPx         param.Opt[float64]                  `json:"player_width_px,omitzero" api:"required"`
	// Any of "ended", "destroyed", "source_changed", "page_exit".
	Reason                      param.Opt[ViewStartedReason] `json:"reason,omitzero" api:"required"`
	RebufferRatio               param.Opt[float64]           `json:"rebuffer_ratio,omitzero" api:"required"`
	RenditionSwitchCount        param.Opt[float64]           `json:"rendition_switch_count,omitzero" api:"required"`
	SelectedRenditionBitrateBps param.Opt[float64]           `json:"selected_rendition_bitrate_bps,omitzero" api:"required"`
	SelectedRenditionCodec      param.Opt[string]            `json:"selected_rendition_codec,omitzero" api:"required"`
	SelectedRenditionHeight     param.Opt[float64]           `json:"selected_rendition_height,omitzero" api:"required"`
	SelectedRenditionWidth      param.Opt[float64]           `json:"selected_rendition_width,omitzero" api:"required"`
	// Any of "started", "failed_before_start", "ended_before_start".
	StartupOutcome   param.Opt[ViewStartedStartupOutcome] `json:"startup_outcome,omitzero" api:"required"`
	StartupTimeMs    param.Opt[float64]                   `json:"startup_time_ms,omitzero" api:"required"`
	TotalVideoFrames param.Opt[float64]                   `json:"total_video_frames,omitzero" api:"required"`
	// Anonymous browser identifier generated by the Chunkify player SDK and stored in
	// the page's localStorage.
	ViewerID               param.Opt[string]                  `json:"viewer_id,omitzero" api:"required"`
	ErrorCode              ViewStartedErrorCodeUnionParam     `json:"error_code,omitzero" api:"required"`
	RenditionPlayingTime   []shared.RenditionPlayingTimeParam `json:"rendition_playing_time,omitzero" api:"required"`
	ConnectorName          string                             `json:"connector_name" api:"required"`
	EventID                string                             `json:"event_id" api:"required" format:"uuid"`
	MaximumPlayheadSeconds float64                            `json:"maximum_playhead_seconds" api:"required"`
	OccurredAt             time.Time                          `json:"occurred_at" api:"required" format:"date-time"`
	PlayingTimeMs          float64                            `json:"playing_time_ms" api:"required"`
	RebufferCount          float64                            `json:"rebuffer_count" api:"required"`
	RebufferDurationMs     float64                            `json:"rebuffer_duration_ms" api:"required"`
	// Name of the SDK or integration that produced the event. Informational metadata,
	// preserved as received. Must contain at least one non-whitespace character.
	SDKName        string  `json:"sdk_name" api:"required"`
	SDKVersion     string  `json:"sdk_version" api:"required"`
	SeekCount      float64 `json:"seek_count" api:"required"`
	SeekDurationMs float64 `json:"seek_duration_ms" api:"required"`
	Sequence       int64   `json:"sequence" api:"required"`
	VideoID        string  `json:"video_id" api:"required"`
	ViewID         string  `json:"view_id" api:"required" format:"uuid"`
	// Current playback rate, rounded to two decimals, such as 1, 1.5 or 2; from 0.01
	// to 16. Optional.
	PlaybackRate param.Opt[float64] `json:"playback_rate,omitzero"`
	// Current media position at this event, not the maximum reached.
	PlayheadSeconds param.Opt[float64] `json:"playhead_seconds,omitzero"`
	// Language tag (BCP 47) of the showing subtitles or captions track; und when the
	// track declares no language; null when subtitles are off. Optional.
	SubtitleLanguage param.Opt[string] `json:"subtitle_language,omitzero"`
	// Optional integrator-supplied user ID, used to link views to the integrator's own
	// users.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Optional human-readable video title for reports.
	VideoTitle param.Opt[string] `json:"video_title,omitzero"`
	// Tab visibility when the event was recorded. Hidden does not imply paused
	// playback or identify the cause of a pause. Missing on older events; null when
	// unavailable.
	//
	// Any of "visible", "hidden".
	VisibilityState param.Opt[ViewStartedVisibilityState] `json:"visibility_state,omitzero"`
	// Cumulative playing time per playback rate for this attempt. At most 32 entries.
	// Optional.
	PlaybackRatePlayingTime []shared.PlaybackRatePlayingTimeParam `json:"playback_rate_playing_time,omitzero"`
	// Cumulative played media ranges for this attempt, sorted and merged without
	// overlaps. Seeking, pauses and buffering are excluded. Missing or null means
	// unavailable; an empty array means measured but no section played. Replaying a
	// section does not add another range.
	PlayedRanges []shared.PlayedRangeParam `json:"played_ranges,omitzero"`
	// Cumulative playing time per subtitle language for this attempt, null language
	// meaning subtitles off. At most 32 entries. Optional.
	SubtitlePlayingTime []shared.SubtitlePlayingTimeParam `json:"subtitle_playing_time,omitzero"`
	// Optional browser information supplied by the player.
	BrowserContext any `json:"browser_context,omitzero"`
	// Integrator-defined dimensions set once when the view starts: at most 10 keys of
	// 1 to 40 characters a-z, 0-9 and \_, each with text of 1 to 256 characters, a
	// number or a boolean. Reports compare values as text, exactly and
	// case-sensitively: numbers as their JSON text (5, 5.5; send whole numbers without
	// a decimal part, as JavaScript does), booleans as true or false. Only
	// view_started accepts custom. Never put personal data here; use user_id for user
	// identity.
	Custom shared.CustomDimensionsParam `json:"custom,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "chunkify.video.view_started".
	EventName constant.ChunkifyVideoViewStarted `json:"event_name" default:"chunkify.video.view_started"`
	// This field can be elided, and will marshal its zero value as 1.
	SchemaVersion float64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r ViewStartedParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewStartedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewStartedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewStartedParam](
		"schema_version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewStartedErrorCodeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u ViewStartedErrorCodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *ViewStartedErrorCodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ViewStartedErrorCodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

// EventIngestResponseUnion contains all possible properties and values from
// [NativeCreated], [NativeDuplicate].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventIngestResponseUnion struct {
	Accepted  bool `json:"accepted"`
	Duplicate bool `json:"duplicate"`
	JSON      struct {
		Accepted  respjson.Field
		Duplicate respjson.Field
		raw       string
	} `json:"-"`
}

func (u EventIngestResponseUnion) AsNativeCreated() (v NativeCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventIngestResponseUnion) AsNativeDuplicate() (v NativeDuplicate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventIngestResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *EventIngestResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
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
	Continent EventListParamsContinent `query:"continent,omitzero" json:"-"`
	// Exact device class. unknown matches an explicit unknown value.
	//
	// Any of "desktop", "mobile", "tablet", "unknown".
	DeviceType EventListParamsDeviceType `query:"device_type,omitzero" json:"-"`
	// Raw-event queries only. Summaries and videos reject this parameter.
	//
	// Any of "chunkify.video.view_started", "chunkify.video.view_snapshot",
	// "chunkify.video.playback_error", "chunkify.video.view_ended",
	// "chunkify.video.playback_state_changed", "chunkify.video.rendition_changed",
	// "chunkify.video.seek_completed", "chunkify.video.subtitle_changed",
	// "chunkify.video.playback_rate_changed".
	EventName EventListParamsEventName `query:"event_name,omitzero" json:"-"`
	// Sort direction. Events and videos default to asc; views default to desc. Events
	// sort by ingestion sequence. Videos and views sort by sort_by with missing values
	// last. View ties use view_id in the requested direction; video metric ties use
	// video_id ascending.
	//
	// Any of "asc", "desc".
	Order EventListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact continent code from the local lookup. Missing locations do not match.
type EventListParamsContinent string

const (
	EventListParamsContinentAf EventListParamsContinent = "AF"
	EventListParamsContinentAn EventListParamsContinent = "AN"
	EventListParamsContinentAs EventListParamsContinent = "AS"
	EventListParamsContinentEu EventListParamsContinent = "EU"
	EventListParamsContinentNa EventListParamsContinent = "NA"
	EventListParamsContinentOc EventListParamsContinent = "OC"
	EventListParamsContinentSa EventListParamsContinent = "SA"
)

// Exact device class. unknown matches an explicit unknown value.
type EventListParamsDeviceType string

const (
	EventListParamsDeviceTypeDesktop EventListParamsDeviceType = "desktop"
	EventListParamsDeviceTypeMobile  EventListParamsDeviceType = "mobile"
	EventListParamsDeviceTypeTablet  EventListParamsDeviceType = "tablet"
	EventListParamsDeviceTypeUnknown EventListParamsDeviceType = "unknown"
)

// Raw-event queries only. Summaries and videos reject this parameter.
type EventListParamsEventName string

const (
	EventListParamsEventNameChunkifyVideoViewStarted          EventListParamsEventName = "chunkify.video.view_started"
	EventListParamsEventNameChunkifyVideoViewSnapshot         EventListParamsEventName = "chunkify.video.view_snapshot"
	EventListParamsEventNameChunkifyVideoPlaybackError        EventListParamsEventName = "chunkify.video.playback_error"
	EventListParamsEventNameChunkifyVideoViewEnded            EventListParamsEventName = "chunkify.video.view_ended"
	EventListParamsEventNameChunkifyVideoPlaybackStateChanged EventListParamsEventName = "chunkify.video.playback_state_changed"
	EventListParamsEventNameChunkifyVideoRenditionChanged     EventListParamsEventName = "chunkify.video.rendition_changed"
	EventListParamsEventNameChunkifyVideoSeekCompleted        EventListParamsEventName = "chunkify.video.seek_completed"
	EventListParamsEventNameChunkifyVideoSubtitleChanged      EventListParamsEventName = "chunkify.video.subtitle_changed"
	EventListParamsEventNameChunkifyVideoPlaybackRateChanged  EventListParamsEventName = "chunkify.video.playback_rate_changed"
)

// Sort direction. Events and videos default to asc; views default to desc. Events
// sort by ingestion sequence. Videos and views sort by sort_by with missing values
// last. View ties use view_id in the requested direction; video metric ties use
// video_id ascending.
type EventListParamsOrder string

const (
	EventListParamsOrderAsc  EventListParamsOrder = "asc"
	EventListParamsOrderDesc EventListParamsOrder = "desc"
)

type EventIngestParams struct {
	// A schema version 1 playback event. Extra properties are preserved.
	Event  AnalyticsEventUnionParam
	Origin param.Opt[string] `header:"Origin,omitzero" json:"-"`
	// Any of "identity".
	ContentEncoding EventIngestParamsContentEncoding `header:"Content-Encoding,omitzero" json:"-"`
	paramObj
}

func (r EventIngestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Event)
}
func (r *EventIngestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventIngestParamsContentEncoding string

const (
	EventIngestParamsContentEncodingIdentity EventIngestParamsContentEncoding = "identity"
)
