package client

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

type AuthStruct struct {
	Username string
	Password string
}

func NewClient(host, username, password *string) *Client {
	if host == nil {
		*host = os.Getenv("ANCHORE_CLI_URL")
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
		Auth: AuthStruct{
			Username: *username,
			Password: *password,
		},
	}

	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	authToken := fmt.Sprintf("%s:%s", c.Auth.Username, c.Auth.Password)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(authToken))))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}