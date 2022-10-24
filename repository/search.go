package repository

import (
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/request"
)

var SearchConfig = &request.Config{
	Action:   "GET",
	Resource: "search",
}

type SearchRequest struct {
	// [REQUIRED] Search Keyword
	Keyword string `json:"keyword"`
}

type SearchResponse = []Repository

func Search(client client.Client, cfg *SearchRequest) (*SearchResponse, error) {
	response, err := client.Request(SearchConfig, &request.Payload{
		Query: map[string]string{
			"scope":  "projects",
			"search": cfg.Keyword,
		},
	})

	if err != nil {
		return nil, err
	}

	var res SearchResponse
	if err = response.UnmarshalJSON(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
