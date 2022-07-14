package branch

import (
	"strconv"

	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "projects/:project_id/repository/branches",
}

type CreateRequest struct {
	ProjectID int64  `json:"project_id"`
	Name      string `json:"name"`
	Ref       string `json:"ref"`
}

type CreateResponse = Branch

func Create(req *CreateRequest) (*CreateResponse, error) {
	response, err := request.Request(CreateConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(req.ProjectID)),
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
