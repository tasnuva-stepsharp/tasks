







package routes

import (
 "github.com/gin-gonic/gin"
 "gorm.io/gorm"
 "tasks/internal/controllers"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
 authGroup := r.Group("/auth")
 {
 authGroup.POST("/signup", controllers.Signup)
 authGroup.POST("/login", controllers.Login)
 }
 taskGroup := r.Group("/tasks")
 {
 taskGroup.POST("", controllers.CreateTask)
 taskGroup.GET("", controllers.GetTasks)
 taskGroup.GET("/:id", controllers.GetTask)
 taskGroup.PUT("/:id", controllers.UpdateTask)
 taskGroup.DELETE("/:id", controllers.DeleteTask)
 }
}






