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

func TestViewGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Get(
		context.TODO(),
		"x",
		chunkifyanalytics.ViewGetParams{
			Cursor: chunkifyanalytics.String("cursor"),
			Limit:  chunkifyanalytics.Int(1),
		},
	)
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.List(context.TODO(), chunkifyanalytics.ViewListParams{
		BrowserName:   chunkifyanalytics.String("x"),
		City:          chunkifyanalytics.String("x"),
		Continent:     chunkifyanalytics.ViewListParamsContinentAf,
		Country:       chunkifyanalytics.String("SE"),
		Cursor:        chunkifyanalytics.String("cursor"),
		DeviceType:    chunkifyanalytics.ViewListParamsDeviceTypeDesktop,
		From:          chunkifyanalytics.Time(time.Now()),
		HasFatalError: chunkifyanalytics.Bool(true),
		Limit:         chunkifyanalytics.Int(1),
		Order:         chunkifyanalytics.ViewListParamsOrderAsc,
		OsName:        chunkifyanalytics.String("x"),
		PageOrigin:    chunkifyanalytics.String("x"),
		PagePath:      chunkifyanalytics.String("x"),
		SortBy:        chunkifyanalytics.ViewListParamsSortByStartedAt,
		To:            chunkifyanalytics.Time(time.Now()),
		VideoID:       chunkifyanalytics.String("video_id"),
		ViewID:        chunkifyanalytics.String("x"),
		ViewerID:      chunkifyanalytics.String("viewer_id"),
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
