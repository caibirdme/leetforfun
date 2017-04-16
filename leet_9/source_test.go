package leet_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  int
	out bool
}

func TestIsPalindrome(t *testing.T) {
	var data = []Case{
		Case{1, true},
		Case{121, true},
		Case{123, false},
		Case{123321, true},
		Case{1234321, true},
		Case{4104, false},
		Case{13531, true},
		Case{-1, false},
		Case{-2147483648, false},
		Case{1000001, true},
		Case{1000011, false},
	}
	ass := assert.New(t)

	for _, tc := range data {
		ass.Equal(tc.out, isPalindrome(tc.in), "input:%d", tc.in)
	}
}
