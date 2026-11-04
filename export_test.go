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

func TestExportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.New(context.TODO(), chunkifyanalytics.ExportNewParams{
		ExportRequest: chunkifyanalytics.ExportRequestParam{
			Resource: chunkifyanalytics.ExportRequestResourceVideos,
			Columns:  []string{"x"},
			Query: chunkifyanalytics.ExportQueryParam{
				BrowserName:   chunkifyanalytics.String("x"),
				City:          chunkifyanalytics.String("x"),
				Continent:     chunkifyanalytics.ExportQueryContinentAf,
				Country:       chunkifyanalytics.String("SE"),
				DeviceType:    chunkifyanalytics.ExportQueryDeviceTypeDesktop,
				Dimension:     chunkifyanalytics.ExportQueryDimensionBrowserName,
				EventName:     chunkifyanalytics.ExportQueryEventNameChunkifyVideoViewStarted,
				From:          chunkifyanalytics.Time(time.Now()),
				GroupBy:       chunkifyanalytics.ExportQueryGroupByBrowserName,
				HasFatalError: chunkifyanalytics.Bool(true),
				Interval:      chunkifyanalytics.ExportQueryIntervalHour,
				Kind:          chunkifyanalytics.ExportQueryKindSummary,
				Order:         chunkifyanalytics.ExportQueryOrderDesc,
				OsName:        chunkifyanalytics.String("x"),
				PageOrigin:    chunkifyanalytics.String("x"),
				PagePath:      chunkifyanalytics.String("x"),
				SortBy:        chunkifyanalytics.ExportQuerySortByViews,
				To:            chunkifyanalytics.Time(time.Now()),
				VideoID:       chunkifyanalytics.String("video_id"),
				ViewID:        chunkifyanalytics.String("x"),
				ViewerID:      chunkifyanalytics.String("viewer_id"),
			},
		},
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportDownload(t *testing.T) {
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
	_, err := client.Exports.Download(context.TODO(), "2e-xAv1b-zE_-N_eC1_0f_sO_-a-i_D--_3f-__c703")
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportListColumnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.ListColumns(context.TODO(), chunkifyanalytics.ExportListColumnsParams{
		Resource: chunkifyanalytics.ExportListColumnsParamsResourceSummary,
		GroupBy:  chunkifyanalytics.ExportListColumnsParamsGroupByBrowserName,
		Kind:     chunkifyanalytics.ExportListColumnsParamsKindSummary,
	})
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
