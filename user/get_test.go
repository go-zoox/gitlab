package user

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/gitlab/client"
	"github.com/go-zoox/gitlab/config"
)

func TestGet(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	user, err := Get(client.NewMockClient(), func(opt *GetRequest) {
		opt.ID = 2
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(user)
}
