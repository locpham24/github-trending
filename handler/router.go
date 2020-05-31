package handler

import (
	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	healthCheckService := HealthCheckHandler{
		Engine: e,
	}
	healthCheckService.inject()
}
