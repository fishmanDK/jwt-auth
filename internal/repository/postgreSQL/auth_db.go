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
	Id   int64  `db:"id"`
	Role string `db:"role"`
}

func (postgre *PostgreDB) Authentication(user jwtauth.User) (AuthResult, error) {
	var res AuthResult

	query := "SELECT users.id, role FROM users JOIN services ON users.id_service = services.id WHERE email = $1 and password_hash = $2 and services.name = $3"
	
	err := postgre.postgre_db.Get(&res, query, user.Email, user.Password, user.AppName)
	if err != nil || res.Id == 0{
		if res.Id == 0{
			return AuthResult{}, errors.New("error: данный пользователь отсутствует в БД")
		}
		return AuthResult{}, err
	}
	return res, nil
}

func (postgre *PostgreDB) CreateSession(userId int64, session Session) error {
	// query := fmt.Sprintf("INSERT INTO mytable (id, column1, column2, column3) VALUES (1, 'value1', 'value2', 'value3') ON CONFLICT (id) DO UPDATE SET column1 = EXCLUDED.column1,")
	query := "INSERT INTO jwt (user_id, refresh_token, expiresAt)"+
			 "VALUES ($1, $2, $3)"+
			 "ON CONFLICT (user_id) DO UPDATE SET " +
			 "refresh_token = EXCLUDED.refresh_token, " +
			 "expiresAt = EXCLUDED.expiresAt"

	_, err := postgre.postgre_db.Exec(query, userId, session.Refresh_token, session.ExpiresAt)

	return err	
}

// query := fmt.Sprintf("INSERT INTO users ")

// postgre.postgre_db.Exec()
