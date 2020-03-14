package api

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/model"
)

func GetAllNote(c *gin.Context) {
	notes := []model.Note{}
	db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer db.Close()

	db.Find(&notes).Limit(5)
	c.JSON(200, gin.H{
		"data": notes,
	})
}

func GetOneNote(c *gin.Context) {
	note := model.Note{}
	noteId := c.Param("id")
	db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	defer db.Close()

	db.Where("id = ?", noteId).Find(&note)

	c.JSON(200, gin.H{
		"data": note,
	})
}

func CreateNote(c *gin.Context) {
	form := model.Note{}
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	catId, _ := strconv.Atoi(c.Query("category_id"))
	fmt.Println("db_uri:", os.Getenv("DB_URI"))
	db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	defer db.Close()
	db.DropTableIfExists(&form)
	db.AutoMigrate(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	form.CreatedAt = time.Now()
	form.CategoryId = catId

	db.Create(&form)
	c.JSON(200, form)
}
