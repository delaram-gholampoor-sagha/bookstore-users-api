package app

import (
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/controllers/ping"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/controllers/users"
)

func MapUrls() {

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	router.POST("/users", users.CreateUser)
}
