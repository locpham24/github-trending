package c_errors

import "errors"

var (
	UserConflict  = errors.New("the email is existed")
	SignUpFail    = errors.New("sign up fail")
	SignInFail    = errors.New("sign in fail")
	UserNotFound  = errors.New("can not found this user")
	Incorrect     = errors.New("email or password is incorrect")
	TokenError    = errors.New("token error")
	UserNotUpdate = errors.New("user not update")
	AddRepoFail   = errors.New("add repo fail")
	RepoNotFound  = errors.New("repo not found")
)
