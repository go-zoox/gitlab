package branch

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
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

	// the name of the project
	RepositoryName string `json:"repository_name"`
}

func Delete(client client.Client, req *DeleteRequest) error {
	if req.RepositoryName != "" && req.RepositoryID == 0 {
		repository, err := repository.Get(client, req.RepositoryName)
		if err != nil {
			return fmt.Errorf("failed to get repository by name: %v", err)
		}

		req.RepositoryID = repository.ID
	}

	_, err := client.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id":  strconv.Itoa(int(req.RepositoryID)),
			"branch_name": req.Name,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
