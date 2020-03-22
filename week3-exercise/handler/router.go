package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(db *gorm.DB, r *gin.Engine) {
	noteHandler := NoteHandler{
		DB:     db,
		Engine: r,
	}
	noteHandler.inject()
}
