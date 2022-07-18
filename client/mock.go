package client

import "github.com/go-zoox/gitlab/request"

type MockClient struct {
}

func NewMockClient() Client {
	return &MockClient{}
}

func (c *MockClient) Request(cfg *request.Config, payload *request.Payload) (*request.Response, error) {
	return request.Request(cfg, payload)
}
