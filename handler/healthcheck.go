package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type HealthCheckHandler struct {
	Engine *echo.Echo
}

func (h HealthCheckHandler) inject() {
	h.Engine.GET("/ping", h.ping)
}

func (h HealthCheckHandler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
