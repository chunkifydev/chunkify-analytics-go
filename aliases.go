// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics

import (
	"github.com/chunkifydev/chunkify-analytics-go/internal/apierror"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Browser context and locally derived continent/country/city. Request IP addresses
// are never added to events or stored context.
//
// This is an alias to an internal type.
type EventContext = shared.EventContext

// Copied from browser_context.browser_name when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextBrowserNameUnion = shared.EventContextBrowserNameUnion

// Copied from browser_context.browser_version when primitive. See BrowserContext
// for the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextBrowserVersionUnion = shared.EventContextBrowserVersionUnion

// Continent code from the local City MMDB lookup: Africa, Antarctica, Asia,
// Europe, North America, Oceania or South America. Null when unknown.
//
// This is an alias to an internal type.
type EventContextContinent = shared.EventContextContinent

// Equals "AF"
const EventContextContinentAf = shared.EventContextContinentAf

// Equals "AN"
const EventContextContinentAn = shared.EventContextContinentAn

// Equals "AS"
const EventContextContinentAs = shared.EventContextContinentAs

// Equals "EU"
const EventContextContinentEu = shared.EventContextContinentEu

// Equals "NA"
const EventContextContinentNa = shared.EventContextContinentNa

// Equals "OC"
const EventContextContinentOc = shared.EventContextContinentOc

// Equals "SA"
const EventContextContinentSa = shared.EventContextContinentSa

// Copied from browser_context.device_family when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextDeviceFamilyUnion = shared.EventContextDeviceFamilyUnion

// Copied from browser_context.device_type when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextDeviceTypeUnion = shared.EventContextDeviceTypeUnion

// Copied from browser_context.language when primitive. See BrowserContext for the
// SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextLanguageUnion = shared.EventContextLanguageUnion

// Copied from browser_context.language_prefix when primitive. See BrowserContext
// for the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextLanguagePrefixUnion = shared.EventContextLanguagePrefixUnion

// Copied from browser_context.os_name when primitive. See BrowserContext for the
// SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextOsNameUnion = shared.EventContextOsNameUnion

// Copied from browser_context.os_version when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextOsVersionUnion = shared.EventContextOsVersionUnion

// Copied from browser_context.pixel_ratio when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextPixelRatioUnion = shared.EventContextPixelRatioUnion

// Copied from browser_context.screen_height_px when primitive. See BrowserContext
// for the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextScreenHeightPxUnion = shared.EventContextScreenHeightPxUnion

// Copied from browser_context.screen_width_px when primitive. See BrowserContext
// for the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextScreenWidthPxUnion = shared.EventContextScreenWidthPxUnion

// Copied from browser_context.timezone when primitive. See BrowserContext for the
// SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextTimezoneUnion = shared.EventContextTimezoneUnion

// Copied from browser_context.timezone_offset_minutes when primitive. See
// BrowserContext for the SDK type; the collector currently accepts any primitive
// here.
//
// This is an alias to an internal type.
type EventContextTimezoneOffsetMinutesUnion = shared.EventContextTimezoneOffsetMinutesUnion

// Copied from browser_context.user_agent when primitive. See BrowserContext for
// the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextUserAgentUnion = shared.EventContextUserAgentUnion

// Copied from browser_context.viewport_height_px when primitive. See
// BrowserContext for the SDK type; the collector currently accepts any primitive
// here.
//
// This is an alias to an internal type.
type EventContextViewportHeightPxUnion = shared.EventContextViewportHeightPxUnion

// Copied from browser_context.viewport_width_px when primitive. See BrowserContext
// for the SDK type; the collector currently accepts any primitive here.
//
// This is an alias to an internal type.
type EventContextViewportWidthPxUnion = shared.EventContextViewportWidthPxUnion

// View count and fraction of its documented denominator. The ratio is null when
// that denominator is zero.
//
// This is an alias to an internal type.
type MetricCountRatio = shared.MetricCountRatio

// Standard OTLP AnyValue in a preserved resource, scope or body. This collector
// does not decode these fields into the event.
//
// This is an alias to an internal type.
type OtlpAnyValue = shared.OtlpAnyValue

// This is an alias to an internal type.
type OtlpAnyValueArrayValue = shared.OtlpAnyValueArrayValue

// Signed int64. Strings avoid JavaScript precision loss.
//
// This is an alias to an internal type.
type OtlpAnyValueIntValueUnion = shared.OtlpAnyValueIntValueUnion

// This is an alias to an internal type.
type OtlpAnyValueKvlistValue = shared.OtlpAnyValueKvlistValue

// An absent or null key decodes to the empty string; duplicate decoded keys reject
// the request. An absent value decodes to null.
//
// This is an alias to an internal type.
type OtlpAttribute = shared.OtlpAttribute

// A typed log attribute. At most one value field can be set. Binary values are
// unsupported.
//
// This is an alias to an internal type.
type OtlpAttributeValueUnion = shared.OtlpAttributeValueUnion

// This is an alias to an internal type.
type OtlpAttributeValueObject = shared.OtlpAttributeValueObject

// This is an alias to an internal type.
type OtlpAttributeValueObject2 = shared.OtlpAttributeValueObject2

// This is an alias to an internal type.
type OtlpAttributeValueObject3 = shared.OtlpAttributeValueObject3

// This is an alias to an internal type.
type OtlpAttributeValueObject4 = shared.OtlpAttributeValueObject4

// This is an alias to an internal type.
type OtlpAttributeValueObject5 = shared.OtlpAttributeValueObject5

// This is an alias to an internal type.
type OtlpAttributeValueObject6 = shared.OtlpAttributeValueObject6

// This is an alias to an internal type.
type OtlpAttributeValueObject7 = shared.OtlpAttributeValueObject7

// This is an alias to an internal type.
type OtlpContextAttribute = shared.OtlpContextAttribute

// A playback event in log attributes. Duplicate attributes and invalid event
// fields reject the request.
//
// This is an alias to an internal type.
type OtlpLogRecord = shared.OtlpLogRecord

// This is an alias to an internal type.
type OtlpResource = shared.OtlpResource

// This is an alias to an internal type.
type OtlpScope = shared.OtlpScope

// This is an alias to an internal type.
type PlaybackScore = shared.PlaybackScore

// This is an alias to an internal type.
type PlayedRange = shared.PlayedRange

// This is an alias to an internal type.
type PlayedRangeParam = shared.PlayedRangeParam

// Furthest progress on the selected final event or latest snapshot. Only views
// with a positive known media duration and a progress ratio in [0,1] are samples.
// Each milestone ratio uses sample_views. Seeking can count; this is not watched
// coverage or audience retention.
//
// This is an alias to an internal type.
type ProgressReached = shared.ProgressReached

// This is an alias to an internal type.
type RenditionPlayingTime = shared.RenditionPlayingTime

// This is an alias to an internal type.
type RenditionPlayingTimeParam = shared.RenditionPlayingTimeParam

// Mutually exclusive outcomes. An explicit startup_outcome on the selected final
// event takes precedence. Otherwise an available startup time, including zero, or
// positive playing time on the selected measurement proves started. All other
// views are unknown. sample_views excludes unknown; every ratio uses all selected
// attempts, including unknown. Missing final delivery does not prove an exit or
// failure.
//
// This is an alias to an internal type.
type StartupOutcomes = shared.StartupOutcomes

// Legacy context present for an event received through OTLP by an earlier
// collector version. Existing context remains readable; new ingestion accepts
// native JSON events only.
//
// This is an alias to an internal type.
type StoredOtlpContext = shared.StoredOtlpContext

// Original resource group excluding scopeLogs.
//
// This is an alias to an internal type.
type StoredOtlpContextResourceLogs = shared.StoredOtlpContextResourceLogs

// Original scope group excluding logRecords.
//
// This is an alias to an internal type.
type StoredOtlpContextScopeLogs = shared.StoredOtlpContextScopeLogs

// Playing time in milliseconds from one selected final event or latest snapshot
// per view. Measured zero is included; missing measurements are excluded.
// Percentiles use linear interpolation and are null without samples. This sample
// denominator differs from average_playing_time_ms, which uses all attempts.
//
// This is an alias to an internal type.
type WatchTimeDistribution = shared.WatchTimeDistribution

// Mutually exclusive playing-time buckets: [0,10s), [10s,30s), [30s,60s), [1m,5m),
// [5m,infinity). Counts sum to sample_views.
//
// This is an alias to an internal type.
type WatchTimeDistributionDistribution = shared.WatchTimeDistributionDistribution
