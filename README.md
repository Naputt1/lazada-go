# go-lazada-v1

> Auto-generated, type-safe Go SDK for the Lazada Open Platform API

**32 service modules**, **341 API methods**, **300 error constants**, and **371 test fixtures** — all auto-generated from [Lazada's official API documentation](https://open.lazada.com/apps/doc/api) using [doclient](https://github.com/Naputt1/doclient).

- HMAC-SHA256 request signing (matching Lazada's specification)
- Per-region API gateways (SG, MY, VN, TH, PH, ID)
- Automatic access token refresh on expiry
- File upload support via multipart/form-data
- Go 1.22 generics for extensible metadata

## Install

```bash
go get github.com/naputt1/go-lazada-v1
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/naputt1/go-lazada-v1"
)

func main() {
    app := golazada.App{
        AppKey:    os.Getenv("LAZADA_APP_KEY"),
        AppSecret: os.Getenv("LAZADA_APP_SECRET"),
    }

    client := golazada.NewDefaultClient(app)
    client.Region = "SG"
    client.Token = os.Getenv("LAZADA_ACCESS_TOKEN")

    orders, err := client.Order.GetOrders(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // orders.Data contains the response payload
    fmt.Printf("Orders: %+v\n", orders)
}
```

## Client

The client is generic over a metadata type `T`.

### Default client

```go
client := golazada.NewDefaultClient(app)
// client is *golazada.Client[any]
```

### With custom metadata

```go
type Meta struct {
    ShopID   uint64
    StoreName string
}

client := golazada.NewClient[Meta](app, golazada.WithMeta(Meta{
    ShopID:    123456,
    StoreName: "my-store",
}))
```

The `Meta` value is passed to your `OnTokenRefresh` callback:

```go
golazada.WithOnTokenRefresh(func(res *golazada.RefreshAccessTokenResponse, meta Meta) {
    db.SaveRefreshToken(meta.ShopID, res.RefreshToken)
})
```

### Options

| Option | Description |
|--------|-------------|
| `WithRegion` | API region (SG, MY, VN, TH, PH, ID) |
| `WithHTTPClient` | Custom `*http.Client` (default: 10s timeout) |
| `WithRetry` | Number of retry attempts on failure |
| `WithLogger` | Custom `LeveledLoggerInterface` |
| `WithProxy` | HTTP proxy URL |
| `WithRefreshToken` | Refresh token for auto-renewal |
| `WithOnTokenRefresh` | Callback after successful token refresh |
| `WithMeta` | Custom metadata (generic) |

### Region gateways

| Region | URL |
|--------|-----|
| `SG` | `https://api.lazada.sg/rest` |
| `MY` | `https://api.lazada.com.my/rest` |
| `VN` | `https://api.lazada.vn/rest` |
| `TH` | `https://api.lazada.co.th/rest` |
| `PH` | `https://api.lazada.com.ph/rest` |
| `ID` | `https://api.lazada.co.id/rest` |
| `AUTH` | `https://auth.lazada.com/rest` |

## Usage

### GET request

```go
result, err := client.Order.GetOrders(context.Background())
```

### POST request

API-specific parameters are sent automatically by the generated service methods. The client handles system-level params (`app_key`, `sign_method`, `timestamp`, `partner_id`, `access_token`, `sign`).

### Pointer helpers

The SDK provides a `Ptr` helper for optional fields:

```go
golazada.Ptr("NORMAL")     // *string
golazada.Ptr(int64(100))   // *int64
golazada.Ptr(29.99)        // *float64
```

## Authentication

### OAuth flow

```go
// 1. Redirect user to Lazada's auth URL:
//    https://auth.lazada.com/oauth/authorize?response_type=code&client_id=APP_KEY&redirect_uri=...

// 2. Lazada redirects back with ?code=...

// 3. Exchange code for tokens
client.Region = "AUTH"
token, err := client.Auth.GetAccessToken(ctx, "code_from_redirect")
// token.AccessToken, token.RefreshToken, token.ExpireIn

// 4. Refresh on expiry
newToken, err := client.Auth.RefreshAccessToken(ctx, "refresh_token")
```

### Auto refresh

When `WithRefreshToken` is set, the client automatically detects expired token errors, refreshes it, and retries:

```go
client := golazada.NewDefaultClient(app,
    golazada.WithRefreshToken("initial_refresh_token"),
    golazada.WithOnTokenRefreshDefault(func(res *golazada.RefreshAccessTokenResponse, meta any) {
        // Persist the new tokens
    }),
)
```

## Error Handling

### ResponseError

All API calls return `ResponseError` on failure:

```go
type ResponseError struct {
    Status    int
    Code      string
    Type      string
    Message   string
    RequestID string
}
```

### Checking for API errors

Lazada returns `code: "0"` on success. Non-zero codes indicate errors:

```go
result, err := client.Order.GetOrders(context.Background())
if err != nil {
    if re, ok := err.(golazada.ResponseError); ok {
        fmt.Printf("Error %s: %s (request: %s)\n", re.Code, re.Message, re.RequestID)
    }
}
```

The SDK includes **300 named error constants** covering all documented Lazada error codes:

```go
ErrMissingParameter                      // Missing required parameter
ErrIncompleteSignature                   // Invalid signature
ErrInvalidCode                           // Invalid authorization code
Err10002                                 // Incorrect product attributes
Err10003                                 // Item not found
Err1000012                               // Invalid date range (>180 days)
// ... 294 more
```

## BaseResponse

Every API response embeds `BaseResponse`:

```go
type BaseResponse struct {
    Code      string `json:"code"`
    Type      string `json:"type"`
    Message   string `json:"message"`
    RequestID string `json:"request_id"`
}
```

## Services

| Service | Methods | File |
|---------|---------|------|
| Auth | 2 | `auth.go` (hand-written) |
| ChoiceCustomized | 12 | `choice_customized.gen.go` |
| Content | 7 | `content.gen.go` |
| CrossBoarderProduct | 11 | `cross_boarder_product.gen.go` |
| ETickets | 8 | `e_tickets.gen.go` |
| EarlyBirdPrice | 4 | `early_bird_price.gen.go` |
| FBL | 51 | `fbl.gen.go` |
| Finance | 4 | `finance.gen.go` |
| FirstMileBigbagOnlyForCN | 9 | `first_mile_bigbagonly_for_cn.gen.go` |
| Flexicombo | 9 | `flexicombo.gen.go` |
| FreeShipping | 11 | `free_shipping.gen.go` |
| Fulfillment | 10 | `fulfillment.gen.go` |
| InstantMessaging | 7 | `instant_messaging.gen.go` |
| LazLike | 13 | `laz_like.gen.go` |
| LazLive | 1 | `laz_live.gen.go` |
| LazPay | 24 | `laz_pay.gen.go` |
| LazadaDG | 7 | `lazada_dg.gen.go` |
| LazadaLogistics | 20 | `lazada_logistics.gen.go` |
| LazadaWalletCorporateTopUp | 5 | `lazada_wallet_corporate_top_up.gen.go` |
| Logistics | 9 | `logistics.gen.go` |
| LogisticsStation | 18 | `logistics_station.gen.go` |
| MediaCenter | 6 | `media_center.gen.go` |
| Membership | 10 | `membership.gen.go` |
| Order | 8 | `order.gen.go` |
| ProductReview | 3 | `product_review.gen.go` |
| RedMart | 8 | `red_mart.gen.go` |
| ReturnAndRefund | 8 | `return_and_refund.gen.go` |
| Seller | 17 | `seller.gen.go` |
| SellerVoucher | 9 | `seller_voucher.gen.go` |
| ServiceMarket | 2 | `service_market.gen.go` |
| SponsoredSolutions | 28 | `sponsored_solutions.gen.go` |
| StoreDecoration | 1 | `store_decoration.gen.go` |
| System | 1 | `system.gen.go` |

## Regenerating

The SDK is generated using [doclient](https://github.com/Naputt1/doclient). To regenerate from the latest Lazada API docs:

```bash
pnpm install
pnpm run generate
```

See `doclient.config.ts` for the generation configuration.

## Testing

```bash
go test ./...
```

All tests use [httpmock](https://github.com/jarcoal/httpmock) with real API response fixtures (371 JSON files in `fixtures/`). No network access required.

## License

MIT
