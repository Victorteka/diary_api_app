package main

import (
	"diary_api/controller"
	"diary_api/middleware"
	"diary_api/model"
	"diary_api/util/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entries", controller.AddEntry)
	protectedRoutes.GET("/entries/:id", controller.GetEntry)
	protectedRoutes.GET("/entries", controller.GetAllEntries)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
