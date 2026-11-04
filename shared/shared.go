// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Browser context and locally derived continent/country/city. Request IP addresses
// are never added to events or stored context.
type EventContext struct {
	SchemaVersion float64 `json:"schema_version" api:"required"`
	// Copied from browser_context.browser_name when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	BrowserName EventContextBrowserNameUnion `json:"browser_name" api:"nullable"`
	// Copied from browser_context.browser_version when primitive. See BrowserContext
	// for the SDK type; the collector currently accepts any primitive here.
	BrowserVersion EventContextBrowserVersionUnion `json:"browser_version" api:"nullable"`
	// Approximate city name in English from the local City MMDB lookup. Null when
	// unknown. Location displays using the bundled DB-IP data must link to
	// https://db-ip.com.
	City string `json:"city" api:"nullable"`
	// Continent code from the local City MMDB lookup: Africa, Antarctica, Asia,
	// Europe, North America, Oceania or South America. Null when unknown.
	//
	// Any of "AF", "AN", "AS", "EU", "NA", "OC", "SA".
	Continent EventContextContinent `json:"continent" api:"nullable"`
	// Country ISO 3166-1 alpha-2 code from the local City MMDB lookup. Null when
	// unknown. The request IP is discarded.
	Country string `json:"country" api:"nullable"`
	// Copied from browser_context.device_family when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	DeviceFamily EventContextDeviceFamilyUnion `json:"device_family" api:"nullable"`
	// Copied from browser_context.device_type when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	DeviceType EventContextDeviceTypeUnion `json:"device_type" api:"nullable"`
	// Copied from browser_context.language when primitive. See BrowserContext for the
	// SDK type; the collector currently accepts any primitive here.
	Language EventContextLanguageUnion `json:"language" api:"nullable"`
	// Copied from browser_context.language_prefix when primitive. See BrowserContext
	// for the SDK type; the collector currently accepts any primitive here.
	LanguagePrefix EventContextLanguagePrefixUnion `json:"language_prefix" api:"nullable"`
	// Copied from browser_context.os_name when primitive. See BrowserContext for the
	// SDK type; the collector currently accepts any primitive here.
	OsName EventContextOsNameUnion `json:"os_name" api:"nullable"`
	// Copied from browser_context.os_version when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	OsVersion EventContextOsVersionUnion `json:"os_version" api:"nullable"`
	// Current document HTTP(S) origin, without credentials. Null when unavailable or
	// longer than 2048 characters.
	PageOrigin string `json:"page_origin" api:"nullable"`
	// Current document pathname, without query parameters or fragment. Null when
	// unavailable or longer than 2048 characters. Paths may contain
	// application-defined identifiers.
	PagePath string `json:"page_path" api:"nullable"`
	// Copied from browser_context.pixel_ratio when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	PixelRatio EventContextPixelRatioUnion `json:"pixel_ratio" api:"nullable"`
	// Copied from browser_context.screen_height_px when primitive. See BrowserContext
	// for the SDK type; the collector currently accepts any primitive here.
	ScreenHeightPx EventContextScreenHeightPxUnion `json:"screen_height_px" api:"nullable"`
	// Copied from browser_context.screen_width_px when primitive. See BrowserContext
	// for the SDK type; the collector currently accepts any primitive here.
	ScreenWidthPx EventContextScreenWidthPxUnion `json:"screen_width_px" api:"nullable"`
	// Copied from browser_context.timezone when primitive. See BrowserContext for the
	// SDK type; the collector currently accepts any primitive here.
	Timezone EventContextTimezoneUnion `json:"timezone" api:"nullable"`
	// Copied from browser_context.timezone_offset_minutes when primitive. See
	// BrowserContext for the SDK type; the collector currently accepts any primitive
	// here.
	TimezoneOffsetMinutes EventContextTimezoneOffsetMinutesUnion `json:"timezone_offset_minutes" api:"nullable"`
	// Copied from browser_context.user_agent when primitive. See BrowserContext for
	// the SDK type; the collector currently accepts any primitive here.
	UserAgent EventContextUserAgentUnion `json:"user_agent" api:"nullable"`
	// Copied from browser_context.viewport_height_px when primitive. See
	// BrowserContext for the SDK type; the collector currently accepts any primitive
	// here.
	ViewportHeightPx EventContextViewportHeightPxUnion `json:"viewport_height_px" api:"nullable"`
	// Copied from browser_context.viewport_width_px when primitive. See BrowserContext
	// for the SDK type; the collector currently accepts any primitive here.
	ViewportWidthPx EventContextViewportWidthPxUnion `json:"viewport_width_px" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SchemaVersion         respjson.Field
		BrowserName           respjson.Field
		BrowserVersion        respjson.Field
		City                  respjson.Field
		Continent             respjson.Field
		Country               respjson.Field
		DeviceFamily          respjson.Field
		DeviceType            respjson.Field
		Language              respjson.Field
		LanguagePrefix        respjson.Field
		OsName                respjson.Field
		OsVersion             respjson.Field
		PageOrigin            respjson.Field
		PagePath              respjson.Field
		PixelRatio            respjson.Field
		ScreenHeightPx        respjson.Field
		ScreenWidthPx         respjson.Field
		Timezone              respjson.Field
		TimezoneOffsetMinutes respjson.Field
		UserAgent             respjson.Field
		ViewportHeightPx      respjson.Field
		ViewportWidthPx       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventContext) RawJSON() string { return r.JSON.raw }
func (r *EventContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextBrowserNameUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextBrowserNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextBrowserNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextBrowserNameUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextBrowserNameUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextBrowserNameUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextBrowserNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextBrowserVersionUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextBrowserVersionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextBrowserVersionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextBrowserVersionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextBrowserVersionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextBrowserVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextBrowserVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Continent code from the local City MMDB lookup: Africa, Antarctica, Asia,
// Europe, North America, Oceania or South America. Null when unknown.
type EventContextContinent string

const (
	EventContextContinentAf EventContextContinent = "AF"
	EventContextContinentAn EventContextContinent = "AN"
	EventContextContinentAs EventContextContinent = "AS"
	EventContextContinentEu EventContextContinent = "EU"
	EventContextContinentNa EventContextContinent = "NA"
	EventContextContinentOc EventContextContinent = "OC"
	EventContextContinentSa EventContextContinent = "SA"
)

// EventContextDeviceFamilyUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextDeviceFamilyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextDeviceFamilyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextDeviceFamilyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextDeviceFamilyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextDeviceFamilyUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextDeviceFamilyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextDeviceTypeUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextDeviceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextDeviceTypeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextDeviceTypeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextDeviceTypeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextDeviceTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextDeviceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextLanguageUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextLanguageUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextLanguageUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextLanguageUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextLanguageUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextLanguageUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextLanguageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextLanguagePrefixUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextLanguagePrefixUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextLanguagePrefixUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextLanguagePrefixUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextLanguagePrefixUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextLanguagePrefixUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextLanguagePrefixUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextOsNameUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextOsNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextOsNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextOsNameUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextOsNameUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextOsNameUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextOsNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextOsVersionUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextOsVersionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextOsVersionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextOsVersionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextOsVersionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextOsVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextOsVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextPixelRatioUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextPixelRatioUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextPixelRatioUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextPixelRatioUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextPixelRatioUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextPixelRatioUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextPixelRatioUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextScreenHeightPxUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextScreenHeightPxUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextScreenHeightPxUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextScreenHeightPxUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextScreenHeightPxUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextScreenHeightPxUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextScreenHeightPxUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextScreenWidthPxUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextScreenWidthPxUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextScreenWidthPxUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextScreenWidthPxUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextScreenWidthPxUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextScreenWidthPxUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextScreenWidthPxUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextTimezoneUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextTimezoneUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextTimezoneUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextTimezoneUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextTimezoneUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextTimezoneUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextTimezoneUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextTimezoneOffsetMinutesUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextTimezoneOffsetMinutesUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextTimezoneOffsetMinutesUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextTimezoneOffsetMinutesUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextTimezoneOffsetMinutesUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextTimezoneOffsetMinutesUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextTimezoneOffsetMinutesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextUserAgentUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextUserAgentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextUserAgentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextUserAgentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextUserAgentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextUserAgentUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextUserAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextViewportHeightPxUnion contains all possible properties and values
// from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextViewportHeightPxUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextViewportHeightPxUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextViewportHeightPxUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextViewportHeightPxUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextViewportHeightPxUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextViewportHeightPxUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventContextViewportWidthPxUnion contains all possible properties and values
// from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventContextViewportWidthPxUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventContextViewportWidthPxUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextViewportWidthPxUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventContextViewportWidthPxUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventContextViewportWidthPxUnion) RawJSON() string { return u.JSON.raw }

func (r *EventContextViewportWidthPxUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// View count and fraction of its documented denominator. The ratio is null when
// that denominator is zero.
type MetricCountRatio struct {
	Ratio float64 `json:"ratio" api:"required"`
	Views int64   `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ratio       respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricCountRatio) RawJSON() string { return r.JSON.raw }
func (r *MetricCountRatio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard OTLP AnyValue in a preserved resource, scope or body. This collector
// does not decode these fields into the event.
type OtlpAnyValue struct {
	ArrayValue  OtlpAnyValueArrayValue `json:"arrayValue" api:"nullable"`
	BoolValue   bool                   `json:"boolValue" api:"nullable"`
	BytesValue  string                 `json:"bytesValue" api:"nullable"`
	DoubleValue float64                `json:"doubleValue" api:"nullable"`
	// Signed int64. Strings avoid JavaScript precision loss.
	IntValue    OtlpAnyValueIntValueUnion `json:"intValue" api:"nullable"`
	KvlistValue OtlpAnyValueKvlistValue   `json:"kvlistValue" api:"nullable"`
	StringValue string                    `json:"stringValue" api:"nullable"`
	ExtraFields map[string]any            `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAnyValue) RawJSON() string { return r.JSON.raw }
func (r *OtlpAnyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAnyValueArrayValue struct {
	Values []OtlpAnyValue `json:"values" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAnyValueArrayValue) RawJSON() string { return r.JSON.raw }
func (r *OtlpAnyValueArrayValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OtlpAnyValueIntValueUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type OtlpAnyValueIntValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u OtlpAnyValueIntValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAnyValueIntValueUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OtlpAnyValueIntValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OtlpAnyValueIntValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAnyValueKvlistValue struct {
	Values []OtlpContextAttribute `json:"values" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAnyValueKvlistValue) RawJSON() string { return r.JSON.raw }
func (r *OtlpAnyValueKvlistValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An absent or null key decodes to the empty string; duplicate decoded keys reject
// the request. An absent value decodes to null.
type OtlpAttribute struct {
	Key string `json:"key" api:"nullable"`
	// A typed log attribute. At most one value field can be set. Binary values are
	// unsupported.
	Value OtlpAttributeValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttribute) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OtlpAttributeValueUnion contains all possible properties and values from
// [OtlpAttributeValueObject], [OtlpAttributeValueObject2],
// [OtlpAttributeValueObject3], [OtlpAttributeValueObject4],
// [OtlpAttributeValueObject5], [OtlpAttributeValueObject6],
// [OtlpAttributeValueObject7].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OtlpAttributeValueUnion struct {
	ArrayValue  any `json:"arrayValue"`
	BoolValue   any `json:"boolValue"`
	BytesValue  any `json:"bytesValue"`
	DoubleValue any `json:"doubleValue"`
	IntValue    any `json:"intValue"`
	KvlistValue any `json:"kvlistValue"`
	StringValue any `json:"stringValue"`
	JSON        struct {
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		raw         string
	} `json:"-"`
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject() (v OtlpAttributeValueObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject2() (v OtlpAttributeValueObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject3() (v OtlpAttributeValueObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject4() (v OtlpAttributeValueObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject5() (v OtlpAttributeValueObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject6() (v OtlpAttributeValueObject6) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OtlpAttributeValueUnion) AsOtlpAttributeValueObject7() (v OtlpAttributeValueObject7) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OtlpAttributeValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OtlpAttributeValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject struct {
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject2 struct {
	StringValue any `json:"stringValue" api:"required"`
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StringValue respjson.Field
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject2) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject3 struct {
	BoolValue   any `json:"boolValue" api:"required"`
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoolValue   respjson.Field
		ArrayValue  respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject3) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject4 struct {
	IntValue    any `json:"intValue" api:"required"`
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntValue    respjson.Field
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject4) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject5 struct {
	DoubleValue any `json:"doubleValue" api:"required"`
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DoubleValue respjson.Field
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject5) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject6 struct {
	ArrayValue  any `json:"arrayValue" api:"required"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	KvlistValue any `json:"kvlistValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		KvlistValue respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject6) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject6) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpAttributeValueObject7 struct {
	KvlistValue any `json:"kvlistValue" api:"required"`
	ArrayValue  any `json:"arrayValue" api:"nullable"`
	BoolValue   any `json:"boolValue" api:"nullable"`
	BytesValue  any `json:"bytesValue" api:"nullable"`
	DoubleValue any `json:"doubleValue" api:"nullable"`
	IntValue    any `json:"intValue" api:"nullable"`
	StringValue any `json:"stringValue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KvlistValue respjson.Field
		ArrayValue  respjson.Field
		BoolValue   respjson.Field
		BytesValue  respjson.Field
		DoubleValue respjson.Field
		IntValue    respjson.Field
		StringValue respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpAttributeValueObject7) RawJSON() string { return r.JSON.raw }
func (r *OtlpAttributeValueObject7) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpContextAttribute struct {
	Key string `json:"key" api:"required"`
	// Standard OTLP AnyValue in a preserved resource, scope or body. This collector
	// does not decode these fields into the event.
	Value *OtlpAnyValue `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpContextAttribute) RawJSON() string { return r.JSON.raw }
func (r *OtlpContextAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A playback event in log attributes. Duplicate attributes and invalid event
// fields reject the request.
type OtlpLogRecord struct {
	Attributes []OtlpAttribute `json:"attributes" api:"required"`
	// Standard OTLP AnyValue in a preserved resource, scope or body. This collector
	// does not decode these fields into the event.
	Body                   OtlpAnyValue `json:"body" api:"nullable"`
	DroppedAttributesCount int64        `json:"droppedAttributesCount" api:"nullable"`
	EventName              string       `json:"eventName" api:"nullable"`
	Flags                  int64        `json:"flags" api:"nullable"`
	ObservedTimeUnixNano   string       `json:"observedTimeUnixNano" api:"nullable"`
	SeverityNumber         int64        `json:"severityNumber" api:"nullable"`
	SeverityText           string       `json:"severityText" api:"nullable"`
	SpanID                 string       `json:"spanId" api:"nullable"`
	TimeUnixNano           string       `json:"timeUnixNano" api:"nullable"`
	TraceID                string       `json:"traceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes             respjson.Field
		Body                   respjson.Field
		DroppedAttributesCount respjson.Field
		EventName              respjson.Field
		Flags                  respjson.Field
		ObservedTimeUnixNano   respjson.Field
		SeverityNumber         respjson.Field
		SeverityText           respjson.Field
		SpanID                 respjson.Field
		TimeUnixNano           respjson.Field
		TraceID                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpLogRecord) RawJSON() string { return r.JSON.raw }
func (r *OtlpLogRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpResource struct {
	Attributes             []OtlpContextAttribute `json:"attributes" api:"nullable"`
	DroppedAttributesCount int64                  `json:"droppedAttributesCount" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes             respjson.Field
		DroppedAttributesCount respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpResource) RawJSON() string { return r.JSON.raw }
func (r *OtlpResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtlpScope struct {
	Attributes             []OtlpContextAttribute `json:"attributes" api:"nullable"`
	DroppedAttributesCount int64                  `json:"droppedAttributesCount" api:"nullable"`
	Name                   string                 `json:"name" api:"nullable"`
	Version                string                 `json:"version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes             respjson.Field
		DroppedAttributesCount respjson.Field
		Name                   respjson.Field
		Version                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtlpScope) RawJSON() string { return r.JSON.raw }
func (r *OtlpScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaybackScore struct {
	// Sessions contributing a score, including recorded failures. A per-session result
	// is zero or one.
	SampleViews int64 `json:"sample_views" api:"required"`
	// Mean of measured per-session playback scores, or null without scored sessions.
	// Each session uses startup = 8 / (8 + startup_time_ms / 1000), smoothness =
	// exp(-10 _ buffering_ratio), and scores 100 _ startup^0.30 \* smoothness^0.70.
	// Buffering ratio is rebuffer duration divided by playing plus rebuffer duration.
	// The 30% startup and 70% smoothness weights are heuristic. A recorded fatal error
	// or failed_before_start outcome scores zero. Other sessions require positive
	// playing time and known startup and rebuffer durations; ended_before_start and
	// unknown outcomes are unscored. Buffering excludes startup, pauses and seeks. The
	// score does not use interruption count or measure picture quality.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SampleViews respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaybackScore) RawJSON() string { return r.JSON.raw }
func (r *PlaybackScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayedRange struct {
	// Exclusive media position in seconds, greater than start_seconds.
	EndSeconds float64 `json:"end_seconds" api:"required"`
	// Inclusive media position in seconds.
	StartSeconds float64 `json:"start_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndSeconds   respjson.Field
		StartSeconds respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayedRange) RawJSON() string { return r.JSON.raw }
func (r *PlayedRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlayedRange to a PlayedRangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlayedRangeParam.Overrides()
func (r PlayedRange) ToParam() PlayedRangeParam {
	return param.Override[PlayedRangeParam](json.RawMessage(r.RawJSON()))
}

// The properties EndSeconds, StartSeconds are required.
type PlayedRangeParam struct {
	// Exclusive media position in seconds, greater than start_seconds.
	EndSeconds float64 `json:"end_seconds" api:"required"`
	// Inclusive media position in seconds.
	StartSeconds float64 `json:"start_seconds" api:"required"`
	paramObj
}

func (r PlayedRangeParam) MarshalJSON() (data []byte, err error) {
	type shadow PlayedRangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayedRangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Furthest progress on the selected final event or latest snapshot. Only views
// with a positive known media duration and a progress ratio in [0,1] are samples.
// Each milestone ratio uses sample_views. Seeking can count; this is not watched
// coverage or audience retention.
type ProgressReached struct {
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	At100Percent MetricCountRatio `json:"at_100_percent" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	At25Percent MetricCountRatio `json:"at_25_percent" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	At50Percent MetricCountRatio `json:"at_50_percent" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	At75Percent MetricCountRatio `json:"at_75_percent" api:"required"`
	SampleViews int64            `json:"sample_views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At100Percent respjson.Field
		At25Percent  respjson.Field
		At50Percent  respjson.Field
		At75Percent  respjson.Field
		SampleViews  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgressReached) RawJSON() string { return r.JSON.raw }
func (r *ProgressReached) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RenditionPlayingTime struct {
	BitrateBps    float64 `json:"bitrate_bps" api:"required"`
	Codec         string  `json:"codec" api:"required"`
	FrameRate     float64 `json:"frame_rate" api:"required"`
	Height        float64 `json:"height" api:"required"`
	PlayingTimeMs float64 `json:"playing_time_ms" api:"required"`
	Width         float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BitrateBps    respjson.Field
		Codec         respjson.Field
		FrameRate     respjson.Field
		Height        respjson.Field
		PlayingTimeMs respjson.Field
		Width         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenditionPlayingTime) RawJSON() string { return r.JSON.raw }
func (r *RenditionPlayingTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RenditionPlayingTime to a RenditionPlayingTimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RenditionPlayingTimeParam.Overrides()
func (r RenditionPlayingTime) ToParam() RenditionPlayingTimeParam {
	return param.Override[RenditionPlayingTimeParam](json.RawMessage(r.RawJSON()))
}

// The properties BitrateBps, Codec, FrameRate, Height, PlayingTimeMs, Width are
// required.
type RenditionPlayingTimeParam struct {
	BitrateBps    param.Opt[float64] `json:"bitrate_bps,omitzero" api:"required"`
	Codec         param.Opt[string]  `json:"codec,omitzero" api:"required"`
	FrameRate     param.Opt[float64] `json:"frame_rate,omitzero" api:"required"`
	Height        param.Opt[float64] `json:"height,omitzero" api:"required"`
	Width         param.Opt[float64] `json:"width,omitzero" api:"required"`
	PlayingTimeMs float64            `json:"playing_time_ms" api:"required"`
	paramObj
}

func (r RenditionPlayingTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow RenditionPlayingTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RenditionPlayingTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mutually exclusive outcomes. An explicit startup_outcome on the selected final
// event takes precedence. Otherwise an available startup time, including zero, or
// positive playing time on the selected measurement proves started. All other
// views are unknown. sample_views excludes unknown; every ratio uses all selected
// attempts, including unknown. Missing final delivery does not prove an exit or
// failure.
type StartupOutcomes struct {
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	EndedBeforeStart MetricCountRatio `json:"ended_before_start" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	FailedBeforeStart MetricCountRatio `json:"failed_before_start" api:"required"`
	SampleViews       int64            `json:"sample_views" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	Started MetricCountRatio `json:"started" api:"required"`
	// View count and fraction of its documented denominator. The ratio is null when
	// that denominator is zero.
	Unknown MetricCountRatio `json:"unknown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndedBeforeStart  respjson.Field
		FailedBeforeStart respjson.Field
		SampleViews       respjson.Field
		Started           respjson.Field
		Unknown           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StartupOutcomes) RawJSON() string { return r.JSON.raw }
func (r *StartupOutcomes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legacy context present for an event received through OTLP by an earlier
// collector version. Existing context remains readable; new ingestion accepts
// native JSON events only.
type StoredOtlpContext struct {
	// A playback event in log attributes. Duplicate attributes and invalid event
	// fields reject the request.
	LogRecord OtlpLogRecord `json:"log_record" api:"required"`
	// Original resource group excluding scopeLogs.
	ResourceLogs StoredOtlpContextResourceLogs `json:"resource_logs" api:"required"`
	// Original scope group excluding logRecords.
	ScopeLogs StoredOtlpContextScopeLogs `json:"scope_logs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogRecord    respjson.Field
		ResourceLogs respjson.Field
		ScopeLogs    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoredOtlpContext) RawJSON() string { return r.JSON.raw }
func (r *StoredOtlpContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Original resource group excluding scopeLogs.
type StoredOtlpContextResourceLogs struct {
	Resource  OtlpResource `json:"resource" api:"nullable"`
	SchemaURL string       `json:"schemaUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Resource    respjson.Field
		SchemaURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoredOtlpContextResourceLogs) RawJSON() string { return r.JSON.raw }
func (r *StoredOtlpContextResourceLogs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Original scope group excluding logRecords.
type StoredOtlpContextScopeLogs struct {
	SchemaURL string    `json:"schemaUrl" api:"nullable"`
	Scope     OtlpScope `json:"scope" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SchemaURL   respjson.Field
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoredOtlpContextScopeLogs) RawJSON() string { return r.JSON.raw }
func (r *StoredOtlpContextScopeLogs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Playing time in milliseconds from one selected final event or latest snapshot
// per view. Measured zero is included; missing measurements are excluded.
// Percentiles use linear interpolation and are null without samples. This sample
// denominator differs from average_playing_time_ms, which uses all attempts.
type WatchTimeDistribution struct {
	// Mutually exclusive playing-time buckets: [0,10s), [10s,30s), [30s,60s), [1m,5m),
	// [5m,infinity). Counts sum to sample_views.
	Distribution WatchTimeDistributionDistribution `json:"distribution" api:"required"`
	Median       float64                           `json:"median" api:"required"`
	P95          float64                           `json:"p95" api:"required"`
	SampleViews  int64                             `json:"sample_views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distribution respjson.Field
		Median       respjson.Field
		P95          respjson.Field
		SampleViews  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WatchTimeDistribution) RawJSON() string { return r.JSON.raw }
func (r *WatchTimeDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mutually exclusive playing-time buckets: [0,10s), [10s,30s), [30s,60s), [1m,5m),
// [5m,infinity). Counts sum to sample_views.
type WatchTimeDistributionDistribution struct {
	AtLeast5Minutes   int64 `json:"at_least_5_minutes" api:"required"`
	From1To5Minutes   int64 `json:"from_1_to_5_minutes" api:"required"`
	From10To30Seconds int64 `json:"from_10_to_30_seconds" api:"required"`
	From30To60Seconds int64 `json:"from_30_to_60_seconds" api:"required"`
	Under10Seconds    int64 `json:"under_10_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtLeast5Minutes   respjson.Field
		From1To5Minutes   respjson.Field
		From10To30Seconds respjson.Field
		From30To60Seconds respjson.Field
		Under10Seconds    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WatchTimeDistributionDistribution) RawJSON() string { return r.JSON.raw }
func (r *WatchTimeDistributionDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
