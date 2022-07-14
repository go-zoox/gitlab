package merge_request

import (
	"testing"

	"github.com/go-zoox/gitlab/config"
)

func TestDelete(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	if err := Delete(&DeleteRequest{
		ProjectID:      3,
		MergeRequestID: 1,
	}); err != nil {
		t.Fatal(err)
	}
}
