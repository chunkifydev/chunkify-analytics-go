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

func TestErrorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Errors.List(context.TODO(), chunkifyanalytics.ErrorListParams{
		BrowserName: chunkifyanalytics.String("x"),
		City:        chunkifyanalytics.String("x"),
		Continent:   chunkifyanalytics.ErrorListParamsContinentAf,
		Country:     chunkifyanalytics.String("SE"),
		Cursor:      chunkifyanalytics.String("cursor"),
		DeviceType:  chunkifyanalytics.ErrorListParamsDeviceTypeDesktop,
		From:        chunkifyanalytics.Time(time.Now()),
		Limit:       chunkifyanalytics.Int(1),
		OsName:      chunkifyanalytics.String("x"),
		PageOrigin:  chunkifyanalytics.String("x"),
		PagePath:    chunkifyanalytics.String("x"),
		To:          chunkifyanalytics.Time(time.Now()),
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
