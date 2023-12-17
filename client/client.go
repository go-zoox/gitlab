package client

import "github.com/go-zoox/gitlab/request"

type Client interface {
	Request(cfg *request.Config, payload *request.Payload) (*request.Response, error)
}

type client struct{}

func New() Client {
	return &client{}
}

func (c *client) Request(cfg *request.Config, payload *request.Payload) (*request.Response, error) {
	return request.Request(cfg, payload)
}
