package mashable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostList(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	posts, err := c.Posts().List(&ListOptions{
		Page:    1,
		PerPage: 1,
	})
	assert.Nil(t, err)
	assert.Len(t, posts.Posts, 1)
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(posts.Posts[0].Title))

	posts, err = c.Posts().List(&ListOptions{
		Page:    5,
		PerPage: 5,
	})
	assert.Nil(t, err)
	assert.Len(t, posts.Posts, 5)
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(posts.Posts[0].Title))
	assert.Equal(t, 5, posts.Collection.Page)
	assert.Equal(t, 5, posts.Collection.PerPage)
}

func TestPostGet(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	posts, err := c.Posts().Get("576377a06df6124759000024")
	assert.Nil(t, err)
	assert.NotNil(t, posts)
	assert.Equal(t, "Litchfield's newest inmate cooks up delicious drama in 'Orange Is the New Black' Season 4", posts.Title)
}

func TestPostGetFromUrl(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	posts, err := c.Posts().GetFromUrl("http://mashable.com/2016/06/17/orange-is-the-new-black-blair-brown")
	assert.Nil(t, err)
	assert.NotNil(t, posts)
	assert.Equal(t, "Litchfield's newest inmate cooks up delicious drama in 'Orange Is the New Black' Season 4", posts.Title)
}
