package group

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/config"
)

func TestSearch(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Search("the_new_group")
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
