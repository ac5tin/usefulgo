package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns hashed version of password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}


// CheckPasswordHash compares hashes between password string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}