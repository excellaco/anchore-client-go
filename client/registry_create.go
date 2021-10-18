package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/excellaco/anchore-client-go/types"
)

func (c *Client) RegistryCreate(registry types.Registry) (*types.Registry, error) {
	rb, err := json.Marshal(registry)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/registries", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newRegistries := []types.Registry{}
	err = json.Unmarshal(body, &newRegistries)
	if err != nil {
		return nil, err
	}

	return &newRegistries[0], nil
}
