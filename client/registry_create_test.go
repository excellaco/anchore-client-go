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

func TestRegistryCreate(t *testing.T) {
	expectedURL := "/registries"
	testRegistry := &types.Registry{
		URL:  "example.com:5000",
		Name: "test",
	}

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPost {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
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

	registries, err := client.RegistryCreate(*testRegistry)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, registries[0].URL, testRegistry.URL, "Registry URL should match")
	assert.Equal(t, registries[0].Name, testRegistry.Name, "Registry Name should match")
}

func TestRegistryCreateError(t *testing.T) {
	expectedURL := "/registries"
	testRegistry := &types.Registry{
		URL:  "example.com:5000",
		Name: "test",
	}

	client := &Client{
		HTTPClient: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPost {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
			}, nil
		}),
	}

	_, err := client.RegistryCreate(*testRegistry)

	clientError := err.(*ClientError)

	assert.Equal(t, clientError.StatusCode, http.StatusInternalServerError, "Status code should match")
	assert.Equal(t, clientError.URL, expectedURL, "Registry URL should match")
}
