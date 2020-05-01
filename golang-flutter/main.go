package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/locpham24/go-training/repository/repo_impl"

	"github.com/locpham24/go-training/db"
	"github.com/locpham24/go-training/handler"
	myLog "github.com/locpham24/go-training/log"
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

	myLog.Info("hahahahahaha")

	userRepoImpl := repo_impl.NewUserRepo(DB)
	handler.InitRouter(e, &userRepoImpl)
	e.Logger.Fatal(e.Start(":8080"))
}
