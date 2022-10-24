package repository

import (
	"fmt"
	"strings"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/group"
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
	// NamespaceID          string `json:"namespace_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	InitializeWithREADME bool   `json:"initialize_with_readme"`
	DefaultBranch        string `json:"default_branch"`
}

type CreateResponse = Repository

func Create(client client.Client, req *CreateRequest) (*CreateResponse, error) {
	parts := strings.Split(req.Name, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repository name: %s, must be group/name", req.Name)
	}

	// @TODO current only support group, but should support user
	// 	here group id is equal to namespace id
	var namespaceID int64
	groupName, repositoryName := parts[0], parts[1]
	groups, err := group.Search(client, groupName)
	if err != nil || len(*groups) == 0 {
		return nil, fmt.Errorf("invalid group name: %s (error: %v)", groupName, err)
	}
	for _, group := range *groups {
		if group.Name == groupName {
			namespaceID = group.ID
		}
	}

	response, err := client.Request(CreateConfig, &request.Payload{
		Body: map[string]interface{}{
			"namespace_id":           namespaceID,
			"name":                   repositoryName,
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
