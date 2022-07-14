package branch

import "time"

type Branch struct {
	Name               string `json:"name"`
	WebURL             string `json:"web_url"`
	CanPush            bool   `json:"can_push"`
	IsDefault          bool   `json:"default"`
	IsProtected        bool   `json:"protected"`
	IsMerged           bool   `json:"merged"`
	DevelopersCanPush  bool   `json:"developers_can_push"`
	DevelopersCanMerge bool   `json:"developers_can_merge"`
	//
	Commit Commit `json:"commit"`
}

type Commit struct {
	ID             string    `json:"id"`
	ShortID        string    `json:"short_id"`
	Title          string    `json:"title"`
	Author         string    `json:"author_name"`
	Message        string    `json:"message"`
	AuthorEmail    string    `json:"author_email"`
	AuthoredDate   time.Time `json:"authored_date"`
	Committer      string    `json:"committer_name"`
	CommitterEmail string    `json:"committer_email"`
	CommittedDate  time.Time `json:"committed_date"`
	WebURL         string    `json:"web_url"`
	CreatedAt      time.Time `json:"created_at"`
	ParentIDs      []string  `json:"parent_ids"`
}

func New() *Branch {
	return &Branch{}
}

func (r *Branch) Create(req *CreateRequest) (*Branch, error) {
	return Create(req)
}

func (r *Branch) Get(req *GetRequest) (*Branch, error) {
	return Get(req)
}

func (r *Branch) Delete(req *DeleteRequest) error {
	return Delete(req)
}
