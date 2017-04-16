package leet_7

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type Case struct {
	in, out int
}

func TestReverse(t *testing.T) {
	var data = []Case{
		Case{
			123,
			321,
		},
		Case{
			-123,
			-321,
		},
		Case{
			0, 0,
		},
		Case{
			math.MaxInt32,
			0,
		},
		Case{
			-math.MaxInt32,
			0,
		},
		Case{
			math.MaxInt32 - 6,
			1463847412,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, reverse(tc.in))
	}
}
