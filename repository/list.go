package repository

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var ListConfig = &request.Config{
	Action:   "GET",
	Resource: "groups/:group_id/projects",
}

type ListRequest struct {
	GroupID int64 `json:"group_id"`
}

type ListResponse = []Repository

func List(cfg *ListRequest) (*ListResponse, error) {
	response, err := request.Request(ListConfig, &request.Payload{
		Params: map[string]string{
			"group_id": strconv.Itoa(int(cfg.GroupID)),
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
