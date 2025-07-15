package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	config "koka_style/config"
	database "koka_style/database"
	handler "koka_style/handlers"
	middleware "koka_style/middlewares"

	_ "koka_style/docs"
)

func main() {

	config.Init()

	db, err := database.Init()
	if err != nil {
		panic("Failed to connect to the database")
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", handler.Root(db))

	router.POST("/sign_up", handler.SignUp(db))
	router.POST("/login", handler.Login(db))

	router.GET("/products", middleware.AuthMiddleware(), handler.GetProducts(db))

	router.Run(":8080")
}
