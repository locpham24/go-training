package repository

import (
	"context"
	"github.com/locpham24/go-training/model"
	"github.com/locpham24/go-training/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckSignIn(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	Select() ([]model.User, error)
	Insert(u model.User) error
}
