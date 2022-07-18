package client

import "github.com/go-zoox/gitlab/request"

type Client interface {
	Request(cfg *request.Config, payload *request.Payload) (*request.Response, error)
}
