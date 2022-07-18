package merge_request

import (
	"strconv"
	"time"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var ApproveConfig = &request.Config{
	Action:   "POST",
	Resource: "projects/:project_id/merge_requests/:merge_request_id/approve",
}

type ApproveRequest struct {
	ProjectID      int64 `json:"project_id"`
	MergeRequestID int64 `json:"merge_request_id"`
}

type ApproveResponse struct {
	ID          int64     `json:"id"`
	IID         int64     `json:"iid"`
	ProjectID   int64     `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// opened, closed, merged, locked
	State string `json:"state"`
	// unchecked, checking, can_be_merged, cannot_be_merged, cannot_be_merged_recheck
	MergeStatus string `json:"merge_status"`
	//
	Approve bool `json:"approve"`
	// ApprovedBy User `json:"approved_by"`
}

func Approve(client client.Client, req *ApproveRequest) (*ApproveResponse, error) {
	response, err := client.Request(ApproveConfig, &request.Payload{
		Params: map[string]string{
			"project_id":       strconv.Itoa(int(req.ProjectID)),
			"merge_request_id": strconv.Itoa(int(req.MergeRequestID)),
		},
	})

	if err != nil {
		return nil, err
	}

	var res ApproveResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
