package mashable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopicList(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	posts, err := c.Topics().List(&ListOptions{
		Page:    1,
		PerPage: 1,
	})
	assert.Nil(t, err)
	assert.Len(t, posts.Topics, 1)
}
