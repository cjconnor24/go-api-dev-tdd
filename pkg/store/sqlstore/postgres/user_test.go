package postgres

import (
	"testing"

	"github.com/cjconnor24/api-dev-tdd/pkg/domain"
)

func TestCreateUser(t *testing.T) {

	pStore := NewPostgresStore(testDB)
	oldPassword := "password"

	user := &domain.User{
		Email:    "john.doe@gmail.com",
		Password: oldPassword,
		Name:     "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fail()
	}

	if createdUser.ID == 0 {
		t.Errorf("wanted id not to be 0")
	}

	if user.Name != createdUser.Name {
		t.Errorf("expected %q; go %q", user.Name, createdUser.Name)
	}

	if createdUser.Password == oldPassword {
		t.Error("password was not hashed")
	}

	err = pStore.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expect nil error during DeleteUserID; got %d", err)
	}
}
