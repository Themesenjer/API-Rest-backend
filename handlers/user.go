package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// User represents a user in the system
type User struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
    Name     string `json:"name"`
    Surname  string `json:"surname"`
}

// SignUp handles user registration
// @Summary Sign up a new user
// @Description Register a new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param   user  body      User  true  "User data"
// @Success 201  {object}  User
// @Failure 400  {object}  ErrorResponse
// @Router /signup [post]
func SignUp(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

// Login handles user login
// @Summary Log in a user
// @Description Authenticate a user
// @Tags User
// @Accept  json
// @Produce  json
// @Param   user  body      User  true  "User data"
// @Success 200  {object}  User
// @Failure 401  {object}  ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
        return
    }
    // Simulate successful login
    c.JSON(http.StatusOK, user)
}
