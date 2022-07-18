package merge_request

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/config"
)

func TestIsConflict(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := IsConflict(&IsConflictRequest{
		ProjectID:      3,
		MergeRequestID: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
