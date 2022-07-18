package repository

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetResponse = Repository

func Get(client client.Client, id int64) (*GetResponse, error) {
	response, err := client.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(id)),
		},
	})

	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
