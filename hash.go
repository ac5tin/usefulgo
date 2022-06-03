package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns hashed version of password
func HashPassword(password string, cost int) (string, error) {
	// ensure cost is within allowed range
	{
		const MIN_COST = 4
		const MAX_COST = 31
		if cost < MIN_COST {
			cost = MIN_COST
		}
		if cost > MAX_COST {
			cost = MAX_COST
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPasswordHash compares hashes between password string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Mapfields - returns subset map
func Mapfields[T map[string]any](d *T, fields *[]string) *T {
	retme := make(T)
	for _, field := range *fields {
		retme[field] = (*d)[field]
	}
	return &retme
}

// Hashmap hashes a map using sha256 and returns the hash as string
func Hashmap[T map[string]any](d *T) (string, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(b)
	return fmt.Sprintf("%x", h), nil
}
