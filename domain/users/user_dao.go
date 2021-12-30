//user_data =>  data  access object
// our access layer to our database
//  important //the only point in our entire applicaion where you work with the database is indeed over here
// here we are going to have the entire logic to persist and to retrieve this user from the database

// by using the dao file we separete the bussiness logic from the persistance layer

package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/datasources/mysql/users_db"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/logger"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail            = "email_UNIQUE"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	GetUser                     = "SELECT id , first_name , last_name , email , date_created , status FROM users WHERE id = ? ;"
	queryUpdateUser             = "UPDATE users SET first_name = ? , last_name = ? , email = ? WHERE id = ? ;"
	queryDeleteUser             = "DELETE FROM users WHERE id = ? ;"
	queryFindByStatus           = "SELECT id , fisrt_name , last_name , email , date_created , status FROM users WHERE status=? ;"
	queryFindByEmailAndPassword = "SELECT id , first_name , last_name , email , date_created , status FROM users WHERE email = ? AND password = ? AND status = ?;"
)

// when we say get we get a user by its id (primary key)
func (user *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(GetUser)
	if err != nil {
		// this what is returned into my system
		logger.Error("error when trying to get user statement", err)
		// this is what we return for the client
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("databse error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	// take what ever you have as an id in the database and use that value to populate these fields
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))

	}

	return nil
}

func (user *User) Save() rest_errors.RestErr {
	// first we want to check if the query is valid
	//better performance
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}

	defer stmt.Close()

	// atemp to execute this sentence with the given values
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user ", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}

	return nil

}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement ", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete user ", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	return nil
}

// => /internal/users/search?status=active
func (user *User) FindUserByStatus(status string) ([]User, rest_errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement ", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))

	}
	defer stmt.Close()

	// if we have any error we are not going to look at  any parameters on the left
	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to find users by status ", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))

	}
	// we always defer once we know we have a valid result
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		// we always have to pass a pointer to the scan function
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct ", err)
			return nil, rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

func (user *User) FIndByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		// this what is returned into my system
		logger.Error("error when trying to get user by email and password statement", err)
		// this is what we return for the client
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	// take what ever you have as an id in the database and use that value to populate these fields
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}

	return nil
}
