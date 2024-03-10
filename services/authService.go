package services

import (
	"log"

	"github.com/prathamesh-88/go-boilerplate/models"
)

func RegisterUser(username string, email string, password string) (models.User, error) {
	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     "admin",
	}

	createStatus, err := models.CreateUser(&user)

	if err != nil {
		log.Println("Error in registering User: ", err)
		return models.User{}, err
	}

	return createStatus, nil
}
