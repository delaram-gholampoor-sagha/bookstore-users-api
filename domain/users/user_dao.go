//user_data =>  data  access object
// our access layer to our database
//  important //the only point in our entire applicaion where you work with the database is indeed over here
// here we are going to have the entire logic to persist and to retrieve this user from the database

package users

import (
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/datasources/mysql/users_db"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/date_utils"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name , last_name , email , date_created) VALUES(?, ?, ?, ?);"
	GetUser          = "SELECT id , first_name , last_name , email , date_created  FROM users WHERE id = ? ;"
	queryUpdateuser  = "UPDATE users SET first_name = ? , last_name = ? , email = ? WHERE id = ? ;"
	queryDeleteuser  = "DELETE FROM users WHERE id = ? ;"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {

		return mysql_utils.ParseError(getErr)

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

	insertResult, saveErr := stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateuser)
	if err != nil {
		return errors.NewIntervalServerError(err.Error())
	}

	defer stmt.Close()

	// atemp to execute this sentence with the given values
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteuser)
	if err != nil {
		return errors.NewIntervalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
