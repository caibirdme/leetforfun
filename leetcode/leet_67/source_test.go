package leet_67

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	var testCase = []struct {
		a, b   string
		expect string
	}{
		{
			a:      "1111101",
			b:      "11",
			expect: "10000000",
		},
		{
			a:      "1111100",
			b:      "11",
			expect: "1111111",
		},
		{
			a:      "1111100",
			b:      "0",
			expect: "1111100",
		},
		{
			a:      "0",
			b:      "1",
			expect: "1",
		},
		{
			a:      "11",
			b:      "1",
			expect: "100",
		},
		{
			a:      "11",
			b:      "11",
			expect: "110",
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := addBinary(tc.a, tc.b)
		should.Equal(tc.expect, actual, "a: %s b: %s", tc.a, tc.b)
	}
}
