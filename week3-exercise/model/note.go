package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	CategoryId int    `json:"category_id"`
}
