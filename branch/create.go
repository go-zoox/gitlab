package branch

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "projects/:project_id/repository/branches",
}

type CreateRequest struct {
	// the id of the project
	RepositoryID int64  `json:"repostory_id"`
	Name         string `json:"name"`

	// the name of the project
	RepositoryName string `json:"repository_name"`
	Ref            string `json:"ref"`
}

type CreateResponse = Branch

func Create(client client.Client, req *CreateRequest) (*CreateResponse, error) {
	if req.RepositoryName != "" && req.RepositoryID == 0 {
		repository, err := repository.Get(client, req.RepositoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to get repository by name: %v", err)
		}

		req.RepositoryID = repository.ID
	}

	response, err := client.Request(CreateConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(req.RepositoryID)),
		},
		Query: map[string]string{
			"branch": req.Name,
			"ref":    req.Ref,
		},
	})

	if err != nil {
		return nil, err
	}

	var res CreateResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
