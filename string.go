package toolbox

import (
	"math"
	"math/rand"
	"unicode/utf8"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)
)

// RandomString returns string of utf8-length size, and with char from charset.
func RandomString(size int, charset []rune) string {
	if size < 0 {
		panic("RandomString: size should greater than 0")
	}
	if len(charset) == 0 {
		panic("charset should not be empty")
	}
	result := make([]rune, 0, size)
	length := len(charset)
	for i := 0; i < size; i++ {
		result = append(result, charset[rand.Intn(length)])
	}
	return string(result)
}

// SubString returns result with offset and size.
func SubString[T ~string](str T, offset int, size uint) T {
	length := len(str)
	if length == 0 {
		return ""
	}
	if offset < 0 {
		offset = length + offset
		if offset < 0 {
			offset = 0
		}
	}
	rightIndex := int(Min(uint(length), uint(offset)+size))
	return str[offset:rightIndex]
}

// ChunkString returns chunks whose lengths are len(size), and from str.
func ChunkString[T ~string](str T, size int) []T {
	if size <= 0 {
		panic("size should be greater than 0")
	}
	length := len(str)
	result := make([]T, int(math.Ceil(float64(len(str))/float64(size))))
	for index := range result {
		result[index] = str[index*size : Min((index+1)*size, length)]
	}
	return result
}

// RuneLength returns rune count in str.
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}
