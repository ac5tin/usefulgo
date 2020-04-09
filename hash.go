package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns hashed version of password
func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPasswordHash compares hashes between password string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Mapfields - returns subset map
func Mapfields(d map[string]interface{}, fields []string) map[string]interface{} {
	retme := make(map[string]interface{})
	for _, field := range fields {
		retme[field] = d[field]
	}
	return retme
}

// Hashmap hashes a map using sha256 and returns the hash as string
func Hashmap(d map[string]interface{}) (string, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(b)
	return fmt.Sprintf("%x", h), nil
}
