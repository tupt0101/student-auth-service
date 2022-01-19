package users

type User struct {
	UserId    int64  `json:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedOn string `json:"created_on"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
