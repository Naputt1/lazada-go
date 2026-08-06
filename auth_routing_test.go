package golazada

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRefreshAccessTokenUsesAuthDomain(t *testing.T) {
	setup()
	defer teardown()

	const authURL = "https://auth.lazada.com/rest/auth/token/refresh"
	httpmock.RegisterResponder("GET", authURL,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{
				"code": "0",
				"type": "0",
				"message": "success",
				"request_id": "test",
				"data": {
					"access_token": "new_access_token",
					"refresh_token": "new_refresh_token",
					"expires_in": 3600
				}
			}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	client.Region = "TH"
	client.Token = "expired_token"
	client.RefreshToken = "refresh_me"

	res, err := client.Auth.RefreshAccessToken(context.Background(), "refresh_me")
	if err != nil {
		t.Fatalf("RefreshAccessToken: %v", err)
	}
	if res.AccessToken != "new_access_token" {
		t.Fatalf("AccessToken = %q, want %q", res.AccessToken, "new_access_token")
	}
	if res.RefreshToken != "new_refresh_token" {
		t.Fatalf("RefreshToken = %q, want %q", res.RefreshToken, "new_refresh_token")
	}
}

func TestGetOrdersAutoRefreshesOnIllegalAccessToken(t *testing.T) {
	setup()
	defer teardown()

	client.Region = "TH"
	client.Token = "expired_token"
	client.RefreshToken = "refresh_me"

	var mu sync.Mutex
	orderCalls := 0

	httpmock.RegisterResponder("GET", "https://api.lazada.co.th/rest/orders/get",
		func(req *http.Request) (*http.Response, error) {
			mu.Lock()
			orderCalls++
			calls := orderCalls
			mu.Unlock()

			body := `{"code": "0", "data": {"count": 1, "countTotal": 1, "orders": [{"order_number": "1001", "order_id": 1001, "statuses": []}]}}`
			if calls == 1 {
				body = `{"code": "IllegalAccessToken", "type": "ISV", "message": "The specified access token is invalid or expired", "request_id": "test"}`
			}
			resp := httpmock.NewStringResponse(200, body)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "https://auth.lazada.com/rest/auth/token/refresh",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{
				"code": "0",
				"type": "0",
				"message": "success",
				"request_id": "test",
				"data": {
					"access_token": "fresh_token",
					"refresh_token": "new_refresh_token",
					"expires_in": 3600
				}
			}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := client.Order.GetOrders(context.Background(), GetOrdersRequest{})
	if err != nil {
		t.Fatalf("GetOrders after auto-refresh: %v", err)
	}
	if len(res.Response.Orders) != 1 {
		t.Fatalf("got %d orders, want 1", len(res.Response.Orders))
	}
	if client.Token != "fresh_token" {
		t.Fatalf("client token = %q, want %q", client.Token, "fresh_token")
	}

	mu.Lock()
	defer mu.Unlock()
	if orderCalls != 2 {
		t.Fatalf("orders/get called %d times, want 2 (original + retry after refresh)", orderCalls)
	}
}
