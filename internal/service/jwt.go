package service

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	refresh_tokenTtl = time.Hour * 24 * 30
	access_tokenTtl = time.Minute * 15
	signInKey = "@(#tf53$*#$(RHfverib}#Rfrte)"
	salt     = "lsd2#tfv%2"
)

type Claims struct{
	jwt.StandardClaims
	id int64
	role string
}


func CreateAccessToken(id int64, role string) (string, error){
	
	acssesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims {
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(access_tokenTtl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
		role,
	})
	return acssesToken.SignedString([]byte(signInKey))
}

func CreateRefreshToken() (string, error){
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	_, err := r.Read(b)
	if err != nil{
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}