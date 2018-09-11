package leet_903

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPow(t *testing.T) {
	var testCase = []struct {
		x, y int
		m    int
	}{
		{
			x: 2,
			y: 5,
			m: 6,
		},
		{
			x: 7,
			y: 8,
			m: 9,
		},
		{
			x: 2,
			y: 20,
			m: 591,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := pow(tc.x, tc.y, tc.m)
		should.Equal(int(math.Pow(float64(tc.x), float64(tc.y)))%tc.m, actual, "x: %d y:%d m:%d", tc.x, tc.y, tc.m)
	}
}

func TestC(t *testing.T) {
	var testCase = []struct {
		n, m   int
		expect int
	}{
		{
			n:      5,
			m:      3,
			expect: 10,
		},
		{
			n:      4,
			m:      2,
			expect: 6,
		},
		{
			n:      5,
			m:      5,
			expect: 1,
		},
		{
			n:      6,
			m:      3,
			expect: 20,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := calc.C(tc.n, tc.m)
		should.Equal(tc.expect, actual, "n:%d m: %d", tc.n, tc.m)
	}
}

func TestNumPermsDISequence(t *testing.T) {
	var testCase = []struct {
		S      string
		expect int
	}{
		{
			S:      "DDDID",
			expect: 14,
		},
		{
			S:      "IDDDIIDIIIIIIIIDIDID",
			expect: 853197538,
		},
		{
			S:      "DDID",
			expect: 9,
		},
		{
			S:      "IDI",
			expect: 5,
		},
		{
			S:      "ID",
			expect: 2,
		},
		{
			S:      "DI",
			expect: 2,
		},
		{
			S:      "DID",
			expect: 5,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := numPermsDISequence(tc.S)
		should.Equal(tc.expect, actual, "S: %s", tc.S)
	}
}
