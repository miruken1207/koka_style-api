package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	database "koka_style/database"
	handler "koka_style/handlers"

	_ "koka_style/docs"
)

func main() {

	db, err := database.Init()
	if err != nil {
		panic("Failed to connect to the database")
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", handler.Root(db))
	router.POST("/register", handler.Register(db))
	router.POST("/login", handler.Login(db))

	router.Run(":8080")
}
