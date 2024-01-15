package user

import "github.com/go-zoox/gitlab/client"

type UserImpl interface {
	List(opts ...func(opt *ListRequest)) (*ListResponse, error)
	Get(opts ...func(opt *GetRequest)) (*GetResponse, error)
	//
	ListALL(opts ...func(opt *ListALLRequest)) (*ListALLResponse, error)
}

type User struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	State           string `json:"state"`
	AvatarURL       string `json:"avatar_url"`
	WebURL          string `json:"web_url"`
	CreatedAt       string `json:"created_at"`
	IsAdmin         bool   `json:"is_admin"`
	Bio             string `json:"bio"`
	Location        string `json:"location"`
	Skype           string `json:"skype"`
	Linkedin        string `json:"linkedin"`
	Twitter         string `json:"twitter"`
	WebsiteURL      string `json:"website_url"`
	Organization    string `json:"organization"`
	LastSignInAt    string `json:"last_sign_in_at"`
	ConfirmedAt     string `json:"confirmed_at"`
	ThemeID         int64  `json:"theme_id"`
	LastActivityOn  string `json:"last_activity_on"`
	ColorSchemeID   int64  `json:"color_scheme_id"`
	ProjectsLimit   int64  `json:"projects_limit"`
	CurrentSignInAt string `json:"current_sign_in_at"`
	Identities      []struct {
		Provider  string `json:"provider"`
		ExternUID string `json:"extern_uid"`
	} `json:"identities"`
	CanCreateGroup   bool `json:"can_create_group"`
	CanCreateProject bool `json:"can_create_project"`
	TwoFactorEnabled bool `json:"two_factor_enabled"`
	External         bool `json:"external"`
	PrivateProfile   bool `json:"private_profile"`
}

type user struct {
	client client.Client
}

func New(client client.Client) UserImpl {
	return &user{
		client: client,
	}
}

func (u *user) List(opts ...func(opt *ListRequest)) (*ListResponse, error) {
	return List(u.client, opts...)
}

func (u *user) Get(opts ...func(opt *GetRequest)) (*GetResponse, error) {
	return Get(u.client, opts...)
}

func (u *user) ListALL(opts ...func(opt *ListALLRequest)) (*ListALLResponse, error) {
	return ListALL(u.client, opts...)
}
