package handler

import (
	"github.com/labstack/echo"
	repo "github.com/locpham24/github-trending/repository"
)

func InitRouter(e *echo.Echo, repo *repo.Repos) {
	healthCheckService := HealthCheckHandler{
		Engine: e,
	}
	healthCheckService.inject()

	userService := UserHandler{
		Engine:   e,
		UserRepo: *repo.UserRepo,
	}
	userService.inject()
}
