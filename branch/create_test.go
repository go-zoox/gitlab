package branch

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/config"
)

func TestCreate(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Create(&CreateRequest{
		ProjectID: 3,
		Name:      "feat/test-from-gosdk",
		Ref:       "master",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
