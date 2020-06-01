package repository

import (
	"context"
	"github.com/locpham24/github-trending/model"
)

type GithubRepo interface {
	SaveRepo(context context.Context, repo model.Repo) (model.Repo, error)
	SelectRepoByName(context context.Context, name string) (model.Repo, error)
	SelectRepos(context context.Context, userId string, limit int) ([]model.Repo, error)
	UpdateRepo(context context.Context, repo model.Repo) (model.Repo, error)

	GetBookmarks(context context.Context, userId string) ([]model.Repo, error)
	AddBookmark(context context.Context, bId string, userId string, repoName string) error
	DeleteBookmark(context context.Context, userId string, repoName string) error
}
