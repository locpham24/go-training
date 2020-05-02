package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/locpham24/go-training/golang-flutter/model"
	"os"
)

func VerifyToken() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JwtSecretKey")),
		Claims:     &model.JwtCustomClaims{},
	}
	return middleware.JWTWithConfig(config)
}
