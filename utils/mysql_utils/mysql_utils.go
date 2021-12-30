package mysql_utils

import (
	"errors"
	"strings"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	// if we werent able to convert the error we are getting from this function to the mysql error ...
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no recors matching the given id ")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}

	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))

}
