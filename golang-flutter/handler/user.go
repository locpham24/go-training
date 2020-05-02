package handler

import (
	"github.com/dgrijalva/jwt-go"
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/locpham24/go-training/golang-flutter/c_errors"
	"github.com/locpham24/go-training/golang-flutter/log"
	myMiddleware "github.com/locpham24/go-training/golang-flutter/middleware"
	"github.com/locpham24/go-training/golang-flutter/model"
	req2 "github.com/locpham24/go-training/golang-flutter/model/req"
	repo "github.com/locpham24/go-training/golang-flutter/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserHandler struct {
	Engine   *echo.Echo
	UserRepo repo.UserRepo
}

func (h UserHandler) inject() {
	h.Engine.POST("/sign-in", h.signIn)
	h.Engine.POST("/sign-up", h.signUp)
	h.Engine.GET("/profile", h.myProfile, myMiddleware.VerifyToken())
}

func (h UserHandler) signIn(c echo.Context) error {
	req := req2.ReqSignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := h.UserRepo.CheckSignIn(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    c_errors.SignInFail.Error(),
			Data:       nil,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    c_errors.Incorrect.Error(),
			Data:       nil,
		})
	}

	// gen token
	token, err := model.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Token = token
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       user,
	})
}

func (h UserHandler) signUp(c echo.Context) error {
	req := req2.ReqSignUp{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	hashedPassoword := string(bytesPassword)

	role := model.Member.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId:    userId.String(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hashedPassoword,
		Role:      role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
	}

	user, err = h.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    c_errors.UserConflict.Error(),
			Data:       nil,
		})
	}
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       user,
	})
}

func (h UserHandler) myProfile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user, err := h.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		if err == c_errors.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    c_errors.UserNotFound.Error(),
				Data:       nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       user,
	})
}
