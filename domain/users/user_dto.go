//user_data  = > data transfer object
// our definition
// so the data transfer object is basically the object that we are going to be transfering from the persistence layer to the application and  backward so if you are working with the user the user is going to be defiend in the dto . because the user is the object that you are going to be moving between the persistence layer and our application

package users

import (
	"strings"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"
)

const (
	StatusActive = "active"
)

// here we have the definition and our dao where we have the access layer to our database
// we have a passoword field but we are not looking at this password when we are working with json
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

// this is a method .... it contains the func keyword , the struct we are assigning this method to the name of the method  , the parameters and at the end what ever we return
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequessrError("Invalid email address")

	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequessrError("Invalid Password")
	}

	return nil
}
