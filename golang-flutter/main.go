package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/locpham24/go-training/golang-flutter/repository/repo_impl"

	"github.com/locpham24/go-training/golang-flutter/db"
	"github.com/locpham24/go-training/golang-flutter/handler"
	myLog "github.com/locpham24/go-training/golang-flutter/log"
)

func init() {
	myLog.NewLogger()
}
func main() {
	e := echo.New()
	if err := godotenv.Load(); err != nil {
		log.Error(err.Error())
	}

	DB := &db.Sql{}
	DB.Connect()
	defer DB.Close()

	e.Use(middleware.AddTrailingSlash())
	myLog.Info("hahahahahaha")

	userRepoImpl := repo_impl.NewUserRepo(DB)
	handler.InitRouter(e, &userRepoImpl)
	e.Logger.Fatal(e.Start(":8080"))
}
