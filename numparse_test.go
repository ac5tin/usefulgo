package utils

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumParse(t *testing.T) {
	t.Run("Single digit number", func(t *testing.T) {
		num, err := NumParse[uint8]("1")
		assert.NoError(t, err)
		assert.Equal(t, uint8(1), num)
	})
	t.Run("Negative numbers", func(t *testing.T) {
		num, err := NumParse[int8]("-16")
		assert.NoError(t, err)
		assert.Equal(t, int8(-16), num)
	})
	t.Run("Non numbers", func(t *testing.T) {
		num, err := NumParse[int8]("yo 16")
		assert.Error(t, err)
		assert.Equal(t, int8(0), num)
	})

	t.Run("Number overflow", func(t *testing.T) {
		_, err := NumParse[uint8]("10000")
		assert.Error(t, err)
	})

	t.Run("Negative number in unsigned integer type", func(t *testing.T) {
		_, err := NumParse[uint8]("-10")
		assert.Error(t, err)
	})

	t.Run("Float number", func(t *testing.T) {
		_, err := NumParse[uint8]("-10.2")
		assert.Error(t, err)
	})
}

func TestFloatParse(t *testing.T) {
	t.Run("float32 number", func(t *testing.T) {
		num, err := FloatParse[float32]("1.12")
		assert.NoError(t, err)
		assert.Equal(t, float32(1.12), num)
	})
	t.Run("overflow float32 precision", func(t *testing.T) {
		// float32 max precision is 7 digits
		num, err := FloatParse[float32]("1.123456789101112")
		assert.NoError(t, err)
		assert.Equal(t, float32(1.12345678), num)
	})
	t.Run("negative float", func(t *testing.T) {
		num, err := FloatParse[float64]("-16.9000000033")
		assert.NoError(t, err)
		assert.Equal(t, float64(-16.9000000033), num)
	})
	t.Run("Non numbers", func(t *testing.T) {
		num, err := FloatParse[float32]("yo 16")
		assert.Error(t, err)
		assert.Equal(t, float32(0), num)
	})
}

func FuzzNumParse(f *testing.F) {
	f.Add("1") // seed corpus #0
	f.Add("-16")
	f.Add("yo 16")
	f.Add("1000") // test overflow
	f.Add("01000")
	f.Add("-1000")
	f.Fuzz(func(t *testing.T, str string) {
		expectVal, expectErr := strconv.ParseInt(str, 10, 8)
		v, parseErr := NumParse[int8](str)
		if expectErr != nil {
			// should return error
			if !assert.Error(t, parseErr) {
				t.Log("expectErr:", expectErr)
			}
			t.Skip()
		}
		// should parse successfully
		assert.NoError(t, parseErr)

		assert.GreaterOrEqual(t, v, int8(math.MinInt8))
		assert.LessOrEqual(t, v, int8(math.MaxInt8))
		// should be equal to expected value
		assert.Equal(t, int8(expectVal), v)
	})

	f.Fuzz(func(t *testing.T, str string) {
		expectVal, expectErr := strconv.ParseUint(str, 10, 8)
		v, parseErr := NumParse[uint8](str)
		if expectErr != nil {
			// should return error
			if !assert.Error(t, parseErr) {
				t.Log("expectErr:", expectErr)
			}
			t.Skip()
		}
		// should parse successfully
		assert.NoError(t, parseErr)

		assert.GreaterOrEqual(t, v, uint8(0)) // no math.MinUint8
		assert.LessOrEqual(t, v, uint8(math.MaxUint8))
		// should be equal to expected value
		assert.Equal(t, uint8(expectVal), v)
	})
}

func BenchmarkNumParse(b *testing.B) {
	b.Run("signed integers (int8)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NumParse[int8]("100")
		}
	})
	b.Run("signed integers (int64)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NumParse[int64]("100000")
		}
	})

	b.Run("unsigned integers (uint64)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NumParse[int64]("100000")
		}
	})

	b.Run("non integers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NumParse[int64]("yoyo 100000")
		}
	})
}
