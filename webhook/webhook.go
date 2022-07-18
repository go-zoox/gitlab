package webhook

import (
	"time"

	"github.com/go-zoox/gitlab/client"
)

type WebHook struct {
	client client.Client
	//
	ID                       int64     `json:"id"`
	ProjectID                int64     `json:"project_id"`
	URL                      string    `json:"url"`
	CreatedAt                time.Time `json:"created_at"`
	PushEvents               bool      `json:"push_events"`
	TagPushEvents            bool      `json:"tag_push_events"`
	IssuesEvents             bool      `json:"issues_events"`
	ConfidentialIssuesEvents bool      `json:"confidential_issues_events"`
	MergeRequestsEvents      bool      `json:"merge_requests_events"`
	ReleasesEvents           bool      `json:"release_events"`
	RepositoryUpdateEvents   bool      `json:"repository_update_events"`
	NoteEvents               bool      `json:"note_events"`
	PipelineEvents           bool      `json:"pipeline_events"`
	WikiPageEvents           bool      `json:"wiki_page_events"`
	DeploymentEvents         bool      `json:"deployment_events"`
	JobEvents                bool      `json:"job_events"`
	PushEventsBranchFilter   string    `json:"push_events_branch_filter"`
	EnableSSLVerification    bool      `json:"enable_ssl_verification"`
}

func New(client client.Client) *WebHook {
	return &WebHook{
		client: client,
	}
}

func (w *WebHook) List(cfg *ListRequest) (*[]WebHook, error) {
	return List(w.client, cfg)
}

func (w *WebHook) Create(req *CreateRequest) (*WebHook, error) {
	return Create(w.client, req)
}

func (w *WebHook) Get(req *GetRequest) (*WebHook, error) {
	return Get(w.client, req)
}

func (w *WebHook) Delete(req *DeleteRequest) error {
	return Delete(w.client, req)
}
