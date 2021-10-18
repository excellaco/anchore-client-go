package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/excellaco/anchore-client-go/types"
)

func (c *Client) RegistryUpdate(registry types.Registry) ([]*types.Registry, error) {
	rb, err := json.Marshal(registry)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/registries/%s", c.HostURL, registry.URL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedRegistries := []*types.Registry{}
	err = json.Unmarshal(body, &updatedRegistries)
	if err != nil {
		return nil, err
	}

	return updatedRegistries, nil
}
