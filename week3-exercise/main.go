package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/locpham24/go-training/week3-exercise/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := route.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
