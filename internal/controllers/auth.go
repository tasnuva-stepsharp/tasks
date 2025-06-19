




package controllers

import (
 "github.com/gin-gonic/gin"
 "tasks/internal/models"
)

func Signup(c *gin.Context) {
 var user models.User
 c.BindJSON(&user)
 // Implement signup logic


var user models.User
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
var existingUser models.User
if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
	return
}
if err := db.Create(&user).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
c.JSON(http.StatusCreated, user)


}

func Login(c *gin.Context) {
 var user models.User
 c.BindJSON(&user)


if err := c.ShouldBindJSON(&loginUser); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
var existingUser models.User
if err := db.Where("username = ?", loginUser.Username).First(&existingUser).Error; err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	return
}
if existingUser.Password != loginUser.Password {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	return
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
"user_id": existingUser.ID,
"exp": time.Now().Add(time.Hour *72).Unix(),
})
tokenString, err := token.SignedString([]byte("secret_key"))
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
c.JSON(http.StatusOK, gin.H{"token": tokenString})


 // Implement login logic
}




