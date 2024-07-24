package main

import (
    "login-signup-api/handlers"
    "login-signup-api/middleware"
    _ "login-signup-api/docs"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
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

    // Public routes
    r.POST("/signup", handlers.SignUp)
    r.POST("/login", handlers.Login)

    // Protected routes
    protected := r.Group("/")
    protected.Use(middleware.Auth())
    {
        // Define protected routes here
        // Example: protected.GET("/profile", handlers.Profile)
    }

    r.Run(":8080")
}
