package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/handler"
)

type APIService struct {
	db *gorm.DB
}

func NewAPIService(db *gorm.DB) APIService {
	return APIService{
		db: db,
	}
}

func (a *APIService) Start() {
	r := gin.Default()
	handler.InitRouter(a.db, r)
	r.Run()
}
