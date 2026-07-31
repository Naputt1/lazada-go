package golazada

import (
	"encoding/json"
	"testing"
)

func TestFlexStringUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"string", `"0.00"`, "0.00"},
		{"string empty", `""`, ""},
		{"integer number", `0`, "0"},
		{"float number", `0.5`, "0.5"},
		{"scientific number", `1e3`, "1e3"},
		{"null", `null`, ""},
		{"true boolean", `true`, "true"},
		{"false boolean", `false`, "false"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var f FlexString
			if err := json.Unmarshal([]byte(tc.in), &f); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.in, err)
			}
			if string(f) != tc.want {
				t.Fatalf("got %q, want %q", string(f), tc.want)
			}
		})
	}
}

func TestResponseDataOrdersNumericVoucherPlatform(t *testing.T) {
	raw := `{
		"count": "2",
		"countTotal": "2",
		"orders": [{
			"voucher_platform": 0,
			"voucher": 0,
			"voucher_seller": 0.5,
			"shipping_fee": 5.25,
			"order_number": 491253082180001,
			"gift_option": false,
			"is_cancel_pending": true,
			"items_count": 2,
			"price": "106.00",
			"statuses": [],
			"order_id": "491253082180001"
		}]
	}`
	var data GetOrdersResponseData
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		t.Fatalf("unmarshal GetOrdersResponseData: %v", err)
	}
	if len(data.Orders) != 1 {
		t.Fatalf("got %d orders, want 1", len(data.Orders))
	}
	if int(data.Count) != 2 {
		t.Fatalf("Count = %d, want 2", data.Count)
	}
	if int(data.CountTotal) != 2 {
		t.Fatalf("CountTotal = %d, want 2", data.CountTotal)
	}
	if got := string(data.Orders[0].VoucherPlatform); got != "0" {
		t.Fatalf("VoucherPlatform = %q, want %q", got, "0")
	}
	if got := string(data.Orders[0].Voucher); got != "0" {
		t.Fatalf("Voucher = %q, want %q", got, "0")
	}
	if got := string(data.Orders[0].VoucherSeller); got != "0.5" {
		t.Fatalf("VoucherSeller = %q, want %q", got, "0.5")
	}
	if got := string(data.Orders[0].OrderNumber); got != "491253082180001" {
		t.Fatalf("OrderNumber = %q, want %q", got, "491253082180001")
	}
	if got := string(data.Orders[0].GiftOption); got != "false" {
		t.Fatalf("GiftOption = %q, want %q", got, "false")
	}
	if got := string(data.Orders[0].IsCancelPending); got != "true" {
		t.Fatalf("IsCancelPending = %q, want %q", got, "true")
	}
	if got := string(data.Orders[0].ItemsCount); got != "2" {
		t.Fatalf("ItemsCount = %q, want %q", got, "2")
	}
	if got := string(data.Orders[0].ShippingFee); got != "5.25" {
		t.Fatalf("ShippingFee = %q, want %q", got, "5.25")
	}
	if got := float64(data.Orders[0].Price); got != 106.0 {
		t.Fatalf("Price = %v, want 106.0", got)
	}
	if got := int64(data.Orders[0].OrderId); got != 491253082180001 {
		t.Fatalf("OrderId = %d, want 491253082180001", got)
	}
}

func TestFlexFloatUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want float64
	}{
		{"number", `106`, 106},
		{"float number", `106.5`, 106.5},
		{"numeric string", `"106.00"`, 106},
		{"empty string", `""`, 0},
		{"null", `null`, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var f FlexFloat
			if err := json.Unmarshal([]byte(tc.in), &f); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.in, err)
			}
			if float64(f) != tc.want {
				t.Fatalf("got %v, want %v", float64(f), tc.want)
			}
		})
	}
}

func TestFlexIntUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int64
	}{
		{"number", `491253082180001`, 491253082180001},
		{"numeric string", `"491253082180001"`, 491253082180001},
		{"string int", `"2"`, 2},
		{"empty string", `""`, 0},
		{"null", `null`, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var i FlexInt
			if err := json.Unmarshal([]byte(tc.in), &i); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.in, err)
			}
			if int64(i) != tc.want {
				t.Fatalf("got %v, want %v", int64(i), tc.want)
			}
		})
	}
}
