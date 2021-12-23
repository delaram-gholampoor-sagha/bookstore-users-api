//user_data =>  data  access object
// our access layer to our database
//  important //the only point in our entire applicaion where you work with the database is indeed over here
// here we are going to have the entire logic to persist and to retrieve this user from the database

// by using the dqo file we separete the bussiness logic from the persistance layer

package users

import (
	"fmt"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/datasources/mysql/users_db"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/date_utils"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	queryInsertUser       = "INSERT INTO users(first_name , last_name , email , date_created) VALUES(?, ?, ?, ?);"
	GetUser               = "SELECT id , first_name , last_name , email , date_created  FROM users WHERE id = ? ;"
	queryUpdateUser       = "UPDATE users SET first_name = ? , last_name = ? , email = ? WHERE id = ? ;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ? ;"
	queryFindUserByStatus = "SELECT id , fisrt_name , last_name , email , date_created , status FROM users WHERE status=? ;"
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

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
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
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
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

// => /internal/users/search?status=active
func (user *User) FindUserByStatus(status string) ([]User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewIntervalServerError(err.Error())
	}
	defer stmt.Close()

	// if we have any error we are not going to look at  any parameters on the left
	rows, err := stmt.Query(queryFindUserByStatus)

	if err != nil {
		return nil, errors.NewIntervalServerError(err.Error())
	}
	// we always defer once we know we have a valid result
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		// we always have to pass a pointer to the scan function
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
