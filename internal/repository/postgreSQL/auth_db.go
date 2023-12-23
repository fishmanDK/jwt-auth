package postgresql

import (
	"errors"
	"time"

	jwtauth "github.com/fishmanDK"
)

type Session struct {
	Refresh_token string
	ExpiresAt     time.Time
}

type AuthResult struct {
	Id       int64  `db:"id"`
	UserName string `db:"user_name"`
	Role     string `db:"role"`
}

func (postgre *PostgreDB) Authentication(user jwtauth.User) (AuthResult, error) {
	var res AuthResult

	query := "SELECT users.id, role, user_name FROM users JOIN services ON users.id_service = services.id WHERE email = $1 and password_hash = $2 and services.name = $3"

	err := postgre.postgre_db.Get(&res, query, user.Email, user.Password, user.AppName)
	if err != nil || res.Id == 0 {
		if res.Id == 0 {
			return AuthResult{}, errors.New("error: данный пользователь отсутствует в БД")
		}
		return AuthResult{}, err
	}
	return res, nil
}

func (postgre *PostgreDB) CreateSession(userId int64, session Session) error {
	query := "INSERT INTO jwt (user_id, refresh_token, expiresAt)" +
		"VALUES ($1, $2, $3)" +
		"ON CONFLICT (user_id) DO UPDATE SET " +
		"refresh_token = EXCLUDED.refresh_token, " +
		"expiresAt = EXCLUDED.expiresAt"

	_, err := postgre.postgre_db.Exec(query, userId, session.Refresh_token, session.ExpiresAt)

	return err
}

func (postgre *PostgreDB) CreateUser(newUser jwtauth.CreateUser) (int64, error) {
	var user_id int64

	query := "INSERT INTO users (id_service, user_name, first_name, last_name, email, password_hash, role) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7);"
	_, err := postgre.postgre_db.Exec(query, newUser.AppId, newUser.UserName, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.Role)
	if err != nil {
		return 0, err
	}

	err = postgre.postgre_db.Get(&user_id, "SELECT id FROM users WHERE email = $1 and id_service = $2", newUser.Email, newUser.AppId)
	if err != nil {
		return 0, nil
	}
	return user_id, nil
}

// query := fmt.Sprintf("INSERT INTO users ")

// postgre.postgre_db.Exec()
