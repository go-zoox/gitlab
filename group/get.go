package group

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "groups/:group_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetRequest struct {
	GroupID int64 `json:"group_id"`
}

type GetResponse = Group

func Get(client client.Client, cfg *GetRequest) (*GetResponse, error) {
	response, err := client.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"group_id": strconv.Itoa(int(cfg.GroupID)),
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
