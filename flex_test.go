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
		"count": 1,
		"countTotal": 1,
		"orders": [{
			"voucher_platform": 0,
			"voucher": "0.00",
			"order_number": "491253082180001",
			"statuses": [],
			"order_id": 491253082180001
		}]
	}`
	var data GetOrdersResponseData
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		t.Fatalf("unmarshal GetOrdersResponseData: %v", err)
	}
	if len(data.Orders) != 1 {
		t.Fatalf("got %d orders, want 1", len(data.Orders))
	}
	if got := string(data.Orders[0].VoucherPlatform); got != "0" {
		t.Fatalf("VoucherPlatform = %q, want %q", got, "0")
	}
}
