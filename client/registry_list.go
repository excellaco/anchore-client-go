package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/excellaco/anchore-client-go/types"
)

func (c *Client) RegistryList() ([]types.Registry, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/registries", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	registries := []types.Registry{}
	err = json.Unmarshal(body, &registries)
	if err != nil {
		return nil, err
	}

	return registries, nil
}
