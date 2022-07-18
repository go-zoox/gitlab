package webhook

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var ListConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id/hooks",
}

type ListRequest struct {
	ProjectID int64 `json:"project_id"`
}

type ListResponse = []WebHook

func List(cfg *ListRequest) (*ListResponse, error) {
	response, err := request.Request(ListConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(cfg.ProjectID)),
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
