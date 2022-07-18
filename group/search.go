package group

import "github.com/go-zoox/gitlab/client"

func Search(client client.Client, keyword string) (*[]Group, error) {
	return List(client, &ListRequest{
		Search: keyword,
	})
}
