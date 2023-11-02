package service

import (
	"crypto/sha256"

	postgresql "github.com/fishmanDK/internal/repository/postgreSQL"
)



type CheckService struct {
	repo *postgresql.PostgreDB
}

func NewCheckService(db *postgresql.PostgreDB) *CheckService {
	return &CheckService{
		repo: db,
	}
}

func (db *CheckService) CheckUser(app_name, email, password string) (int64, error) {
	hashPassword := HashPassword(password)

	return db.CheckUser(app_name, email, hashPassword)

}

func HashPassword(password string) string {
	data := []byte(password + salt)
	hashData := sha256.Sum256(data)

	return string(hashData[:])
}
