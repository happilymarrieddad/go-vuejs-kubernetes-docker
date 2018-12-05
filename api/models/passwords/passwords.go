package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

// Encrypt - takes a password and encrypts it
func Encrypt(pass string) (string, error) {
	val, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(val), err
}

// Compare - this compares the hashed value in DB to the passed in value from the endpoint
func Compare(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}

	return true
}
