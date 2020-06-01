package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/locpham24/github-trending/db"
	"github.com/locpham24/github-trending/handler"
	"github.com/locpham24/github-trending/helper"
	repo "github.com/locpham24/github-trending/repository"
	"github.com/locpham24/github-trending/repository/repo_impl"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "7000"
	}

	if err := godotenv.Load(); err != nil {
		log.Error(err.Error())
	}

	DB := &db.Sql{}
	DB.Connect()
	defer DB.Close()

	e := echo.New()

	e.Use(middleware.AddTrailingSlash())

	userRepoImpl := repo_impl.NewUserRepo(DB)
	githubRepoImpl := repo_impl.NewGithubRepo(DB)

	repoImpl := repo.Repos{
		UserRepo:   &userRepoImpl,
		GithubRepo: &githubRepoImpl,
	}
	handler.InitRouter(e, &repoImpl)

	helper.CrawlRepo(githubRepoImpl)

	e.Logger.Fatal(e.Start(":" + port))
}
