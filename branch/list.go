package branch

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
	"github.com/go-zoox/gitlab/request"
)

var ListConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id/repository/branches",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type ListRequest struct {
	// the id of the project
	RepositoryID int64 `json:"repository_id"`
	// the name of the project
	RepositoryName string `json:"repository_name"`
}

type ListResponse = []Branch

func List(client client.Client, req *ListRequest) (*ListResponse, error) {
	if req.RepositoryName != "" && req.RepositoryID == 0 {
		repository, err := repository.Get(client, req.RepositoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to get repository by name: %v", err)
		}

		req.RepositoryID = repository.ID
	}

	response, err := client.Request(ListConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(req.RepositoryID)),
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
