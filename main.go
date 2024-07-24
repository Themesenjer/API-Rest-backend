package main

import (
	_ "login-signup-api/docs"
	"login-signup-api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Login Signup API
// @version 1.0
// @description This is a sample server for a user login/signup API.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	// Swagger setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
