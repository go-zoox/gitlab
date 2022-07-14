package repository

import (
	"testing"

	"github.com/go-zoox/gitlab/config"
)

func TestDelete(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	if err := Delete(6); err != nil {
		t.Fatal(err)
	}
}
