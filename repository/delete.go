package repository

import (
	"fmt"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var DeleteConfig = &request.Config{
	Action:   "DELETE",
	Resource: "projects/:project_id",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

type DeleteResponse struct {
	Repository
}

func Delete(client client.Client, name string) error {
	project, err := Get(client, name)
	if err != nil {
		return err
	}

	_, err = client.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id": fmt.Sprintf("%d", project.ID),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
