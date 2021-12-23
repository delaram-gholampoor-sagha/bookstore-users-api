package services

import (
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// if your function needs to return an error , it needs to be at the end
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	// take the current user that exist
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := user.Update(); err != nil {
		return nil, err
	}

	return current, nil

}
