package merge_request

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id/merge_requests/:merge_request_id",
}

type GetRequest struct {
	ProjectID      int64 `json:"project_id"`
	MergeRequestID int64 `json:"merge_request_id"`
}

type GetResponse = MergeRequest

func Get(cfg *GetRequest) (*GetResponse, error) {
	response, err := request.Request(GetConfig, &request.Payload{
		Params: map[string]string{
			"project_id":       strconv.Itoa(int(cfg.ProjectID)),
			"merge_request_id": strconv.Itoa(int(cfg.MergeRequestID)),
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
