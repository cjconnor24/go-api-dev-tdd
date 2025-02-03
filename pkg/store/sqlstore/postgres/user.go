package postgres

import (
	"database/sql"
	"errors"

	"github.com/cjconnor24/api-dev-tdd/pkg/common"
	"github.com/cjconnor24/api-dev-tdd/pkg/domain"
)

const (
	sqlCreateUser = `INSERT INTO users
	(name, email, password) VALUES($1, $2, $3)
	RETURNING id, name, email, password`
	sqlDeleteUserByID  = `DELETE FROM users WHERE id = $1`
	sqlFindUserByEmail = `SELECT id, name, email FROM users WHERE email =$1`
	sqlFindUserByID    = `SELECT id, name, email FROM users WHERE id=$1`
)

func (q *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	user.Password, _ = common.PasswordHash(user.Password)

	err := q.db.QueryRow(sqlCreateUser, user.Name, user.Email, user.Password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (q *postgresStore) DeleteUserByID(ID int64) error {
	_, err := q.db.Exec(sqlDeleteUserByID, ID)
	if err != nil {
		return err
	}

	return nil
}

func (q *postgresStore) FindUserByEmail(Email string) (*domain.User, error) {
	user := &domain.User{}

	err := q.db.QueryRow(sqlFindUserByEmail, Email).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	// 54:47
	return user, nil

}

func (q *postgresStore) FindUserByID(ID int64) (*domain.User, error) {
	user := &domain.User{}

	err := q.db.QueryRow(sqlFindUserByID, ID).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
