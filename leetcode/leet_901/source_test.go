package leet_901

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStockSpanner(t *testing.T) {
	var testCase = []struct {
		price  int
		expect int
	}{
		{
			price:  100,
			expect: 1,
		},
		{
			price:  80,
			expect: 1,
		},
		{
			price:  60,
			expect: 1,
		},
		{
			price:  70,
			expect: 2,
		},
		{
			price:  60,
			expect: 1,
		},
		{
			price:  75,
			expect: 4,
		},
		{
			price:  85,
			expect: 6,
		},
	}
	should := require.New(t)
	obj := Constructor()
	for idx, tc := range testCase {
		actual := obj.Next(tc.price)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
