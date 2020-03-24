package utils

// ArrContains check is array contains a string
func ArrContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// ArrRMS removes string element from array (ordered)
func ArrRMS(slice *[]string, t string) {
	for i, a := range *slice {
		if a == t {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return
		}
	}
	return
}
