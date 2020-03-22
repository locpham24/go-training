package model

import (
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/form"
)

type Note struct {
	gorm.Model
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	CategoryId int    `json:"category_id"`
}

func (n *Note) Fill(input form.Note) {
	n.Title = input.Title
	n.Completed = input.Completed
	n.CategoryId = input.CategoryId
}
