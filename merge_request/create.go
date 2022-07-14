package merge_request

import (
	"strconv"

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
	ProjectID    int64  `json:"project_id"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Title        string `json:"title"`
	//
	Description string `json:"description"`
	AssigneeID  int64  `json:"assignee_id"`
	Labels      string `json:"labels"`
	//
	RemoveSourceBranch bool `json:"remove_source_branch"`
}

type CreateResponse = MergeRequest

func Create(req *CreateRequest) (*CreateResponse, error) {
	response, err := request.Request(CreateConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(req.ProjectID)),
		},
		Query: map[string]string{
			"source_branch": req.SourceBranch,
			"target_branch": req.TargetBranch,
			"title":         req.Title,
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
