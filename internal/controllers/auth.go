




package controllers

import (
 "github.com/gin-gonic/gin"
 "tasks/internal/models"
)

func Signup(c *gin.Context) {
 var user models.User
 c.BindJSON(&user)
 // Implement signup logic
}

func Login(c *gin.Context) {
 var user models.User
 c.BindJSON(&user)
 // Implement login logic
}




