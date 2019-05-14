package leet_201

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeBitwiseAnd(t *testing.T) {
	should := require.New(t)
	should.Equal(expect(5, 7), rangeBitwiseAnd(5, 7))
	should.Equal(expect(4894, 5949), rangeBitwiseAnd(4894, 5949))
	q := 10000000
	for i := 0; i < 5000; i++ {
		m := rand.Intn(q)
		n := m + rand.Intn(m/3)
		actual := rangeBitwiseAnd(m, n)
		should.Equal(expect(m, n), actual, "m: %d n: %d", m, n)
	}
}

func expect(m, n int) int {
	var pos uint
	for m != n {
		m >>= 1
		n >>= 1
		pos++
	}
	return m << pos
}
