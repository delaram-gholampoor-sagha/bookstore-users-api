package services

import (
	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/crypto_utils"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(userId int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(userId int64) rest_errors.RestErr
	SearchUser(status string) (users.Users, rest_errors.RestErr)
	LogInUser(request users.LogInRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

// if your function needs to return an error , it needs to be at the end
func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	// take the current user that exist
	// in both cases partial and not partial we need the current user
	current := &users.User{Id: user.Id}
	// if we dont have any user return nil
	if err := current.Get(); err != nil {
		return nil, err
	}

	// if we have a user validate it
	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if user.FirstName == "" {
			current.FirstName = user.FirstName
		}
		if user.LastName == "" {
			current.LastName = user.LastName
		}
		if user.Email == "" {
			current.Email = user.Email
		}
		// if its not partial modify every field
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

}

// what are the possible results that you might get from deleting a user ? probably just an error

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindUserByStatus(status)
}

func (s *usersService) LogInUser(request users.LogInRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}

	if err := dao.FIndByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
