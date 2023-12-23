package jwtauth

type User struct {
	AppName  string `json:"app_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	AppId     int64  `json:"app_id"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
