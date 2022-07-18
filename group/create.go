package group

import (
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var CreateConfig = &request.Config{
	Action:   "POST",
	Resource: "groups",
}

type CreateRequest struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"` // private, internal, public, default: private
}

type CreateResponse = Group

func Create(client client.Client, req *CreateRequest) (*CreateResponse, error) {
	if req.Visibility == "" {
		req.Visibility = "private"
	}

	response, err := client.Request(CreateConfig, &request.Payload{
		Body: map[string]interface{}{
			"name":        req.Name,
			"path":        req.Path,
			"description": req.Description,
			"visibility":  req.Visibility,
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
