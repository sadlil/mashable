package mashable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	res, err := c.Search().Query("hello world")
	assert.Nil(t, err)
	assert.NotEqual(t, res.Search.Results.Posts, 0)
}
