package repository

import (
	"context"
	"github.com/locpham24/github-trending/model"
	"github.com/locpham24/github-trending/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckSignIn(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
	Insert(u model.User) error
}
