package postgres

import (
	"github.com/cjconnor24/api-dev-tdd/pkg/common"
	"github.com/cjconnor24/api-dev-tdd/pkg/domain"
)

const (
	sqlCreateUser = `INSERT INTO users
	(name, email, password) VALUES($1, $2, $3)
	RETURNING id, name, email, password`
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
