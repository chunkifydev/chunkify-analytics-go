// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/chunkifydev/chunkify-analytics-go"
	"github.com/chunkifydev/chunkify-analytics-go/internal/testutil"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/shared"
)

func TestEventListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkifyanalytics.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPublicKey("My Public Key"),
		option.WithSecretKey("My Secret Key"),
		option.WithAdminKey("My Admin Key"),
	)
	_, err := client.Events.List(context.TODO(), chunkifyanalytics.EventListParams{
		BrowserName:      chunkifyanalytics.String("x"),
		City:             chunkifyanalytics.String("x"),
		Continent:        chunkifyanalytics.EventListParamsContinentAf,
		Country:          chunkifyanalytics.String("SE"),
		Cursor:           chunkifyanalytics.String("cursor"),
		DeviceType:       chunkifyanalytics.EventListParamsDeviceTypeDesktop,
		EventName:        chunkifyanalytics.EventListParamsEventNameChunkifyVideoViewStarted,
		From:             chunkifyanalytics.Time(time.Now()),
		Limit:            chunkifyanalytics.Int(1),
		Order:            chunkifyanalytics.EventListParamsOrderAsc,
		OsName:           chunkifyanalytics.String("x"),
		PageOrigin:       chunkifyanalytics.String("x"),
		PagePath:         chunkifyanalytics.String("x"),
		PlaybackRate:     chunkifyanalytics.String("269125115713"),
		SubtitleLanguage: chunkifyanalytics.String("P"),
		To:               chunkifyanalytics.Time(time.Now()),
		UserID:           chunkifyanalytics.String("x"),
		VideoID:          chunkifyanalytics.String("P"),
		VideoTitle:       chunkifyanalytics.String("x"),
		ViewID:           chunkifyanalytics.String("x"),
		ViewerID:         chunkifyanalytics.String("x"),
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventIngestWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkifyanalytics.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPublicKey("My Public Key"),
		option.WithSecretKey("My Secret Key"),
		option.WithAdminKey("My Admin Key"),
	)
	_, err := client.Events.Ingest(context.TODO(), chunkifyanalytics.EventIngestParams{
		Event: chunkifyanalytics.AnalyticsEventUnionParam{
			OfViewStarted: &chunkifyanalytics.ViewStartedParam{
				ConnectorName:            "html-media",
				ConnectorVersion:         param.Null[string](),
				DeliveryHostname:         param.Null[string](),
				DeliveryRequestCount:     param.Null[float64](),
				DeliveryThroughputBps:    param.Null[float64](),
				DeliveryTransferredBytes: param.Null[float64](),
				DroppedVideoFrames:       param.Null[float64](),
				Environment:              param.Null[string](),
				ErrorCategory:            param.Null[chunkifyanalytics.ViewStartedErrorCategory](),
				ErrorCode: chunkifyanalytics.ViewStartedErrorCodeUnionParam{
					OfString: chunkifyanalytics.String("string"),
				},
				ErrorFatal:             param.Null[bool](),
				ErrorPlayheadSeconds:   param.Null[float64](),
				EventID:                "35c9ab04-7c09-4c09-ae70-a214f4a637c0",
				FirstVideoFrameTimeMs:  param.Null[float64](),
				FurthestProgressRatio:  param.Null[float64](),
				MaximumPlayheadSeconds: 0,
				MediaDurationSeconds:   param.Null[float64](),
				MediaHeightPx:          param.Null[float64](),
				MediaWidthPx:           param.Null[float64](),
				OccurredAt:             time.Now(),
				PlayerHeightPx:         param.Null[float64](),
				PlayerName:             param.Null[string](),
				PlayerVersion:          param.Null[string](),
				PlayerWidthPx:          param.Null[float64](),
				PlayingTimeMs:          0,
				Reason:                 param.Null[chunkifyanalytics.ViewStartedReason](),
				RebufferCount:          0,
				RebufferDurationMs:     0,
				RebufferRatio:          param.Null[float64](),
				RenditionPlayingTime: []shared.RenditionPlayingTimeParam{{
					BitrateBps:    chunkifyanalytics.Float(0),
					Codec:         chunkifyanalytics.String("codec"),
					FrameRate:     chunkifyanalytics.Float(0),
					Height:        chunkifyanalytics.Float(0),
					PlayingTimeMs: 0,
					Width:         chunkifyanalytics.Float(0),
				}},
				RenditionSwitchCount:        param.Null[float64](),
				SchemaVersion:               1,
				SDKName:                     "@chunkify/player-analytics",
				SDKVersion:                  "0.3.0",
				SeekCount:                   0,
				SeekDurationMs:              0,
				SelectedRenditionBitrateBps: param.Null[float64](),
				SelectedRenditionCodec:      param.Null[string](),
				SelectedRenditionHeight:     param.Null[float64](),
				SelectedRenditionWidth:      param.Null[float64](),
				Sequence:                    1,
				StartupOutcome:              param.Null[chunkifyanalytics.ViewStartedStartupOutcome](),
				StartupTimeMs:               param.Null[float64](),
				TotalVideoFrames:            param.Null[float64](),
				VideoID:                     "demo-video",
				ViewID:                      "616992d2-6db6-4578-8361-3f8813e963be",
				ViewerID:                    param.Null[string](),
				BrowserContext:              map[string]any{},
				Custom: shared.CustomDimensionsParam{
					"plan": shared.CustomDimensionUnionParam{
						OfString: chunkifyanalytics.String("premium"),
					},
					"ab_test": shared.CustomDimensionUnionParam{
						OfString: chunkifyanalytics.String("B"),
					},
					"seats": shared.CustomDimensionUnionParam{
						OfFloat: chunkifyanalytics.Float(5),
					},
					"trial": shared.CustomDimensionUnionParam{
						OfBool: chunkifyanalytics.Bool(false),
					},
				},
				PlaybackRate: chunkifyanalytics.Float(0.01),
				PlaybackRatePlayingTime: []shared.PlaybackRatePlayingTimeParam{{
					PlayingTimeMs: 0,
					Rate:          0.01,
				}},
				PlayedRanges: []shared.PlayedRangeParam{{
					EndSeconds:   0,
					StartSeconds: 0,
				}},
				PlayheadSeconds:  chunkifyanalytics.Float(0),
				SubtitleLanguage: chunkifyanalytics.String("P"),
				SubtitlePlayingTime: []shared.SubtitlePlayingTimeParam{{
					Language:      chunkifyanalytics.String("P"),
					PlayingTimeMs: 0,
				}},
				UserID:          chunkifyanalytics.String("x"),
				VideoTitle:      chunkifyanalytics.String("x"),
				VisibilityState: chunkifyanalytics.Opt(chunkifyanalytics.ViewStartedVisibilityStateVisible),
			},
		},
		ContentEncoding: chunkifyanalytics.EventIngestParamsContentEncodingIdentity,
		Origin:          chunkifyanalytics.String("Origin"),
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
