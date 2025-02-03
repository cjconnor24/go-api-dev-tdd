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

func TestFindUserByEmail(t *testing.T) {

	pStore := NewPostgresStore(testDB)

	user := &domain.User{
		Email:    "test@test.com",
		Password: "passwrd",
		Name:     "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	uByEmail, err := pStore.FindUserByEmail(createdUser.Email)
	if err != nil {
		t.Errorf("Expected to find user; but got error %q", err)
	}

	_, err = pStore.FindUserByEmail("unknownemail")
	if err == nil {
		t.Errorf("Expected error; but got %q", err)
	}

	if user.ID != createdUser.ID {
		t.Errorf("Expteced %q; but got %q", createdUser.ID, uByEmail.Email)
	}

	_ = pStore.DeleteUserByID(createdUser.ID)

}

func TestFindUserByID(t *testing.T) {

	pStore := NewPostgresStore(testDB)

	user := &domain.User{
		Email:    "test@test.com",
		Password: "passwrd",
		Name:     "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	uByID, err := pStore.FindUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("Expected to find user; but got error %q", err)
	}

	_, err = pStore.FindUserByID(-10)
	if err == nil {
		t.Errorf("Expected error; but got %q", err)
	}

	if uByID.ID != createdUser.ID {
		t.Errorf("Expteced %q; but got %q", createdUser.ID, uByID.ID)
	}

	_ = pStore.DeleteUserByID(createdUser.ID)

}
