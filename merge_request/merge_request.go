package merge_request

import (
	"time"

	"github.com/go-zoox/gitlab/client"
)

type MergeRequestImpl interface {
	Create(req *CreateRequest) (*MergeRequest, error)
	Get(req *GetRequest) (*MergeRequest, error)
	Delete(req *DeleteRequest) error
	Approve(req *ApproveRequest) (*ApproveResponse, error)
	Merge(req *MergeeRequest) (*MergeResponse, error)
}

type MergeRequest struct {
	client client.Client
	//
	ID              int64     `json:"id"`
	IID             int64     `json:"iid"`
	ProjectID       int64     `json:"project_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	SourceBranch    string    `json:"source_branch"`
	TargetBranch    string    `json:"target_branch"`
	SourceProjectID int64     `json:"source_project_id"`
	TargetProjectID int64     `json:"target_project_id"`
	Labels          []string  `json:"labels"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	WebURL          string    `json:"web_url"`
	//
	Author User `json:"author"`
	//
	MergedAt   time.Time `json:"merged_at"`
	MergedUser User      `json:"merged_by"`
	// opened, closed, merged, locked
	State string `json:"state"`
	// unchecked, checking, can_be_merged, cannot_be_merged, cannot_be_merged_recheck
	MergeStatus  string `json:"merge_status"`
	MergeError   string `json:"merge_error"`
	HasConflicts bool   `json:"has_conflicts"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

func New(client client.Client) MergeRequestImpl {
	return &MergeRequest{
		client: client,
	}
}

func (r *MergeRequest) Create(req *CreateRequest) (*MergeRequest, error) {
	return Create(r.client, req)
}

func (r *MergeRequest) Get(req *GetRequest) (*MergeRequest, error) {
	return Get(r.client, req)
}

func (r *MergeRequest) Delete(req *DeleteRequest) error {
	return Delete(r.client, req)
}

func (r *MergeRequest) Approve(req *ApproveRequest) (*ApproveResponse, error) {
	return Approve(r.client, req)
}

func (r *MergeRequest) Merge(req *MergeeRequest) (*MergeResponse, error) {
	return Merge(r.client, req)
}
