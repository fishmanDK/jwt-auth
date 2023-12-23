package service

import (
	"github.com/fishmanDK/internal/repository"
	"time"

	jwtauth "github.com/fishmanDK"
	postgresql "github.com/fishmanDK/internal/repository/postgreSQL"
)

type AuthService struct {
	repo *repository.Storage
}

func NewAuthService(db *repository.Storage) *AuthService {
	return &AuthService{
		repo: db,
	}
}

func (s *AuthService) Authentication(user jwtauth.User) (Tokens, error) {
	var (
		tokens Tokens
	)

	res, err := s.repo.Authentication(user)

	tokens.Access_token, err = CreateAccessToken(res.Id, res.UserName, res.Role)
	if err != nil {
		return tokens, err
	}

	tokens.Refresh_token, err = CreateRefreshToken()
	if err != nil {
		return tokens, err
	}

	session := postgresql.Session{
		Refresh_token: tokens.Refresh_token,
		ExpiresAt:     time.Now().Add(refresh_tokenTtl).UTC(),
	}

	err = s.repo.CreateSession(res.Id, session)
	return tokens, err
}

func (s *AuthService) CreateUser(newUser jwtauth.CreateUser) (int64, error) {
	return s.repo.CreateUser(newUser)
}

//func (s *AuthService) ParceToken() {
//
//}
