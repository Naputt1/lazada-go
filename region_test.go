package golazada

import "testing"

func TestGetServerURLCaseInsensitive(t *testing.T) {
	cases := []struct {
		region string
		want   string
	}{
		{"TH", "https://api.lazada.co.th/rest"},
		{"th", "https://api.lazada.co.th/rest"},
		{"tH", "https://api.lazada.co.th/rest"},
		{"SG", "https://api.lazada.sg/rest"},
		{"sg", "https://api.lazada.sg/rest"},
		{"AUTH", "https://auth.lazada.com/rest"},
		{"auth", "https://auth.lazada.com/rest"},
		{"MY", "https://api.lazada.com.my/rest"},
		{"vn", "https://api.lazada.vn/rest"},
		{"ZZ", "https://api.lazada.sg/rest"},
	}
	for _, tc := range cases {
		c := &Client[any]{Region: tc.region}
		if got := c.getServerURL(); got != tc.want {
			t.Errorf("getServerURL(%q) = %q, want %q", tc.region, got, tc.want)
		}
	}
}
