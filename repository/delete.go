package repository

import (
	"strconv"

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

func Delete(id int64) error {
	_, err := request.Request(DeleteConfig, &request.Payload{
		Params: map[string]string{
			"project_id": strconv.Itoa(int(id)),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
