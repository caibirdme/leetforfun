package leet_224

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	var testData = []struct {
		in  string
		out int
	}{
		{
			in:  `       2               -      1  +         2       `,
			out: 3,
		},
		{
			in:  `1+1`,
			out: 2,
		},
		{
			in:  `2-1  +2`,
			out: 3,
		},

		{
			in:  `(1+(4+5+2)-3)+(6+8)`,
			out: 23,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, calculate(tc.in), "%s", tc.in)
	}
}
