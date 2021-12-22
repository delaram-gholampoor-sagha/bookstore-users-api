package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// the only place that we are interacting with our application with gingonic package (http server) is in our app and controller package
func StartApplication() {
	MapUrls()
	router.Run(":8080")

}
