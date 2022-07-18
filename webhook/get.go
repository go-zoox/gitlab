package webhook

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id/hooks/:webhook_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetRequest struct {
	ProjectID int64 `json:"project_id"`
	WebHookID int64 `json:"webhook_id"`
}

type GetResponse = WebHook

func Get(cfg *GetRequest) (*GetResponse, error) {
	response, err := request.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(cfg.ProjectID)),
			"webhook_id": strconv.Itoa(int(cfg.WebHookID)),
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
