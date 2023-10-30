package utils

import (
	"reflect"
	"strconv"
)

// NumParse parses a string to a number type.
func NumParse[T WholeNumber](str string) (T, error) {
	bits := 64
	signed := true
	switch reflect.TypeOf(T(0)).Kind() {
	case reflect.Uint8:
		signed = false
		fallthrough
	case reflect.Int8:
		bits = 8
	case reflect.Uint16:
		signed = false
		fallthrough
	case reflect.Int16:
		bits = 16
	case reflect.Uint32:
		signed = false
		fallthrough
	case reflect.Int32:
		bits = 32
	case reflect.Uint64:
		signed = false
	}

	// signed
	if signed {
		num, parseErr := strconv.ParseInt(str, 10, bits)
		if parseErr != nil {
			return 0, parseErr
		}
		return T(num), nil
	}
	// unsigned
	num, parseErr := strconv.ParseUint(str, 10, bits)
	if parseErr != nil {
		return 0, parseErr
	}

	return T(num), nil
}

// FloatParse parses a string to a float type.
func FloatParse[T Float](str string) (T, error) {
	f, parseErr := strconv.ParseFloat(str, 64)
	if parseErr != nil {
		return 0, parseErr
	}

	return T(f), nil
}
