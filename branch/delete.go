package branch

import (
	"strconv"

	"github.com/go-zoox/gitlab/client"
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
	// the id of the project
	RepositoryID int64  `json:"repository_id"`
	Name         string `json:"branch_name"`
}

func Delete(client client.Client, cfg *DeleteRequest) error {
	_, err := client.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id":  strconv.Itoa(int(cfg.RepositoryID)),
			"branch_name": cfg.Name,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
