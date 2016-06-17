package mashable

import (
	"errors"
)

type userInterface interface {
	Get(string) (*User, error)
}

type user struct {
	c *Client
}

func newUser(c *Client) userInterface {
	return &user{c}
}

func (u *user) Get(id string) (*User, error) {
	user := &User{}
	var err error
	u.c.resource = "users"
	u.c.uri = append(u.c.uri, id)
	u.c.request = u.c.newRequest("GET")
	u.c.response, err = u.c.httpClient.Do(u.c.request)
	if err != nil {
		return user, err
	}

	if !isStatusOK(u.c.response.StatusCode) {
		return user, errors.New(u.c.response.Status)
	}

	if err := u.c.unMarshalInto(user); err != nil {
		return user, err
	}

	if user.Error != nil {
		return user, errors.New(user.Error.Message)
	}

	return user, nil
}
