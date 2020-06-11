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

// ArrRmiF - fast array remover by index struct (order unmaintained)
type ArrRmiF struct{}

// ArrRmiS - Slow array remover by index struct (order maintained)
type ArrRmiS struct{}

// NewArrRmiF returns new ArrRmiF (fast array index remover)
func NewArrRmiF() ArrRmiF {
	return ArrRmiF{}
}

// NewArrRmiS returns new ArrRmiS (slow array index remover)
func NewArrRmiS() ArrRmiS {
	return ArrRmiS{}
}

// String - []string remover
func (a ArrRmiF) String(slice *[]string, index uint32) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	(*slice)[len(*slice)-1] = ""
	(*slice) = (*slice)[:len(*slice)-1]
}

// Object - []map[string]interface{} remover
func (a ArrRmiF) Object(slice *[]map[string]interface{}, index uint32) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	(*slice)[len(*slice)-1] = map[string]interface{}{}
	(*slice) = (*slice)[:len(*slice)-1]
}

// Int - []int remover
func (a ArrRmiF) Int(slice *[]int, index uint32) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	(*slice)[len(*slice)-1] = 0
	(*slice) = (*slice)[:len(*slice)-1]
}

// Object - []map[string]interface{} remover
func (a ArrRmiS) Object(slice *[]map[string]interface{}, index uint32) {
	copy((*slice)[index:], (*slice)[index+1:])
	(*slice)[len(*slice)-1] = map[string]interface{}{}
	(*slice) = (*slice)[:len(*slice)-1]
}

// String - []string remover
func (a ArrRmiS) String(slice *[]string, index uint32) {
	copy((*slice)[index:], (*slice)[index+1:])
	(*slice)[len(*slice)-1] = ""
	(*slice) = (*slice)[:len(*slice)-1]
}

// Int - []int remover
func (a ArrRmiS) Int(slice *[]int, index uint32) {
	copy((*slice)[index:], (*slice)[index+1:])
	(*slice)[len(*slice)-1] = 0
	(*slice) = (*slice)[:len(*slice)-1]
}
