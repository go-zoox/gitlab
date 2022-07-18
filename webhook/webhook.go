package webhook

import "time"

type WebHook struct {
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

func New() *WebHook {
	return &WebHook{}
}

func (r *WebHook) Create(req *CreateRequest) (*WebHook, error) {
	return Create(req)
}

func (r *WebHook) Get(req *GetRequest) (*WebHook, error) {
	return Get(req)
}

func (r *WebHook) Delete(req *DeleteRequest) error {
	return Delete(req)
}
