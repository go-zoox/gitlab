package group

import (
	"github.com/go-zoox/gitlab/request"
)

var ListConfig = &request.Config{
	Action:   "GET",
	Resource: "groups",
}

type ListRequest struct {
	Search string `json:"search"`
}

type ListResponse = []Group

func List(cfg *ListRequest) (*ListResponse, error) {
	response, err := request.Request(ListConfig, &request.Payload{
		Query: map[string]string{
			"search": cfg.Search,
		},
	})

	if err != nil {
		return nil, err
	}

	var res ListResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
