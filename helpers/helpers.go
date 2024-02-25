package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hash)
}