package webhook

// {
//   "method": "POST",
//   "url": "/",
//   "query": {},
//   "body": {
//     "object_kind": "push",
//     "event_name": "push",
//     "before": "0000000000000000000000000000000000000000",
//     "after": "bf850ae470b59d53ece6ac344adb9fba25ddcfd8",
//     "ref": "refs/heads/feat/test-from-gosdk",
//     "checkout_sha": "bf850ae470b59d53ece6ac344adb9fba25ddcfd8",
//     "message": null,
//     "user_id": 1,
//     "user_name": "Administrator",
//     "user_username": "root",
//     "user_email": null,
//     "user_avatar": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
//     "project_id": 3,
//     "project": {
//       "id": 3,
//       "name": "test-nodejs",
//       "description": null,
//       "web_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//       "avatar_url": null,
//       "git_ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//       "git_http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git",
//       "namespace": "GitLab Instance",
//       "visibility_level": 0,
//       "path_with_namespace": "gitlab-instance-32797568/test-nodejs",
//       "default_branch": "master",
//       "ci_config_path": null,
//       "homepage": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//       "url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//       "ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//       "http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git"
//     },
//     "commits": [],
//     "total_commits_count": 0,
//     "push_options": {},
//     "repository": {
//       "name": "test-nodejs",
//       "url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//       "description": null,
//       "homepage": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//       "git_http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git",
//       "git_ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//       "visibility_level": 0
//     }
//   },
//   "headers": {
//     "content-type": "application/json",
//     "user-agent": "GitLab/14.10.2-ee",
//     "x-gitlab-event": "Push Hook",
//     "x-gitlab-token": "the_webhook_token",
//     "x-gitlab-event-uuid": "75a98ce1-d20a-420c-81b9-a633ef90dde6",
//     "connection": "close",
//     "host": "10.208.200.121:8080",
//     "content-length": "1653"
//   },
//   "origin": "http://10.208.200.121:8080"
// }

type RequestHeader struct {
	UserAgent string `header:"user-agent"`
	EventName string `header:"x-gitlab-event"`
	Token     string `header:"x-gitlab-token"`
	EventUUID string `header:"x-gitlab-event-uuid"`
}

type RequestBody struct {
	ObjectKind   string `json:"object_kind"` // push, tag_push, etc
	EventName    string `json:"event_name"`  // push, tag_push, etc
	Before       string `json:"before"`      // commit sha before push
	After        string `json:"after"`       // commit sha after push
	Ref          string `json:"ref"`         // ref name
	CheckoutSha  string `json:"checkout_sha"`
	Message      string `json:"message"`
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserAvatar   string `json:"user_avatar"`
	ProjectID    int    `json:"project_id"`
	Project      struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
		CiConfigPath      string `json:"ci_config_path"`
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []string `json:"added"`
		Modified []string `json:"modified"`
		Removed  []string `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
		Force bool `json:"force"`
	} `json:"push_options"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}
