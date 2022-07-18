package gitlab

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/group"
)

func TestNew(t *testing.T) {
	// client, _ := New(&Config{
	// 	// Service: "https://gitlab.com",
	// 	// Token:   "token",
	// 	// Version: "v4",
	// })

	client, _ := New()

	repository := client.Repository()
	if repository == nil {
		t.Fatal("repository is nil")
	}

	branch := client.Branch()
	if branch == nil {
		t.Fatal("branch is nil")
	}

	mergeRequest := client.MergeRequest()
	if mergeRequest == nil {
		t.Fatal("mergeRequest is nil")
	}

	webhook := client.WebHook()
	if webhook == nil {
		t.Fatal("webhook is nil")
	}

	g := client.Group()
	if g == nil {
		t.Fatal("group is nil")
	}

	fmt.PrintJSON(g.List(&group.ListRequest{}))
}
