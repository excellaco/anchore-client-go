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

func TestRegistryUpdate(t *testing.T) {
	expectedRegistryURL := "example.com:5000"
	expectedURL := fmt.Sprintf("/registries/%s", expectedRegistryURL)

	testRegistry := &types.Registry{
		URL:  expectedRegistryURL,
		Name: "test",
	}

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPut {
				return nil, fmt.Errorf("expected PUT method, got %s", req.Method)
			}

			content, err := json.Marshal([]*types.Registry{testRegistry})
			if err != nil {
				return nil, err
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(content))),
			}, nil
		}),
	}

	registries, err := client.RegistryUpdate(*testRegistry)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, registries[0].URL, testRegistry.URL, "Registry URL should match")
	assert.Equal(t, registries[0].Name, testRegistry.Name, "Registry Name should match")
}
