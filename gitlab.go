package gitlab

import (
	"github.com/go-zoox/gitlab/branch"
	"github.com/go-zoox/gitlab/config"
	"github.com/go-zoox/gitlab/repository"
)

type Client struct {
	Config *Config
}

type Config struct {
	Token string
}

func New(cfg *Config) *Client {
	if cfg == nil {
		panic("config is nil")
	}

	if cfg.Token == "" {
		panic("config.SeTokencretId is required")
	}

	return &Client{
		Config: cfg,
		// @TODO
		// Cvm: cvm.New(&cvm.Config{
		// 	SecretId:  config.SecretId,
		// 	SecretKey: config.SecretKey,
		// 	Region:    config.Region,
		// }),
	}
}

func (c *Client) Repository() *repository.Repository {
	return repository.New()
}

func (c *Client) Branch() *branch.Branch {
	return branch.New()
}

func LoadConfig() error {
	return config.Load()
}
