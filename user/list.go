package user

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var ListConfig = &request.Config{
	Action:   "GET",
	Resource: "users",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type ListRequest struct {
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
	Username  string `json:"username"`
	Active    bool   `json:"active"`
	Blocked   bool   `json:"blocked"`
	OrderBy   string `json:"order_by"`
	Sort      string `json:"sort"`
	TwoFactor bool   `json:"two_factor"`
}

type ListResponse struct {
	Data       []User `json:"data"`
	IsLastPage bool   `json:"is_last_page"`
}

func List(client client.Client, opts ...func(opt *ListRequest)) (*ListResponse, error) {
	opt := &ListRequest{}
	for _, o := range opts {
		o(opt)
	}

	if opt.Page == 0 {
		opt.Page = 1
	}
	if opt.PerPage == 0 {
		opt.PerPage = 20
	}
	if opt.PerPage > 100 {
		opt.PerPage = 100
	}

	query := map[string]string{
		"page":     strconv.Itoa(int(opt.Page)),
		"per_page": strconv.Itoa(int(opt.PerPage)),
		"username": opt.Username,
	}
	if opt.Username != "" {
		query["username"] = opt.Username
	}
	if opt.Active {
		query["active"] = "true"
	}
	if opt.Blocked {
		query["blocked"] = "true"
	}
	if opt.OrderBy != "" {
		query["order_by"] = opt.OrderBy
	}
	if opt.Sort != "" {
		query["sort"] = opt.Sort
	}
	if opt.TwoFactor {
		query["two_factor"] = "true"
	}

	response, err := client.Request(ListConfig, &request.Payload{
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	users := []User{}
	if err = response.UnmarshalJSON(&users); err != nil {
		return nil, err
	}

	return &ListResponse{
		Data:       users,
		IsLastPage: len(users) < opt.PerPage,
	}, nil
}
