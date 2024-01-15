package user

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestList(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := List(client.NewMockClient(), func(opt *ListRequest) {
		opt.Page = 6
		// opt.PerPage = 1000
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo, repo.IsLastPage)
}
