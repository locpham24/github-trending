package repo_impl

import (
	"context"
	"database/sql"
	"github.com/locpham24/github-trending/c_errors"
	"github.com/locpham24/github-trending/db"
	"github.com/locpham24/github-trending/model"
	repo "github.com/locpham24/github-trending/repository"
	"time"
)

type GithubRepoImpl struct {
	sql *db.Sql
}

func NewGithubRepo(sql *db.Sql) repo.GithubRepo {
	return &GithubRepoImpl{
		sql: sql,
	}
}

func (g *GithubRepoImpl) SaveRepo(ctx context.Context, repo model.Repo) (model.Repo, error) {
	statement := `
		INSERT INTO repos(name, description, url, color, lang, fork, stars, stars_today, author, created_at, updated_at)
		VALUES (:name, :description, :url, :color, :lang, :fork, :stars, :stars_today, :author, :created_at, :updated_at)
	`
	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()

	_, err := g.sql.DB.NamedExecContext(ctx, statement, repo)
	if err != nil {
		return repo, c_errors.AddRepoFail
	}

	return repo, nil
}

func (g *GithubRepoImpl) SelectRepoByName(ctx context.Context, name string) (model.Repo, error) {
	var githubRepo model.Repo
	err := g.sql.DB.GetContext(ctx, &githubRepo, `SELECT * FROM repos WHERE name = $1`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return githubRepo, c_errors.RepoNotFound
		}
		return githubRepo, err
	}
	return githubRepo, nil
}

func (g *GithubRepoImpl) UpdateRepo(ctx context.Context, repo model.Repo) (model.Repo, error) {
	statement := `
		UPDATE repos
		SET 
			stars  = :stars,
			fork = :fork,
			stars_today = :stars_today,
			author = :author,
			updated_at = :updated_at
		WHERE name = :name
	`
	repo.UpdatedAt = time.Now()
	result, err := g.sql.DB.NamedExecContext(ctx, statement, repo)
	if err != nil {
		return repo, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return repo, c_errors.UserNotUpdate
	}

	return repo, nil
}

func (g *GithubRepoImpl) SelectRepos(context context.Context, userId string, limit int) ([]model.Repo, error) {
	var repos []model.Repo
	err := g.sql.DB.SelectContext(context, &repos,
		`
			SELECT repos.*, COALESCE (bookmarks.repo_name IS NOT NULL, FALSE) as bookmarked
			FROM repos
			LEFT JOIN bookmarks
			ON repos.name = bookmarks.repo_name AND bookmarks.user_id = $1 
			ORDER BY repos.updated_at ASC LIMIT $2
		`, userId, limit)
	if err != nil {
		return repos, err
	}

	return repos, nil
}

func (g *GithubRepoImpl) GetBookmarks(context context.Context, userId string) ([]model.Repo, error) {
	var repos []model.Repo
	err := g.sql.DB.SelectContext(context, &repos,
		`
			SELECT repos.*
			FROM bookmarks
			INNER JOIN repos
			ON bookmarks.user_id = $1 AND bookmarks.repo_name = repos.name
		`, userId)

	if err != nil {
		return repos, err
	}

	return repos, nil
}

func (g *GithubRepoImpl) AddBookmark(ctx context.Context, bId string, userId string, repoName string) error {
	statement := `
		INSERT INTO bookmarks(bid, user_id, repo_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	now := time.Now()
	_, err := g.sql.DB.ExecContext(ctx, statement, bId, userId, repoName, now, now)
	if err != nil {
		return err
	}

	return nil
}

func (g *GithubRepoImpl) DeleteBookmark(ctx context.Context, userId string, repoName string) error {
	statement := `
		DELETE FROM bookmarks
		WHERE user_id = $1 AND repo_name = $2
	`

	_, err := g.sql.DB.ExecContext(ctx, statement, userId, repoName)
	if err != nil {
		return err
	}

	return nil
}
