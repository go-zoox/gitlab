package branch

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestGet(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Get(client.NewMockClient(), &GetRequest{
		RepositoryID: 3,
		Name:         "fix/test-create-branch",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
