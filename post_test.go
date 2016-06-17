package mashable

import (
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
}
