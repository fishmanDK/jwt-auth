package repository

import (
	jwtauth "github.com/fishmanDK"
	postgresql "github.com/fishmanDK/internal/repository/postgreSQL"
)

type Storage struct {
	DB
}

type DB interface {
	Authentication(user jwtauth.User) (postgresql.AuthResult, error)
	CreateSession(userId int64, session postgresql.Session) error
	CreateUser(newUser jwtauth.CreateUser) (int64, error)
}

func NewStorage(db_name string) (*Storage, error) {
	var (
		db  DB
		err error
	)
	switch db_name {
	case "postgres":
		cfg := postgresql.InitPostgreConfig()
		db, err = postgresql.NewPostgreDB(cfg)
	default:
		cfg := postgresql.InitPostgreConfig()
		db, err = postgresql.NewPostgreDB(cfg)
	}
	return &Storage{
		DB: db,
	}, err
}
