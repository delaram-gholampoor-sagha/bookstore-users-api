package app

import (
	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// the only place that we are interacting with our application with gingonic package (http server) is in our app and controller package
func StartApplication() {

	mapUrls()

	logger.Info("about to start the application...")

	router.Run(":8082")

}
