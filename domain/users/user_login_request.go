package users

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"passowrd"`
}
