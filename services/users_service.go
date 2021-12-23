package services

import (
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/date_utils"
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

	user.DateCreated = date_utils.GetNowDBFormat()

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	// take the current user that exist
	// in both cases partial and not partial we need the current user
	current, err := GetUser(user.Id)
	// if we dont have any user return nil
	if err != nil {
		return nil, err
	}

	// if we have a user validate it
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName == "" {
			current.FirstName = user.FirstName
		}
		if user.LastName == "" {
			current.LastName = user.LastName
		}
		if user.Email == "" {
			current.Email = user.Email
		}
		// if its not partial modify every field
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := user.Update(); err != nil {
		return nil, err
	}

	return current, nil

}

// what are the possible results that you might get from deleting a user ? probably just an error

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindUserByStatus(status)
}
