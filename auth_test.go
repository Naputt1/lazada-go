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

	httpmock.RegisterResponder("GET", serverURL+"/auth/token/create",
		httpmock.NewBytesResponder(200, loadFixture("auth.token.create_resp.json")))

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

	httpmock.RegisterResponder("GET", serverURL+"/auth/token/refresh",
		httpmock.NewBytesResponder(200, loadFixture("auth.token.refresh_resp.json")))

	res, err := client.Auth.RefreshAccessToken(context.Background(), "test_refresh_token")
	if err != nil {
		t.Fatalf("Auth.RefreshAccessToken error: %s", err)
	}

	if res.AccessToken == "" {
		t.Error("Expected non-empty AccessToken")
	}
}
