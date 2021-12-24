package users

import "encoding/json"

// this is the inforamtion that you are going to retrieve when it is a public request
type PublicUser struct {
	Id int64 `json:"user_id"`
	// FirstName   string `json:"first_name"`
	// LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

// this is the inforamtion that you are going to retrieve when it is an internal request
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

// is this a public or private request ? based on this condition we wanna show different inforamtion of the user
func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			// here we have different keys in json format
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// if we had the same name for id we would do the following approach
	userjson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userjson, &privateUser)
	return privateUser
}
