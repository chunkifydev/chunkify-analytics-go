// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/chunkifydev/chunkify-analytics-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type ChunkifyVideoPlaybackError string        // Always "chunkify.video.playback_error"
type ChunkifyVideoPlaybackStateChanged string // Always "chunkify.video.playback_state_changed"
type ChunkifyVideoRenditionChanged string     // Always "chunkify.video.rendition_changed"
type ChunkifyVideoSeekCompleted string        // Always "chunkify.video.seek_completed"
type ChunkifyVideoViewEnded string            // Always "chunkify.video.view_ended"
type ChunkifyVideoViewSnapshot string         // Always "chunkify.video.view_snapshot"
type ChunkifyVideoViewStarted string          // Always "chunkify.video.view_started"
type Engagement string                        // Always "engagement"
type Ok string                                // Always "ok"
type Quality string                           // Always "quality"
type Summary string                           // Always "summary"
type TextCsv string                           // Always "text/csv"
type Utc string                               // Always "UTC"

func (c ChunkifyVideoPlaybackError) Default() ChunkifyVideoPlaybackError {
	return "chunkify.video.playback_error"
}
func (c ChunkifyVideoPlaybackStateChanged) Default() ChunkifyVideoPlaybackStateChanged {
	return "chunkify.video.playback_state_changed"
}
func (c ChunkifyVideoRenditionChanged) Default() ChunkifyVideoRenditionChanged {
	return "chunkify.video.rendition_changed"
}
func (c ChunkifyVideoSeekCompleted) Default() ChunkifyVideoSeekCompleted {
	return "chunkify.video.seek_completed"
}
func (c ChunkifyVideoViewEnded) Default() ChunkifyVideoViewEnded { return "chunkify.video.view_ended" }
func (c ChunkifyVideoViewSnapshot) Default() ChunkifyVideoViewSnapshot {
	return "chunkify.video.view_snapshot"
}
func (c ChunkifyVideoViewStarted) Default() ChunkifyVideoViewStarted {
	return "chunkify.video.view_started"
}
func (c Engagement) Default() Engagement { return "engagement" }
func (c Ok) Default() Ok                 { return "ok" }
func (c Quality) Default() Quality       { return "quality" }
func (c Summary) Default() Summary       { return "summary" }
func (c TextCsv) Default() TextCsv       { return "text/csv" }
func (c Utc) Default() Utc               { return "UTC" }

func (c ChunkifyVideoPlaybackError) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ChunkifyVideoPlaybackStateChanged) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c ChunkifyVideoRenditionChanged) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ChunkifyVideoSeekCompleted) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ChunkifyVideoViewEnded) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ChunkifyVideoViewSnapshot) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ChunkifyVideoViewStarted) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Engagement) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Ok) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c Quality) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Summary) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c TextCsv) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Utc) MarshalJSON() ([]byte, error)                               { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
