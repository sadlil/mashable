package mashable

import (
	"errors"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

type PostInterface interface {
	List(*ListOptions) (*PostList, error)
	Get(string) (*Post, error)
	GetFromUrl(string) (*Post, error)
}

type post struct {
	c *Client
}

func newPost(c *Client) PostInterface {
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

	if err := p.c.unmarshalInto(posts); err != nil {
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

	if err := p.c.unmarshalInto(post); err != nil {
		return post, err
	}

	if post.Error != nil {
		return post, errors.New(post.Error.Message)
	}

	return post, nil
}

func (p *post) GetFromUrl(url string) (*Post, error) {
	post := &Post{}
	resp, err := http.Get(url)
	if err != nil {
		return post, err
	}
	if isStatusOK(resp.StatusCode) {
		root, err := html.Parse(resp.Body)
		if err != nil {
			return post, err
		}

		content, ok := scrape.Find(root, scrape.ByClass("post-slider"))
		if !ok {
			return post, errors.New("failed to find content")
		}

		article, ok := scrape.Find(content, scrape.ByClass("post"))
		if !ok {
			return post, errors.New("failed to find article")
		}

		id := scrape.Attr(article, "data-id")
		return p.Get(id)
	}
	return post, errors.New(resp.Status)
}
