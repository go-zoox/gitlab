package gitlab

import (
	"os"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/gitlab/branch"
	"github.com/go-zoox/gitlab/config"
	"github.com/go-zoox/gitlab/group"
	"github.com/go-zoox/gitlab/merge_request"
	"github.com/go-zoox/gitlab/repository"
	"github.com/go-zoox/gitlab/request"
	"github.com/go-zoox/gitlab/user"
	"github.com/go-zoox/gitlab/webhook"
)

type Client struct {
	Cfg *Config
}

type Config struct {
	Service string
	Token   string
	Version string
}

func New(cfg ...*Config) (*Client, error) {
	cfgX := &Config{}
	if len(cfg) > 0 {
		cfgX = cfg[0]

		if cfgX == nil {
			panic("config is nil")
		}
	}

	dotEnvPath := fs.JoinPath(fs.CurrentDir(), ".env")
	if fs.IsExist(dotEnvPath) {
		if err := LoadConfig(); err != nil {
			return nil, err
		}
	}

	if cfgX.Service == "" {
		if os.Getenv("GITLAB_SERVICE") != "" {
			cfgX.Service = os.Getenv("GITLAB_SERVICE")
		} else {
			panic("service is required")
		}
	}

	if cfgX.Token == "" {
		if os.Getenv("GITLAB_TOKEN") != "" {
			cfgX.Token = os.Getenv("GITLAB_TOKEN")
		} else {
			panic("token is required")
		}
	}

	return &Client{
		Cfg: cfgX,
	}, nil
}

func (c *Client) Request(cfg *request.Config, payload *request.Payload) (*request.Response, error) {
	cfg.Service = c.Cfg.Service
	cfg.Token = c.Cfg.Token
	cfg.Version = c.Cfg.Version

	return request.Request(cfg, payload)
}

func (c *Client) Repository() repository.RepositoryImpl {
	return repository.New(c)
}

func (c *Client) Branch() branch.BranchImpl {
	return branch.New(c)
}

func (c *Client) MergeRequest() merge_request.MergeRequestImpl {
	return merge_request.New(c)
}

func (c *Client) WebHook() webhook.WebHookImpl {
	return webhook.New(c)
}

func (c *Client) Group() group.GroupImpl {
	return group.New(c)
}

func (c *Client) User() user.UserImpl {
	return user.New(c)
}

func LoadConfig() error {
	return config.Load()
}
