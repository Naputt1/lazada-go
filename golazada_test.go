package golazada

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if len(skippedRoutes) > 0 {
		fmt.Println("\nSkipped routes (missing fixtures):")
		for _, r := range skippedRoutes {
			fmt.Println("  -", r)
		}
	}
	os.Exit(code)
}

func TestCheckResponseError(t *testing.T) {
	setup()
	defer teardown()

	rawErr := `{"code":"IllegalArgument","type":"ISV","message":"Invalid parameter","request_id":"req123"}`
	serverURL := regionURLs["SG"]
	httpmock.RegisterResponder("GET", serverURL+"/test",
		httpmock.NewStringResponder(400, rawErr))

	_, err := client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	re, ok := err.(ResponseError)
	if !ok {
		t.Fatalf("Expected ResponseError, got %T", err)
	}

	if re.Code != "IllegalArgument" {
		t.Errorf("Expected Code IllegalArgument, got %s", re.Code)
	}

	if re.RequestID != "req123" {
		t.Errorf("Expected RequestID req123, got %s", re.RequestID)
	}
}

func TestGenericMeta(t *testing.T) {
	type MyMeta struct {
		ID   int
		Name string
	}

	meta := MyMeta{ID: 123, Name: "Test"}
	app := App{
		AppKey:    "test_app_key",
		AppSecret: "test_app_secret",
	}

	var capturedMeta MyMeta
	c := NewClient(app,
		WithMeta(meta),
		WithOnTokenRefresh(func(res *RefreshAccessTokenResponse, m MyMeta) {
			capturedMeta = m
		}),
	)

	if c.Meta.ID != 123 {
		t.Errorf("Expected meta ID 123, got %d", c.Meta.ID)
	}

	c.OnTokenRefresh(nil, c.Meta)
	if capturedMeta.ID != 123 {
		t.Errorf("Expected captured meta ID 123, got %d", capturedMeta.ID)
	}
}

func TestNewDefaultClient(t *testing.T) {
	app := App{
		AppKey:    "test_app_key",
		AppSecret: "test_app_secret",
	}

	c := NewDefaultClient(app,
		WithRetryDefault(3),
		WithMetaDefault("some meta"),
	)

	if c.retries != 3 {
		t.Errorf("Expected retries 3, got %d", c.retries)
	}

	if c.Meta != "some meta" {
		t.Errorf("Expected meta 'some meta', got %v", c.Meta)
	}

	client = c
}

func TestRegionURLs(t *testing.T) {
	if regionURLs["SG"] != "https://api.lazada.sg/rest" {
		t.Errorf("Unexpected SG URL: %s", regionURLs["SG"])
	}
	if regionURLs["AUTH"] != "https://auth.lazada.com/rest" {
		t.Errorf("Unexpected AUTH URL: %s", regionURLs["AUTH"])
	}
}
