package merge_request

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var MergeConfig = &request.Config{
	Action:   "PUT",
	Resource: "projects/:project_id/merge_requests/:merge_request_id/merge",
}

type MergeeRequest struct {
	ProjectID      int64 `json:"project_id"`
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

func Merge(req *MergeeRequest) (*MergeResponse, error) {
	response, err := request.Request(MergeConfig, &request.Payload{
		Params: map[string]string{
			"project_id":       strconv.Itoa(int(req.ProjectID)),
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
