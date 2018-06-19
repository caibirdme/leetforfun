package leet_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type tCase struct {
	str  string
	rows int
	ans  string
}

func TestConvert(t *testing.T) {
	var data = []tCase{
		tCase{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		tCase{
			"HELLOMYNAMEISDEEN",
			4,
			"HYSEMNIDLOAEENLME",
		},
		tCase{
			"A",
			1,
			"A",
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, convert(tc.str, tc.rows))
	}
}
