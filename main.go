package main

import (
	"github.com/gin-gonic/gin"

	database "koka_style/database"
	handler "koka_style/handlers"
)

func main() {

	db, err := database.Init()
	if err != nil {
		panic("Failed to connect to the database")
	}
	
	router := gin.Default()

	router.GET("/", handler.Root(db))
	router.POST("/register", handler.Register(db))
	router.POST("/login", handler.Login(db))

	router.Run(":8080")
}
