package utils

import (
	"strings"
)

// RmDash removes dashes from strings
func RmDash(tgt string) string {
	return strings.Replace(tgt, "-", "", -1)
}
