package handler

import (
	"github.com/labstack/echo"
	repo "github.com/locpham24/go-training/golang-flutter/repository"
)

func InitRouter(e *echo.Echo, userRepo *repo.UserRepo) {
	healthCheckService := HealthCheckHandler{
		Engine: e,
	}
	healthCheckService.inject()

	userService := UserHandler{
		Engine:   e,
		UserRepo: *userRepo,
	}
	userService.inject()
}
