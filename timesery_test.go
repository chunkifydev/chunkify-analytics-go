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
)

func TestTimeseryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Timeseries.Get(context.TODO(), chunkifyanalytics.TimeseryGetParams{
		From:        time.Now(),
		Interval:    chunkifyanalytics.TimeseryGetParamsIntervalHour,
		To:          time.Now(),
		BrowserName: chunkifyanalytics.String("x"),
		City:        chunkifyanalytics.String("x"),
		Continent:   chunkifyanalytics.TimeseryGetParamsContinentAf,
		Country:     chunkifyanalytics.String("SE"),
		DeviceType:  chunkifyanalytics.TimeseryGetParamsDeviceTypeDesktop,
		Kind:        chunkifyanalytics.TimeseryGetParamsKindSummary,
		OsName:      chunkifyanalytics.String("x"),
		PageOrigin:  chunkifyanalytics.String("x"),
		PagePath:    chunkifyanalytics.String("x"),
		VideoID:     chunkifyanalytics.String("video_id"),
		ViewID:      chunkifyanalytics.String("x"),
		ViewerID:    chunkifyanalytics.String("viewer_id"),
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
