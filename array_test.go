package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	// test array contains
	t.Run("test array contains", func(t *testing.T) {
		// string
		{
			arr := []string{"a", "b", "c"}
			if !assert.True(t, ArrContains(arr, "a")) {
				t.Errorf("array contains failed")
			}
			if !assert.True(t, ArrContains(arr, "b")) {
				t.Errorf("array contains failed")
			}
			if !assert.True(t, ArrContains(arr, "c")) {
				t.Errorf("array contains failed")
			}
			if !assert.False(t, ArrContains(arr, "d")) {
				t.Errorf("array contains failed")
			}
		}

	})
	// test array remove
	t.Run("test array remove", func(t *testing.T) {
		// string
		{
			arr := []string{"a", "b", "c"}
			ArrRm(&arr, "b")
			if !assert.Equal(t, []string{"a", "c"}, arr) {
				t.Errorf("array remove failed")
			}
		}

		// int
		{
			arr := []int{1, 2, 3}
			ArrRm(&arr, 2)
			if !assert.Equal(t, []int{1, 3}, arr) {
				t.Errorf("array remove failed")
			}
		}

	})
	// test array remove fast
	t.Run("test array remove fast", func(t *testing.T) {
		// string
		{
			arr := []string{"a", "b", "c"}
			ArrRmF(&arr, 1)
			if !assert.Equal(t, []string{"a", "c"}, arr) {
				t.Errorf("array remove failed")
			}
		}

		// int
		{
			arr := []int{1, 2, 3}
			ArrRmF(&arr, 1)
			if !assert.Equal(t, []int{1, 3}, arr) {
				t.Errorf("array remove failed")
			}
		}
		// custom struct
		{
			type myStruct struct {
				Name string
				Age  int
			}
			arr := make([]myStruct, 0)
			arr = append(arr, myStruct{Name: "a", Age: 1})
			arr = append(arr, myStruct{Name: "b", Age: 2})
			arr = append(arr, myStruct{Name: "c", Age: 3})
			ArrRmF(&arr, 1)
			if !assert.Equal(t, []myStruct{{Name: "a", Age: 1}, {Name: "c", Age: 3}}, arr) {
				t.Errorf("array remove failed")
			}
		}

	})
	// test array remove stable
	t.Run("test array remove stable", func(t *testing.T) {
		// string
		{
			arr := []string{"a", "b", "c"}
			ArrRmS(&arr, 1)
			if !assert.Equal(t, []string{"a", "c"}, arr) {
				t.Errorf("array remove failed")
			}
		}

		// int
		{
			arr := []int{1, 2, 3}
			ArrRmS(&arr, 1)
			if !assert.Equal(t, []int{1, 3}, arr) {
				t.Errorf("array remove failed")
			}
		}
		// custom struct
		{
			type myStruct struct {
				Name string
				Age  int
			}
			arr := make([]myStruct, 0)
			arr = append(arr, myStruct{Name: "a", Age: 1})
			arr = append(arr, myStruct{Name: "b", Age: 2})
			arr = append(arr, myStruct{Name: "c", Age: 3})
			ArrRmS(&arr, 1)
			if !assert.Equal(t, []myStruct{{Name: "a", Age: 1}, {Name: "c", Age: 3}}, arr) {
				t.Errorf("array remove failed")
			}
		}

	})
}
