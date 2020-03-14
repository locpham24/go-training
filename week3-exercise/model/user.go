package model

import (
	"errors"
	"log"

	"github.com/locpham24/go-training/week3-exercise/auth"
)

type User struct {
	Id   int
	Name string `json:"name"`
}

func ValidateUser(username interface{}, password interface{}) (string, error) {
	var err error
	if username == "admin" && password == "admin" {
		//Create token
		token, err := auth.CreateJwtToken("admin", "1")
		if err != nil {
			log.Println("Error Creating Jwt Token:", err)
			return "", err
		}
		return token, nil
	}

	if username == "chris" && password == "123456" {
		//Create token
		token, err := auth.CreateJwtToken("chris", "2")
		if err != nil {
			log.Println("Error Creating Jwt Token:", err)
			return "", err
		}
		return token, nil
	}
	err = errors.New("Wrong username or password")
	return "", err
}
