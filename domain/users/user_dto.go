//user_data transfer object object
// our definition

package users

import (
	"strings"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequessrError("invalid email address")

	}

	return nil
}
