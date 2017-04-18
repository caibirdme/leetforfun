package leet_13

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in string
	out int
}

func TestRomanToInt(t *testing.T) {
	var data = []Case{
		Case{"I", 1},
		Case{"XI", 11},
		Case{"V", 5},
		Case{"IV", 4},
		Case{"DCCXLVII", 747},
		Case{"MMMCMXCIX", 3999},
		Case{"MMDCCCXCV", 2895},
		Case{"L", 50},
		Case{"XCIX", 99},
		Case{"D", 500},
	}
	ass := assert.New(t)
	for _,tc := range data {
		ass.Equal(tc.out, romanToInt(tc.in), "in: %s", tc.in)
	}
}
