// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics_test

import (
	"context"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-analytics-go"
	"github.com/chunkifydev/chunkify-analytics-go/internal/testutil"
	"github.com/chunkifydev/chunkify-analytics-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Events.ListAutoPaging(context.TODO(), chunkifyanalytics.EventListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		event := iter.Current()
		t.Logf("%+v\n", event.IngestKeyID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
