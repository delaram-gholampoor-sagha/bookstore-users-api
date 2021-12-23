package app

import (
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/controllers/ping"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/controllers/users"
)

func MapUrls() {
	// we are defining the function that needs to be executed against this path
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	router.POST("/users", users.CreateUser)

	router.PUT("/users/:user_id", users.UpdateUser)
}
