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
)

func TestRegistryList(t *testing.T) {
	expectedURL := "/registries"

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			content, err := json.Marshal([]types.Registry{
				{
					Name: "test",
					URL:  "example.com:5000",
				},
			})
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
}
