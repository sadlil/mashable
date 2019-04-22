package mashable

import (
	"errors"
	"net/http"
)

type Client struct {
	apiEndpoint string
	version     string
	resource    string
	uri         []string

	httpClient *http.Client
	request    *http.Request
	response   *http.Response
}

func New() (*Client, error) {
	client := &Client{
		apiEndpoint: "https://api.mashable.com",
		version:     "v1",
		uri:         make([]string, 0),

		httpClient: http.DefaultClient,
	}

	resp, err := http.Get(client.apiEndpoint)
	if err != nil {
		return nil, err
	}

	if !isStatusOK(resp.StatusCode) {
		return nil, errors.New(resp.Status)
	}

	return client, nil
}

func (c *Client) Posts() PostInterface {
	return newPost(c)
}

func (c *Client) Search() SearchInterface {
	return newSearch(c)
}

func (c *Client) Topics() TopicInterface {
	return newTopic(c)
}

func (c *Client) Users() UserInterface {
	return newUser(c)
}
