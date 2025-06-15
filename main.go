package main

import (
	"log"
	"email-service/internal/config"
	"email-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sqlDB from gorm DB:", err)
	}
	defer sqlDB.Close()

	r := gin.Default()

	r.POST("/password/recover", handlers.RecoverPasswordHandler)

	if err := r.Run(":4000"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
