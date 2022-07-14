package branch

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var DeleteConfig = &request.Config{
	Action:   "DELETE",
	Resource: "projects/:project_id/repository/branches/:branch_name",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type DeleteRequest struct {
	ProjectID int64  `json:"project_id"`
	Name      string `json:"branch_name"`
}

func Delete(cfg *DeleteRequest) error {
	_, err := request.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id":  strconv.Itoa(int(cfg.ProjectID)),
			"branch_name": cfg.Name,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
