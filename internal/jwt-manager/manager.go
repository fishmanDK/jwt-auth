package jwtmanager

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTtl = time.Hour * 12
)

type TokenManager interface{
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parce(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct{
	signinKey string
}

func NewManager(signinKey string) (*Manager, error){
	if signinKey == ""{ return nil, errors.New("error: absent signinKey")}
	return &Manager{signinKey: signinKey}, nil
}

func(m *Manager) NewJWT(userId string, ttl time.Duration) (string, error){
	acssesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTtl).Unix(),
		Subject: 	userId,
	})
	return acssesToken.SignedString(m.signinKey)
}

func(m *Manager) Parce(accessToken string) (string, error){
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{},err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error: unexpected signing method")
		}
		return []byte(m.signinKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token нельзя привести к типу *tokenClaims")
	}

	return claims["sub"].(string), nil
}

func(m *Manager) NewRefreshToken() (string, error){
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	_, err := r.Read(b)
	if err != nil{
		return "", nil
	}

	return fmt.Sprintf("%x", b), nil
}