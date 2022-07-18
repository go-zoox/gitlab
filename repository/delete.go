package repository

import (
	"strconv"

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

func Delete(client client.Client, id int64) error {
	_, err := client.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(id)),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
