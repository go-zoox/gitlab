package user

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "users/:id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetRequest struct {
	ID uint `json:"id"`
}

type GetResponse = User

func Get(client client.Client, opts ...func(opt *GetRequest)) (*GetResponse, error) {
	opt := &GetRequest{}
	for _, o := range opts {
		o(opt)
	}

	response, err := client.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"id": strconv.Itoa(int(opt.ID)),
		},
	})
	if err != nil {
		return nil, err
	}

	var resp GetResponse
	if err = response.UnmarshalJSON(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
