package user

import (
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var ListALLConfig = &request.Config{
	Action:   "GET",
	Resource: "users",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type ListALLRequest struct {
	Active    bool   `json:"active"`
	Blocked   bool   `json:"blocked"`
	OrderBy   string `json:"order_by"`
	Sort      string `json:"sort"`
	TwoFactor bool   `json:"two_factor"`
}

type ListALLResponse struct {
	Data  []User `json:"data"`
	Total int    `json:"total"`
}

func ListALL(client client.Client, opts ...func(opt *ListALLRequest)) (*ListALLResponse, error) {
	opt := &ListALLRequest{}
	for _, o := range opts {
		o(opt)
	}

	users := []User{}
	page := 0
	for {
		page += 1

		response, err := List(client, func(opt *ListRequest) {
			opt.Page = page
			opt.PerPage = 100
		})
		if err != nil {
			return nil, err
		}

		users = append(users, response.Data...)
		if response.IsLastPage {
			break
		}
	}

	return &ListALLResponse{
		Data:  users,
		Total: len(users),
	}, nil
}
