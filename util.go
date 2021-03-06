package mashable

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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

func (c *Client) newRequest(method string) *http.Request {
	req, _ := http.NewRequest(method, c.url(), nil)
	return req
}

func (c *Client) unmarshalInto(i interface{}) error {
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
	return c.response.Body.Close()
}

func (l *ListOptions) Encode() string {
	query, _ := query.Values(l)
	return query.Encode()
}

func (t *Topic) parseName(name string) string {
	name = strings.Replace(strings.ToLower(name), " ", "-", -1)
	return name
}

func isStatusOK(c int) bool {
	if c >= 200 && c <= 299 {
		return true
	}
	return false
}
