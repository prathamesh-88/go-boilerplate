package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prathamesh-88/go-boilerplate/handlers"
)

func BindAuthRoutes(r *gin.Engine) {
	authHandlers := handlers.Auth{}

	rg := r.Group("/auth")
	{
		// Display views
		rg.GET("/login", authHandlers.LoginScreen)

		// POST Handlers
		rg.POST("/signup", authHandlers.SignUp)
		rg.POST("/login", authHandlers.Login)
		rg.POST("/logout", authHandlers.SignOut)
	}
}
