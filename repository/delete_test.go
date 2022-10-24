package repository

import (
	"testing"

	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestDelete(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	if err := Delete(&client.MockClient{}, "eunomia/test-from-gosdk"); err != nil {
		t.Fatal(err)
	}
}
