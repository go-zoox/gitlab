package merge_request

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/repository"
	"github.com/go-zoox/gitlab/request"
)

var MergeConfig = &request.Config{
	Action:   "PUT",
	Resource: "projects/:project_id/merge_requests/:merge_request_id/merge",
}

type MergeeRequest struct {
	RepositoryID int64 `json:"repository_id"`
	// the name of the project
	RepositoryName string `json:"repository_name"`
	//
	MergeRequestID int64 `json:"merge_request_id"`
	//
	MergeCommitMessage string `json:"merge_commit_message"`
	//
	Squash              bool   `json:"squash"`
	SquashCommitMessage string `json:"squash_commit_message"`
	//
	ShouldRemoveSourceBranch bool `json:"should_remove_source_branch"`
}

type MergeResponse = MergeRequest

func Merge(client client.Client, req *MergeeRequest) (*MergeResponse, error) {
	if req.RepositoryName != "" && req.RepositoryID == 0 {
		repository, err := repository.Get(client, req.RepositoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to get repository by name: %v", err)
		}

		req.RepositoryID = repository.ID
	}

	// fmt.Println("req.RepositoryID", req.RepositoryID)
	// fmt.Println("req.MergeRequestID", req.MergeRequestID)

	response, err := client.Request(MergeConfig, &request.Payload{
		Params: map[string]string{
			"project_id":       strconv.Itoa(int(req.RepositoryID)),
			"merge_request_id": strconv.Itoa(int(req.MergeRequestID)),
		},
		Body: map[string]any{
			"merge_commit_message":        req.MergeCommitMessage,
			"squash":                      req.Squash,
			"squash_commit_message":       req.SquashCommitMessage,
			"should_remove_source_branch": req.ShouldRemoveSourceBranch,
		},
	})
	if err != nil {
		return nil, err
	}

	var res MergeResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	if res.MergeError != "" {
		return nil, fmt.Errorf(res.MergeError)
	}

	return &res, nil
}
