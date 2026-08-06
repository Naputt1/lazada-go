package golazada

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAuthGetAccessToken(t *testing.T) {
	setup()
	defer teardown()

	client.Region = "AUTH"
	serverURL := regionURLs["AUTH"]

	data, err := loadFixtureSafe("auth.token.create_resp.json")
	if err != nil {
		t.Fatalf("load fixture: %v", err)
	}

	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Fatalf("create responder: %v", err)
	}
	httpmock.RegisterResponder("GET", serverURL+"/auth/token/create", responder)

	res, err := client.Auth.GetAccessToken(context.Background(), "test_code")
	if err != nil {
		t.Fatalf("Auth.GetAccessToken error: %s", err)
	}

	if res.AccessToken == "" {
		t.Error("Expected non-empty AccessToken")
	}
}

func TestAuthRefreshAccessToken(t *testing.T) {
	setup()
	defer teardown()

	client.Region = "AUTH"
	serverURL := regionURLs["AUTH"]

	data, err := loadFixtureSafe("auth.token.refresh_resp.json")
	if err != nil {
		t.Fatalf("load fixture: %v", err)
	}

	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Fatalf("create responder: %v", err)
	}
	httpmock.RegisterResponder("GET", serverURL+"/auth/token/refresh", responder)

	res, err := client.Auth.RefreshAccessToken(context.Background(), "test_refresh_token")
	if err != nil {
		t.Fatalf("Auth.RefreshAccessToken error: %s", err)
	}

	if res.AccessToken == "" {
		t.Error("Expected non-empty AccessToken")
	}
}
