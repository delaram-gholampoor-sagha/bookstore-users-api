//user_data =>  data  access object
// our access layer to our database
//  important //the only point in our entire applicaion where you work with the database is indeed over here
// here we are going to have the entire logic to persist and to retrieve this user from the database

package users

import (
	"fmt"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// when we say get we get a user by its id (primary key)
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequessrError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequessrError(fmt.Sprintf("user %d already exists", user.Id))

	}

	usersDB[user.Id] = user
	return nil
}
