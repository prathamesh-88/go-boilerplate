package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/prathamesh-88/go-boilerplate/schemas"
	"github.com/prathamesh-88/go-boilerplate/services"
)

type Auth struct{}

func (a *Auth) Login(c *gin.Context) {

	// Parse request in the user format
	var userInput schemas.LoginSchema
	c.BindJSON(&userInput)
	log.Println("User Input:", userInput)
}

func (a *Auth) SignUp(c *gin.Context) {
	// Parse request in the JSON format
	var userInput schemas.SignUpSchema
	c.BindJSON(&userInput)
	// Register user
	log.Println("User Input:", userInput.Email, userInput.Password, userInput.Username)

	createdUser, err := services.RegisterUser(userInput.Username, userInput.Email, userInput.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"createdUser": createdUser,
	})
}

func (a *Auth) SignOut(c *gin.Context) {

}

func (a *Auth) LoginScreen(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
