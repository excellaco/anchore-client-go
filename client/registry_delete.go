package client

import (
	"fmt"
	"net/http"
)

func (c *Client) RegistryDelete(registryURL *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/registries/%s", c.HostURL, *registryURL), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	return err
}
