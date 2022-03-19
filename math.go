package utils

// IntAbs - returns the absolute value of x
func IntAbs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
