package repository

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
		NamespaceID: "2",
		Name:        "test-from-gosdk",
		Description: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(repo)
}
