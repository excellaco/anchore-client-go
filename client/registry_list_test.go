package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/excellaco/anchore-client-go/types"
	"gotest.tools/assert"
)

func TestRegistryList(t *testing.T) {
	expectedURL := "/registries"

	testRegistry := &types.Registry{
		URL:  "example:5000",
		Name: "test",
	}

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodGet {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}
			content, err := json.Marshal([]types.Registry{*testRegistry})
			if err != nil {
				return nil, err
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(content)),
			}, nil
		}),
	}

	registries, err := client.RegistryList()
	if err != nil {
		t.Fatal(err)
	}

	if len(registries) != 1 {
		t.Fatalf("expected 1 checkpoint, got %v", registries)
	}

	assert.Equal(t, registries[0].URL, testRegistry.URL, "Registry URL should match")
	assert.Equal(t, registries[0].Name, testRegistry.Name, "Registry Name should match")
}
