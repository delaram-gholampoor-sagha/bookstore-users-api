package services

import (
	"net/http"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
