package jwtauth

type User struct {
	AppName  string `json:"app_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	AppName   string `json:"app_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
