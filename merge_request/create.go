package merge_request

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "projects/:project_id/merge_requests",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type CreateRequest struct {
	// the id of the project
	RepositoryID int64 `json:"repository_id"`
	// the name of the project
	RepositoryName string `json:"repository_name"`
	SourceBranch   string `json:"source_branch"`
	TargetBranch   string `json:"target_branch"`
	Title          string `json:"title"`
	//
	Description string `json:"description"`
	AssigneeID  int64  `json:"assignee_id"`
	Labels      string `json:"labels"`
	//
	RemoveSourceBranch bool `json:"remove_source_branch"`
	Squash             bool `json:"squash"`
	//
}

type CreateResponse = MergeRequest

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
		Body: map[string]any{
			"source_branch":        req.SourceBranch,
			"target_branch":        req.TargetBranch,
			"title":                req.Title,
			"description":          req.Description,
			"assignee_id":          req.AssigneeID,
			"labels":               req.Labels,
			"remove_source_branch": req.RemoveSourceBranch,
			"squash":               req.Squash,
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
