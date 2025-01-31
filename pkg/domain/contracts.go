package domain

type Store interface {
	CreateUser(user *User) (*User, error)
	DeleteUserByID(ID int64) error
}
