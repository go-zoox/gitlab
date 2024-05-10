package merge_request

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestMerge(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Merge(&client.MockClient{}, &MergeeRequest{
		RepositoryID:   3,
		MergeRequestID: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
