package route

import (
	"github.com/gin-gonic/gin"
	"github.com/locpham24/go-training/week3-exercise/api"
)

func Init() *gin.Engine {
	r := gin.Default()
	v1Note := r.Group("api/v1/note")
	v1Note.GET("/", api.GetAllNote)
	v1Note.GET("/:id", api.GetOneNote)
	v1Note.POST("/", api.CreateNote)

	v1 := r.Group("api/v1/")

	v1.POST("/login", api.Login)
	return r
}
