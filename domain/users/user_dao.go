//user_data =>  data  access object
// our access layer to our database
//  important //the only point in our entire applicaion where you work with the database is indeed over here
// here we are going to have the entire logic to persist and to retrieve this user from the database

package users

import (
	"fmt"
	"strings"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/datasources/mysql/users_db"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/date_utils"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
)

var (
	usersDB = make(map[int64]*User)
)

// when we say get we get a user by its id (primary key)
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	// first we want to check if the query is valid
	//better performance
	stmt, err := users_db.Client.Prepare("INSERT INTO users(`first_name` , last_name , email , date_created) VALUES (?, ?, ?, ?);")
	if err != nil {
		return errors.NewIntervalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequessrError(fmt.Sprintf("email %s already exists", user.Email))
		}
		errors.NewIntervalServerError(fmt.Sprintf("error when trying to save user : %s ", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewIntervalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}

	user.Id = userId
	return nil
}
