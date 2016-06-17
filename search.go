package mashable

import (
	"errors"
	"net/url"
)

type searchInterface interface {
	Query(string) (*SearchResult, error)
}

type search struct {
	c *Client
}

func newSearch(c *Client) searchInterface {
	return &search{c}
}

func (s *search) Query(q string) (*SearchResult, error) {
	result := &SearchResult{}
	var err error
	s.c.resource = "search"
	s.c.request = s.c.newRequest("GET")
	s.c.request.URL.RawQuery = "q=" + url.QueryEscape(q)
	s.c.response, err = s.c.httpClient.Do(s.c.request)
	if err != nil {
		return result, err
	}

	if !isStatusOK(s.c.response.StatusCode) {
		return result, errors.New(s.c.response.Status)
	}

	if err := s.c.unMarshalInto(result); err != nil {
		return result, err
	}

	return result, nil
}
