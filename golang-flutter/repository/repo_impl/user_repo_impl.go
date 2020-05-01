package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/locpham24/go-training/c_errors"
	"github.com/locpham24/go-training/db"
	"github.com/locpham24/go-training/log"
	"github.com/locpham24/go-training/model"
	"github.com/locpham24/go-training/model/req"
	repo "github.com/locpham24/go-training/repository"
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
		log.Error(err.Error())
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
		log.Error(err.Error())
		return user, c_errors.SignInFail
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
