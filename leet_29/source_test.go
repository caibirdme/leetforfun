package leet_29

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type Case struct {
	in  [2]int
	out int
}

func TestDivide(t *testing.T) {
	var data = []Case{
		Case{
			[2]int{1, -1},
			-1,
		},
		Case{
			[2]int{0, 0},
			math.MaxInt32,
		},
		Case{
			[2]int{100, -2},
			-50,
		},
		Case{
			[2]int{100, 3},
			33,
		},
		Case{
			[2]int{0, -1000000},
			0,
		},
		Case{
			[2]int{123213123, 3},
			123213123 / 3,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, divide(tc.in[0], tc.in[1]), "in: %v", tc.in)
	}
}
