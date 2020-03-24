package utils

import (
	"github.com/google/uuid"
)

// GenUUIDV4 generates uuidv4 uuid
func GenUUIDV4() string {
	uid := uuid.New()

	return uid.String()
}
