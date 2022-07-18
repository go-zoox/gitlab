package group

import (
	"time"

	"github.com/go-zoox/gitlab/client"
)

type Group struct {
	client client.Client
	//
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	WebURL      string    `json:"web_url"`
	FullName    string    `json:"full_name"`
	FullPath    string    `json:"full_path"`
	CreatedAt   time.Time `json:"created_at"`
	ParentID    int64     `json:"parent_id"`
	//
	Visibility                     string `json:"visibility"` // internal, public, internal_only
	ShareWithGroupLock             bool   `json:"share_with_group_lock"`
	RequireTwoFactorAuthentication bool   `json:"require_two_factor_authentication"`
	TwoFactorGracePeriod           int    `json:"two_factor_grace_period"`
	ProjectCreationLevel           string `json:"project_creation_level"` // owner, developer, reporter, guest
	AutoDevOpsEnabled              bool   `json:"auto_devops_enabled"`
	SubgroupCreationLevel          string `json:"subgroup_creation_level"` // owner, developer, reporter, guest
	EmailsDisabled                 bool   `json:"emails_disabled"`
	MentionsDisabled               bool   `json:"mentions_disabled"`
	LFSEnabled                     bool   `json:"lfs_enabled"`
	DefaultBranchProtection        int    `json:"default_branch_protection"` // 0, 1, 2
	AvatarURL                      string `json:"avatar_url"`
	RequestAccessEnabled           bool   `json:"request_access_enabled"`
	LDAPCN                         string `json:"ldap_cn"`
	LDAPAccess                     string `json:"ldap_access"`
}

func New(client client.Client) *Group {
	return &Group{
		client: client,
	}
}

func (g *Group) List(cfg *ListRequest) (*[]Group, error) {
	return List(g.client, cfg)
}

func (g *Group) Get(cfg *GetRequest) (*Group, error) {
	return Get(g.client, cfg)
}

func (g *Group) Search(keyword string) (*[]Group, error) {
	return Search(g.client, keyword)
}
