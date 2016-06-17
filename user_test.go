package mashable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	user, err := c.Users().Get("5396405b97b2f841ce000c99")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "Sandra Gonzalez")
}
