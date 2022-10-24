package branch

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestCreate(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Create(client.NewMockClient(), &CreateRequest{
		RepositoryID: 445,
		Name:         "sprint_dev",
		Ref:          "master",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}

func TestCreateByProjectName(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Create(client.NewMockClient(), &CreateRequest{
		RepositoryName: "eunomia/eunomia-frontend",
		Name:           "sprint_dev",
		Ref:            "master",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
