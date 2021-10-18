package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/excellaco/anchore-client-go/types"
)

func (c *Client) RegistryRead(registryURL *string) (*types.Registry, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/registries/%s", c.HostURL, *registryURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	registries := []*types.Registry{}
	err = json.Unmarshal(body, &registries)
	if err != nil {
		return nil, err
	}

	if len(registries) == 0 {
		return nil, fmt.Errorf("%s not found", *registryURL)
	}

	return registries[0], nil
}
