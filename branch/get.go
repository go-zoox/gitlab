package branch

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id/repository/branches/:branch_name",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetRequest struct {
	ProjectID int64  `json:"project_id"`
	Name      string `json:"branch_name"`
}

type GetResponse = Branch

func Get(client client.Client, cfg *GetRequest) (*GetResponse, error) {
	response, err := client.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"project_id":  strconv.Itoa(int(cfg.ProjectID)),
			"branch_name": cfg.Name,
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
