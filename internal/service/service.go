package service

import (
	jwtauth "github.com/fishmanDK"
	"github.com/fishmanDK/internal/repository"
)

type Tokens struct {
	Access_token  string
	Refresh_token string
}

type Authentication interface {
	Authentication(user jwtauth.User) (Tokens, error)
	CreateUser(newUser jwtauth.CreateUser) (int64, error)
}

type JWT interface {
	CreateAccessToken(id int64) (string, error)
	CreateRefreshToken() (string, error)
}

type Service struct {
	Authentication
}

func NewService(repo *repository.Storage) *Service {
	return &Service{
		Authentication: NewAuthService(repo),
	}
}
