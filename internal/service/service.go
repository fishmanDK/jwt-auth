package service

import (
	jwtauth "github.com/fishmanDK"
	postgresql "github.com/fishmanDK/internal/repository/postgreSQL"
)

type Tokens struct {
	Access_token  string
	Refresh_token string
}

type Validate interface {
}

type Authentication interface{
	Authentication(user jwtauth.User) (Tokens, error)
}

type JWT interface {
	CreateAccessToken(id int64) (string, error)
	CreateRefreshToken() (string, error)
}

type Service struct {
	Validate
	Authentication

}

func NewService(repo *postgresql.PostgreDB) *Service {
	return &Service{
		Validate: NewValidateService(repo),
		Authentication: NewAuthService(repo),
	}
}
