package repository

import "time"

type Repository struct {
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
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Create(req *CreateRequest) (*Repository, error) {
	return Create(req)
}

func (r *Repository) Get(projectID int64) (*Repository, error) {
	return Get(projectID)
}

func (r *Repository) Delete(projectID int64) error {
	return Delete(projectID)
}
