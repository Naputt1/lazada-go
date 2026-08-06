package golazada

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/jarcoal/httpmock"
)

var (
	client *DefaultClient
	app    App

	skippedMu     sync.Mutex
	skippedRoutes []string
)

func setup() {
	httpmock.Activate()
	app = App{
		AppKey:    "test_app_key",
		AppSecret: "test_app_secret",
	}
	client = NewDefaultClient(app)
	client.Region = "SG"
	client.Token = "test_access_token"
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixtureSafe(path string) (interface{}, error) {
	f, err := os.Open("fixtures/" + path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
