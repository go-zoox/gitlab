package branch

import (
	"testing"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestDelete(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	if err := Delete(client.NewMockClient(), &DeleteRequest{
		RepositoryID: 3,
		Name:         "feat/test-from-gosdk",
	}); err != nil {
		t.Fatal(err)
	}
}
