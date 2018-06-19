package leet_9

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in int
	out bool
}

func TestIsPalindrome(t *testing.T) {
	var data = []Case{
		Case{121, true},
		Case{2, true},
		Case{0, true},
		Case{1000001, true},
		Case{123321, true},
		Case{1233211, false},
		Case{1000010001, false},
	}
	ass := assert.New(t)
	for _,tc := range data {
		ass.Equal(tc.out, isPalindrome(tc.in), "in: %d", tc.in)
	}
}