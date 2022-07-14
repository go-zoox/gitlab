package repository

import (
	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "projects",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type CreateRequest struct {
	NamespaceID          string `json:"namespace_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	InitializeWithREADME bool   `json:"initialize_with_readme"`
	DefaultBranch        string `json:"default_branch"`
}

type CreateResponse = Repository

func Create(req *CreateRequest) (*CreateResponse, error) {
	response, err := request.Request(CreateConfig, &request.Payload{
		Body: map[string]interface{}{
			"namespace_id":           req.NamespaceID,
			"name":                   req.Name,
			"description":            req.Description,
			"initialize_with_readme": req.InitializeWithREADME,
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
