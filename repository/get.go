package repository

import (
	"fmt"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var GetConfig = &request.Config{
	Action:   "GET",
	Resource: "projects/:project_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type GetResponse = Repository

func Get(client client.Client, name string) (*GetResponse, error) {
	list, err := Search(client, &SearchRequest{
		Keyword: name,
	})

	if err != nil {
		return nil, err
	}

	for _, one := range *list {
		if one.Path == name {
			return &one, nil
		}
	}

	return nil, fmt.Errorf("repository(%s) not found", name)
}
