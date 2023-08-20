package cache

import "github.com/xanzy/go-gitlab"

type GitLabCache struct {
	client    *gitlab.Client
	projectId int
}
