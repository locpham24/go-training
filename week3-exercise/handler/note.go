package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/form"
	"github.com/locpham24/go-training/week3-exercise/model"
	"strconv"
)

type NoteHandler struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

func (h *NoteHandler) inject() {
	h.Engine.GET("/note", h.list)
	h.Engine.POST("/note", h.create)
}
func (h *NoteHandler) list(c *gin.Context) {
	notes := []model.Note{}
	err := h.DB.
		Limit(10).
		Offset(0).
		Find(&notes).
		Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, notes)
	return
}

func (h NoteHandler) delete(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw) //Ascii
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	// Delete
	err = h.DB.Delete(model.Note{}, "id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Deleted is success",
	})
}

func (h NoteHandler) get(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw) //Ascii
	if err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}
	note := model.Note{}
	err = h.DB.First(&note, id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, note)
}

func (h NoteHandler) create(c *gin.Context) {
	input := form.Note{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	note := model.Note{}
	note.Fill(input)

	err := h.DB.Create(&note).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":    2000,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, note)
}
