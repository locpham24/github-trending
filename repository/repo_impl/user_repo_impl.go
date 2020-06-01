package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/locpham24/github-trending/c_errors"
	"github.com/locpham24/github-trending/db"
	"github.com/locpham24/github-trending/model"
	"github.com/locpham24/github-trending/model/req"
	repo "github.com/locpham24/github-trending/repository"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repo.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES (:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.DB.NamedExecContext(ctx, statement, user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, c_errors.UserConflict
			}
		}
		return user, c_errors.SignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) CheckSignIn(ctx context.Context, loginReq req.ReqSignIn) (model.User, error) {
	statement := `
		SELECT * FROM users WHERE email=$1
	`
	user := model.User{}
	err := u.sql.DB.GetContext(ctx, &user, statement, loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, c_errors.UserNotFound
		}
		return user, c_errors.SignInFail
	}

	return user, nil
}

func (u *UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error) {
	var user model.User

	err := u.sql.DB.GetContext(context, &user, "SELECT * FROM users WHERE user_id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, c_errors.UserNotFound
		}
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		UPDATE users
		SET 
			full_name = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			updated_at = COALESCE (:updated_at, updated_at)
		WHERE user_id = :user_id
	`
	user.UpdatedAt = time.Now()
	result, err := u.sql.DB.NamedExecContext(context, statement, user)
	if err != nil {
		return user, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return user, c_errors.UserNotUpdate
	}

	return user, nil
}

func (u *UserRepoImpl) Select() ([]model.User, error) {
	users := []model.User{}
	return users, nil
}

func (u *UserRepoImpl) Insert(user model.User) error {
	return nil
}
