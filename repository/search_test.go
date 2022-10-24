package repository

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestSearch(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Search(&client.MockClient{}, &SearchRequest{
		Keyword: "eunomia/eunomia-frontend",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
