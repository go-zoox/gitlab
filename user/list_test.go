package user

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestListALL(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := ListALL(client.NewMockClient())
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo.Data, repo.Total)
}
