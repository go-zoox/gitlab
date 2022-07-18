package webhook

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var DeleteConfig = &request.Config{
	Action:   "DELETE",
	Resource: "projects/:project_id/hooks/:webhook_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type DeleteRequest struct {
	ProjectID int64 `json:"project_id"`
	WebHookID int64 `json:"webhook_id"`
}

func Delete(cfg *DeleteRequest) error {
	_, err := request.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(cfg.ProjectID)),
			"webhook_id": strconv.Itoa(int(cfg.WebHookID)),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
