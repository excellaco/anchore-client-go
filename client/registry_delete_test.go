package client

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestRegistryDelete(t *testing.T) {
	expectedRegistryURL := "example.com:5000"
	expectedURL := fmt.Sprintf("/registries/%s", expectedRegistryURL)

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodDelete {
				return nil, fmt.Errorf("expected DELETE method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		}),
	}

	err := client.RegistryDelete(&expectedRegistryURL)
	if err != nil {
		t.Fatal(err)
	}
}
