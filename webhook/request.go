package webhook

import "time"

type RequestHeader struct {
	UserAgent string `header:"user-agent"`
	EventName string `header:"x-gitlab-event"` // Push Hook, System Hook,
	Token     string `header:"x-gitlab-token"`
	EventUUID string `header:"x-gitlab-event-uuid"`
}

//	{
//	  "method": "POST",
//	  "url": "/",
//	  "query": {},
//	  "body": {
//	    "object_kind": "push",
//	    "event_name": "push",
//	    "before": "0000000000000000000000000000000000000000",
//	    "after": "bf850ae470b59d53ece6ac344adb9fba25ddcfd8",
//	    "ref": "refs/heads/feat/test-from-gosdk",
//	    "checkout_sha": "bf850ae470b59d53ece6ac344adb9fba25ddcfd8",
//	    "message": null,
//	    "user_id": 1,
//	    "user_name": "Administrator",
//	    "user_username": "root",
//	    "user_email": null,
//	    "user_avatar": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
//	    "project_id": 3,
//	    "project": {
//	      "id": 3,
//	      "name": "test-nodejs",
//	      "description": null,
//	      "web_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//	      "avatar_url": null,
//	      "git_ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//	      "git_http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git",
//	      "namespace": "GitLab Instance",
//	      "visibility_level": 0,
//	      "path_with_namespace": "gitlab-instance-32797568/test-nodejs",
//	      "default_branch": "master",
//	      "ci_config_path": null,
//	      "homepage": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//	      "url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//	      "ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//	      "http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git"
//	    },
//	    "commits": [],
//	    "total_commits_count": 0,
//	    "push_options": {},
//	    "repository": {
//	      "name": "test-nodejs",
//	      "url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//	      "description": null,
//	      "homepage": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs",
//	      "git_http_url": "http://tgit.shushu.work/gitlab-instance-32797568/test-nodejs.git",
//	      "git_ssh_url": "git@tgit.shushu.work:gitlab-instance-32797568/test-nodejs.git",
//	      "visibility_level": 0
//	    }
//	  },
//	  "headers": {
//	    "content-type": "application/json",
//	    "user-agent": "GitLab/14.10.2-ee",
//	    "x-gitlab-event": "Push Hook",
//	    "x-gitlab-token": "the_webhook_token",
//	    "x-gitlab-event-uuid": "75a98ce1-d20a-420c-81b9-a633ef90dde6",
//	    "connection": "close",
//	    "host": "10.208.200.121:8080",
//	    "content-length": "1653"
//	  },
//	  "origin": "http://10.208.200.121:8080"
//	}
type RequestPushBody struct {
	ObjectKind   string `json:"object_kind"` // push
	EventName    string `json:"event_name"`  // push
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

//	{
//	  "method": "POST",
//	  "url": "/",
//	  "query": {},
//	  "body": {
//	    "object_kind": "merge_request",
//	    "event_type": "merge_request",
//	    "user": {
//	      "id": 1,
//	      "name": "Administrator",
//	      "username": "root",
//	      "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
//	      "email": "[REDACTED]"
//	    },
//	    "project": {
//	      "id": 9,
//	      "name": "xxx",
//	      "description": "",
//	      "web_url": "http://tgit.shushu.work/test-group/xxx",
//	      "avatar_url": null,
//	      "git_ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	      "git_http_url": "http://tgit.shushu.work/test-group/xxx.git",
//	      "namespace": "test-group",
//	      "visibility_level": 0,
//	      "path_with_namespace": "test-group/xxx",
//	      "default_branch": "main",
//	      "ci_config_path": null,
//	      "homepage": "http://tgit.shushu.work/test-group/xxx",
//	      "url": "git@tgit.shushu.work:test-group/xxx.git",
//	      "ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	      "http_url": "http://tgit.shushu.work/test-group/xxx.git"
//	    },
//	    "object_attributes": {
//	      "assignee_id": null,
//	      "author_id": 1,
//	      "created_at": "2022-07-18 08:35:00 UTC",
//	      "description": "",
//	      "head_pipeline_id": null,
//	      "id": 4,
//	      "iid": 1,
//	      "last_edited_at": null,
//	      "last_edited_by_id": null,
//	      "merge_commit_sha": null,
//	      "merge_error": null,
//	      "merge_params": {
//	        "force_remove_source_branch": "1"
//	      },
//	      "merge_status": "preparing",
//	      "merge_user_id": null,
//	      "merge_when_pipeline_succeeds": false,
//	      "milestone_id": null,
//	      "source_branch": "root-main-patch-62946",
//	      "source_project_id": 9,
//	      "state_id": 1,
//	      "target_branch": "main",
//	      "target_project_id": 9,
//	      "time_estimate": 0,
//	      "title": "Update README.md",
//	      "updated_at": "2022-07-18 08:35:00 UTC",
//	      "updated_by_id": null,
//	      "url": "http://tgit.shushu.work/test-group/xxx/-/merge_requests/1",
//	      "source": {
//	        "id": 9,
//	        "name": "xxx",
//	        "description": "",
//	        "web_url": "http://tgit.shushu.work/test-group/xxx",
//	        "avatar_url": null,
//	        "git_ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "git_http_url": "http://tgit.shushu.work/test-group/xxx.git",
//	        "namespace": "test-group",
//	        "visibility_level": 0,
//	        "path_with_namespace": "test-group/xxx",
//	        "default_branch": "main",
//	        "ci_config_path": null,
//	        "homepage": "http://tgit.shushu.work/test-group/xxx",
//	        "url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "http_url": "http://tgit.shushu.work/test-group/xxx.git"
//	      },
//	      "target": {
//	        "id": 9,
//	        "name": "xxx",
//	        "description": "",
//	        "web_url": "http://tgit.shushu.work/test-group/xxx",
//	        "avatar_url": null,
//	        "git_ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "git_http_url": "http://tgit.shushu.work/test-group/xxx.git",
//	        "namespace": "test-group",
//	        "visibility_level": 0,
//	        "path_with_namespace": "test-group/xxx",
//	        "default_branch": "main",
//	        "ci_config_path": null,
//	        "homepage": "http://tgit.shushu.work/test-group/xxx",
//	        "url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "ssh_url": "git@tgit.shushu.work:test-group/xxx.git",
//	        "http_url": "http://tgit.shushu.work/test-group/xxx.git"
//	      },
//	      "last_commit": {
//	        "id": "8fe782d9c4ef760e51274ba27f39377fc7d0223e",
//	        "message": "Update README.md",
//	        "title": "Update README.md",
//	        "timestamp": "2022-07-18T08:34:42+00:00",
//	        "url": "http://tgit.shushu.work/test-group/xxx/-/commit/8fe782d9c4ef760e51274ba27f39377fc7d0223e",
//	        "author": {
//	          "name": "Administrator",
//	          "email": "admin@example.com"
//	        }
//	      },
//	      "work_in_progress": false,
//	      "total_time_spent": 0,
//	      "time_change": 0,
//	      "human_total_time_spent": null,
//	      "human_time_change": null,
//	      "human_time_estimate": null,
//	      "assignee_ids": [],
//	      "labels": [],
//	      "state": "opened",
//	      "blocking_discussions_resolved": true,
//	      "action": "open"
//	    },
//	    "labels": [],
//	    "changes": {
//	      "merge_status": {
//	        "previous": "unchecked",
//	        "current": "preparing"
//	      }
//	    },
//	    "repository": {
//	      "name": "xxx",
//	      "url": "git@tgit.shushu.work:test-group/xxx.git",
//	      "description": "",
//	      "homepage": "http://tgit.shushu.work/test-group/xxx"
//	    }
//	  },
//	  "headers": {
//	    "content-type": "application/json",
//	    "user-agent": "GitLab/14.10.2-ee",
//	    "x-gitlab-event": "System Hook",
//	    "x-gitlab-token": "94b3804b1922834ad505d3fe1699b87a",
//	    "x-gitlab-event-uuid": "aa2984ee-18a9-4127-b18c-8a89d1ae0064",
//	    "connection": "close",
//	    "host": "10.208.200.121:8080",
//	    "content-length": "3432"
//	  },
//	  "origin": "http://10.208.200.121:8080"
//	}
type RequestMergeRequestBody struct {
	ObjectKind string `json:"object_kind"` // merge_request
	EventName  string `json:"event_name"`  // merge_request
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
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
	ObjectAttributes struct {
		AssigneeID     int    `json:"assignee_id"`
		AuthorID       int    `json:"author_id"`
		CreatedAt      string `json:"created_at"`
		HeadPipelineID int    `json:"head_pipeline_id"`
		ID             int    `json:"id"`
		IID            int    `json:"iid"`
		LastEditedAt   string `json:"last_edited_at"`
		LastEditedByID int    `json:"last_edited_by_id"`
		MergeCommitSha string `json:"merge_commit_sha"`
		MergeError     string `json:"merge_error"`
		MergeParams    struct {
			ForceRemoveSourceBranch bool `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string `json:"merge_status"` // preparing, success, failed
		MergeUserID               int    `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool   `json:"merge_when_pipeline_succeeds"`
		MilestoneID               int    `json:"milestone_id"`
		SourceBranch              string `json:"source_branch"`
		SourceProjectID           int    `json:"source_project_id"`
		StateID                   int    `json:"state_id"` // 1: opened, 2: closed, 3: merged
		TargetBranch              string `json:"target_branch"`
		TargetProjectID           int    `json:"target_project_id"`
		Title                     string `json:"title"`
		UpdatedAt                 string `json:"updated_at"`
		UpdatedByID               int    `json:"updated_by_id"`
		URL                       string `json:"url"`
		Source                    struct {
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
		} `json:"source"`
		Target struct {
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
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Title     string    `json:"title"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool   `json:"work_in_progress"`
		TotalTimeSpent      int    `json:"total_time_spent"`
		TimeChange          int    `json:"time_stats_is_outdated"`
		HumanTotalTimeSpent string `json:"human_total_time_spent"`
		HumanTimeEstimate   string `json:"human_time_estimate"`
		AssigneeIDs         []int  `json:"assignee_ids"`
		Labels              []struct {
			Color string `json:"color"`
			Name  string `json:"name"`
			ID    int    `json:"id"`
		} `json:"labels"`
		State                       string `json:"state"` // opened, closed, merged
		BlockingDiscussionsResolved bool   `json:"blocking_discussions_resolved"`
		Action                      string `json:"action"` // open, close, merge, reopen, update, reopen, close, merge, update, assignee, label, milestone, remove_label, destroy
	} `json:"object_attributes"`
	Labels []struct {
		Color string `json:"color"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
	} `json:"labels"`
	Changes struct {
		MergeStatus struct {
			Previous string `json:"previous"`
			Current  string `json:"current"`
		} `json:"merge_status"`
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}
