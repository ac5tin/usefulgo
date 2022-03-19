package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	// test hash password
	t.Run("test hash password", func(t *testing.T) {
		// hash password
		hash, err := HashPassword("password", 12)
		if err != nil {
			t.Errorf("hash password failed")
		}
		if !assert.NotEmpty(t, hash) {
			t.Errorf("hash password failed")
		}

		// verify hash
		if !assert.True(t, CheckPasswordHash("password", hash)) {
			t.Errorf("check password hash failed")
		}
	})
}

func FuzzHash(f *testing.F) {
	// fuzz pw
	f.Add("password", 12)
	// test hash password
	f.Fuzz(func(t *testing.T, a string, b int) {
		// hash password
		hash, err := HashPassword(a, b)
		if err != nil {
			t.Errorf("hash password failed")
		}
		if !assert.NotEmpty(t, hash) {
			t.Errorf("hash password failed")
		}

		// verify hash
		if !assert.True(t, CheckPasswordHash("password", hash)) {
			t.Errorf("check password hash failed")
		}
	})

}
