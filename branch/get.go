package branch

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
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
	// the id of the project
	RepositoryID int64  `json:"repository_id"`
	Name         string `json:"branch_name"`

	// the name of the project
	RepositoryName string `json:"repository_name"`
}

type GetResponse = Branch

func Get(client client.Client, req *GetRequest) (*GetResponse, error) {
	if req.RepositoryName != "" && req.RepositoryID == 0 {
		repository, err := repository.Get(client, req.RepositoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to get repository by name: %v", err)
		}

		req.RepositoryID = repository.ID
	}

	response, err := client.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"project_id":  strconv.Itoa(int(req.RepositoryID)),
			"branch_name": req.Name,
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
