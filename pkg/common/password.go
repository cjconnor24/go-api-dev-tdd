package common

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(h), nil
}

func CheckPassword(plainText, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
}
