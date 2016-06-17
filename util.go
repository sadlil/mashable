package mashable

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (c *Client) url() string {
	url := c.apiEndpoint + "/" + c.version
	if c.resource != "" {
		url = url + "/" + c.resource
	}
	for _, uri := range c.uri {
		url = url + "/" + uri
	}
	return url
}

func (l *ListOptions) Encode() string {
	query, _ := query.Values(l)
	return query.Encode()
}

func (c *Client) newRequest(method string) *http.Request {
	req, _ := http.NewRequest(method, c.url(), nil)
	return req
}

func (c *Client) unMarshalInto(i interface{}) error {
	if c.response == nil {
		return errors.New("received nil response")
	}

	if c.response.Body == nil {
		return errors.New("response body empty")
	}

	err := json.NewDecoder(c.response.Body).Decode(i)
	if err != nil {
		return err
	}
	c.response.Body.Close()
	return nil
}

func isStatusOK(c int) bool {
	if c >= 200 && c <= 299 {
		return true
	}
	return false
}
