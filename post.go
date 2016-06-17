package mashable

import (
	"errors"
)

type postInterface interface {
	List(*ListOptions) (*PostList, error)
	Get(string) (*Post, error)
	// GetFromUrl(string) (*Post, error)
}

type post struct {
	c *Client
}

func newPost(c *Client) postInterface {
	return &post{c}
}

func (p *post) List(o *ListOptions) (*PostList, error) {
	posts := &PostList{}
	var err error
	p.c.resource = "posts"
	p.c.request = p.c.newRequest("GET")
	p.c.request.URL.RawQuery = o.Encode()
	p.c.response, err = p.c.httpClient.Do(p.c.request)
	if err != nil {
		return posts, err
	}

	if !isStatusOK(p.c.response.StatusCode) {
		return posts, errors.New(p.c.response.Status)
	}

	if err := p.c.unMarshalInto(posts); err != nil {
		return posts, err
	}

	if posts.Error != nil {
		return posts, errors.New(posts.Error.Message)
	}

	return posts, nil
}

func (p *post) Get(id string) (*Post, error) {
	post := &Post{}
	var err error
	p.c.resource = "posts"
	p.c.uri = append(p.c.uri, id)
	p.c.request = p.c.newRequest("GET")
	p.c.response, err = p.c.httpClient.Do(p.c.request)
	if err != nil {
		return post, err
	}

	if !isStatusOK(p.c.response.StatusCode) {
		return post, errors.New(p.c.response.Status)
	}

	if err := p.c.unMarshalInto(post); err != nil {
		return post, err
	}

	if post.Error != nil {
		return post, errors.New(post.Error.Message)
	}

	return post, nil
}
