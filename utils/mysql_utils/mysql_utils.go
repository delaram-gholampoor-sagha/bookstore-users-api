package mysql_utils

import (
	"strings"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	// if we werent able to convert the error we are getting from this function to the mysql error ...
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no recors matching the given id ")
		}
		return errors.NewIntervalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequessrError("invalid data")
	}

	return errors.NewIntervalServerError("error processing request")

}
