package domain

type Store interface {
	CreateUser(user *User) (*User, error)
	DeleteUserByID(ID int64) error
	FindUserByEmail(Email string) (*User, error)
	FindUserById(ID int64) (*User, error)
}
