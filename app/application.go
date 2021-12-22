package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

// the only place that we are interacting with our application with gingonic package (http server) is in our app and controller package
func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MapUrls()
	router.Run(":8080")

}
