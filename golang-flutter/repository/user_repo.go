package repository

import (
	"context"
	"github.com/locpham24/go-training/golang-flutter/model"
	"github.com/locpham24/go-training/golang-flutter/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckSignIn(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	Insert(u model.User) error
}
