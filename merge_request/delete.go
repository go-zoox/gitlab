package merge_request

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var DeleteConfig = &request.Config{
	Action:   "DELETE",
	Resource: "projects/:project_id/merge_requests/:merge_request_id",
}

type DeleteRequest struct {
	ProjectID      int64 `json:"project_id"`
	MergeRequestID int64 `json:"merge_request_id"`
}

func Delete(cfg *DeleteRequest) error {
	_, err := request.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id":       strconv.Itoa(int(cfg.ProjectID)),
			"merge_request_id": strconv.Itoa(int(cfg.MergeRequestID)),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
