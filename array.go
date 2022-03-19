package utils

// ArrContains check is array contains a string

func ArrContains[T Primitive](arr []T, tgt T) bool {
	for _, a := range arr {
		if a == tgt {
			return true
		}
	}
	return false
}

// ArrRM removes element from array (ordered)
func ArrRm[T Primitive](slice *[]T, t T) {
	for i, a := range *slice {
		if a == t {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return
		}
	}
	return
}

// ArrRmF removes element from array (unordered,fast)
func ArrRmF[T any](slice *[]T, index uint32) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	(*slice)[len(*slice)-1] = *new(T)
	(*slice) = (*slice)[:len(*slice)-1]
}

// ArrRmS removes element from array (ordered,slow)
func ArrRmS[T any](slice *[]T, index uint32) {
	copy((*slice)[index:], (*slice)[index+1:])
	(*slice)[len(*slice)-1] = *new(T)
	(*slice) = (*slice)[:len(*slice)-1]
}
