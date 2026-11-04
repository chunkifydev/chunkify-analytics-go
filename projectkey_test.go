// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkifyanalytics_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-analytics-go"
	"github.com/chunkifydev/chunkify-analytics-go/internal/testutil"
	"github.com/chunkifydev/chunkify-analytics-go/option"
)

func TestProjectKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.Keys.New(context.TODO(), chunkifyanalytics.ProjectKeyNewParams{
		CreateProjectKeyRequest: chunkifyanalytics.CreateProjectKeyRequestParam{
			Name:       "Dashboard",
			Permission: chunkifyanalytics.CreateProjectKeyRequestPermissionSecret,
			ReadOnly:   chunkifyanalytics.Bool(true),
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

func TestProjectKeyList(t *testing.T) {
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
	_, err := client.Project.Keys.List(context.TODO())
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectKeyRotate(t *testing.T) {
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
	_, err := client.Project.Keys.Rotate(context.TODO(), "key_id")
	if err != nil {
		var apierr *chunkifyanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
