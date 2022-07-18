package webhook

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

	repo, err := List(&client.MockClient{}, &ListRequest{
		ProjectID: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
