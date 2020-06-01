package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	myMiddleware "github.com/locpham24/github-trending/middleware"
	"github.com/locpham24/github-trending/model"
	req2 "github.com/locpham24/github-trending/model/req"
	repo "github.com/locpham24/github-trending/repository"
	"net/http"
)

type GithubHandler struct {
	Engine     *echo.Echo
	GithubRepo repo.GithubRepo
}

func (h GithubHandler) inject() {
	github := h.Engine.Group("/github", myMiddleware.VerifyToken())
	github.GET("/trending", h.trending)

	bookmark := h.Engine.Group("/bookmark", myMiddleware.VerifyToken())
	bookmark.GET("/list", h.getBookmarks)
	bookmark.POST("/add", h.addBookmark)
	bookmark.DELETE("/delete", h.deleteBookmark)
}

func (h GithubHandler) trending(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	repos, _ := h.GithubRepo.SelectRepos(c.Request().Context(), claims.UserId, 25)

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get github trending success",
		Data:       repos,
	})
}

func (h GithubHandler) getBookmarks(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	repos, _ := h.GithubRepo.GetBookmarks(c.Request().Context(), claims.UserId)

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get bookmarks success",
		Data:       repos,
	})
}

func (h GithubHandler) addBookmark(c echo.Context) error {
	req := req2.ReqBookmark{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	bId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = h.GithubRepo.AddBookmark(c.Request().Context(), bId.String(), claims.UserId, req.RepoName)

	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Add bookmark success",
		Data:       nil,
	})
}

func (h GithubHandler) deleteBookmark(c echo.Context) error {
	req := req2.ReqBookmark{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	err := h.GithubRepo.DeleteBookmark(c.Request().Context(), claims.UserId, req.RepoName)
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete bookmark success",
		Data:       nil,
	})
}
