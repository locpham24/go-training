package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/locpham24/go-training/week3-exercise/model"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	form := LoginForm{}
	err := c.BindJSON(&form)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	username := form.Username
	password := form.Password

	fmt.Println("usename: ", username)
	fmt.Println("password: ", password)
	token, err := model.ValidateUser(username, password)
	fmt.Println("err: ", err)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token":   token,
		"message": "You were log in!",
	})
}
