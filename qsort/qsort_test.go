package qsort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	const times = 100
	const length = 10000
	expect := make([]int, length)
	ass := require.New(t)
	for i := 0; i < times; i++ {
		nums := rand.Perm(length)
		copy(expect, nums)
		sort.Ints(expect)
		sort.Ints(nums)
		ass.Equal(expect, nums)
	}
}
