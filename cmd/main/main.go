
package main

import (
 "github.com/gin-gonic/gin"
 "gorm.io/gorm"
 "log"
 "tasks/config"
 "tasks/internal/routes"
)

func main() {
 db, err := config.InitDB()
 if err != nil {
 log.Fatal(err)
 }
 r := gin.Default()
 routes.SetupRoutes(r, db)
 r.Run(":8080")
}

