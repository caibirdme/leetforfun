package qsort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	const times = 100
	const length = 100
	expect := make([]int, length)
	original := make([]int, length)
	ass := require.New(t)
	for i := 0; i < times; i++ {
		nums := rand.Perm(length)
		copy(expect, nums)
		copy(original, nums)
		sort.Ints(expect)
		Sort(original)
		ass.Equal(expect, original, "%v", nums)
	}
}
