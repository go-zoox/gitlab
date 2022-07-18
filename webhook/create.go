package webhook

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "projects/:project_id/hooks",
}

type CreateRequest struct {
	ProjectID int64  `json:"project_id"`
	URL       string `json:"url"`
	Token     string `json:"token"`
	//
	PushEvents          bool `json:"push_events"`
	TagPushEvents       bool `json:"tag_push_events"`
	MergeRequestsEvents bool `json:"merge_requests_events"`
	//
	ReleasesEvents   bool `json:"release_events"`
	IssuesEvents     bool `json:"issues_events"`
	DeploymentEvents bool `json:"deployment_events"`
	//
	EnableSSLVerification bool `json:"enable_ssl_verification"`
}

type CreateResponse = WebHook

func Create(req *CreateRequest) (*CreateResponse, error) {
	response, err := request.Request(CreateConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(req.ProjectID)),
		},
		Body: map[string]interface{}{
			"url":                     req.URL,
			"token":                   req.Token,
			"push_events":             req.PushEvents,
			"tag_push_events":         req.TagPushEvents,
			"merge_requests_events":   req.MergeRequestsEvents,
			"release_events":          req.ReleasesEvents,
			"issues_events":           req.IssuesEvents,
			"deployment_events":       req.DeploymentEvents,
			"enable_ssl_verification": req.EnableSSLVerification,
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
