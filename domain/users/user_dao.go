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
	queryInsertUser  = "INSERT INTO users(first_name , last_name , email , date_created) VALUES(?, ?, ?, ?);"
	GetUser          = "SELECT id , first_name , last_name , email , date_created  FROM users WHERE id = ?"
	errorNoRows      = "no rows in result set"
)

// when we say get we get a user by its id (primary key)
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(GetUser)
	if err != nil {
		return errors.NewIntervalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow()
	// take what ever you have as an id in the database and use that value to populate these fields
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found  ", user.Id))
		}
		return errors.NewIntervalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))

	}

	return nil
}
func (user *User) Save() *errors.RestErr {
	// first we want to check if the query is valid
	//better performance
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewIntervalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequessrError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewIntervalServerError(fmt.Sprintf("error when trying to save user : %s ", err.Error()))

	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewIntervalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}

	user.Id = userId
	return nil
}
