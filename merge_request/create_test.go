package merge_request

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

	repo, err := Create(&client.MockClient{}, &CreateRequest{
		ProjectID:    3,
		SourceBranch: "fix/test-create-branch2",
		TargetBranch: "master",
		Title:        "test-create-branch2 merge request",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
