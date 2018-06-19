package leet_60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	var testData = []struct {
		n, k int
		out  string
	}{
		{
			n:   3,
			k:   4,
			out: "231",
		},
		{
			n:   4,
			k:   17,
			out: "3412",
		},
		{
			out: "",
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, getPermutation(tc.n, tc.k))
	}
}
