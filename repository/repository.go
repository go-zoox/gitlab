package repository

import (
	"time"

	"github.com/go-zoox/gitlab/client"
)

type RepositoryImpl interface {
	List(cfg *ListRequest) (*[]Repository, error)
	Create(req *CreateRequest) (*Repository, error)
	Get(name string) (*Repository, error)
	Delete(name string) error
}

type Repository struct {
	client client.Client
	//
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Path          string    `json:"path_with_namespace"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	DefaultBranch string    `json:"default_branch"`
	SSHURL        string    `json:"ssh_url_to_repo"`
	HTTPURL       string    `json:"http_url_to_repo"`
	WebURL        string    `json:"web_url"`
	READMEURL     string    `json:"readme_url"`
	//
	LastActivityAt time.Time `json:"last_activity_at"`
	Visibility     string    `json:"visibility"` // internal, public, internal_only
	//
	Namespace Namespace `json:"namespace"`
}

type Namespace struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Kind string `json:"kind"` // group, user
	// FullPath  string `json:"full_path"`
	ParentID  int64  `json:"parent_id"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

func New(client client.Client) RepositoryImpl {
	return &Repository{
		client: client,
	}
}

func (r *Repository) List(cfg *ListRequest) (*[]Repository, error) {
	return List(r.client, cfg)
}

func (r *Repository) Create(req *CreateRequest) (*Repository, error) {
	return Create(r.client, req)
}

func (r *Repository) Get(name string) (*Repository, error) {
	return Get(r.client, name)
}

func (r *Repository) Delete(name string) error {
	return Delete(r.client, name)
}
