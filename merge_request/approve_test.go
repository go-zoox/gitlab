package merge_request

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestApprove(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	repo, err := Approve(&client.MockClient{}, &ApproveRequest{
		ProjectID:      3,
		MergeRequestID: 2,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
