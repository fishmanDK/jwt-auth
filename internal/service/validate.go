package service

import postgresql "github.com/fishmanDK/internal/repository/postgreSQL"

type ValidateService struct{
	repo *postgresql.PostgreDB
}

func NewValidateService(db *postgresql.PostgreDB) *ValidateService{
	return &ValidateService{
		repo: db,
	}
}