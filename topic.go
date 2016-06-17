package mashable

import (
	"errors"
)

type topicInterface interface {
	List(*ListOptions) (*TopicList, error)
	Get(string) (*Topic, error)
}

type topic struct {
	c *Client
}

func newTopic(c *Client) topicInterface {
	return &topic{c}
}

func (t *topic) List(o *ListOptions) (*TopicList, error) {
	topics := &TopicList{}
	var err error
	t.c.resource = "topics"
	t.c.request = t.c.newRequest("GET")
	t.c.request.URL.RawQuery = o.Encode()
	t.c.response, err = t.c.httpClient.Do(t.c.request)
	if err != nil {
		return topics, err
	}

	if !isStatusOK(t.c.response.StatusCode) {
		return topics, errors.New(t.c.response.Status)
	}

	if err := t.c.unMarshalInto(topics); err != nil {
		return topics, err
	}

	if topics.Error != nil {
		return topics, errors.New(topics.Error.Message)
	}

	return topics, nil
}

func (t *topic) Get(slug string) (*Topic, error) {
	topic := &Topic{}
	var err error
	t.c.resource = "topics"
	t.c.uri = append(t.c.uri, topic.parseName(slug))
	t.c.request = t.c.newRequest("GET")
	t.c.response, err = t.c.httpClient.Do(t.c.request)
	if err != nil {
		return topic, err
	}

	if !isStatusOK(t.c.response.StatusCode) {
		return topic, errors.New(t.c.response.Status)
	}

	if err := t.c.unMarshalInto(topic); err != nil {
		return topic, err
	}

	if topic.Error != nil {
		return topic, errors.New(topic.Error.Message)
	}

	return topic, nil
}
