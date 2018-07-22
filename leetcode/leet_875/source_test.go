package leet_875

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinEatingSpeed(t *testing.T) {
	var testCase = []struct {
		piles  []int
		H      int
		expect int
	}{
		{
			piles:  []int{312884470},
			H:      968709470,
			expect: 1,
		},
		{
			piles:  []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184},
			H:      823855818,
			expect: 14,
		},
		{
			piles:  []int{3, 6, 7, 11},
			H:      8,
			expect: 4,
		},
		{
			piles:  []int{30, 11, 23, 4, 20},
			H:      5,
			expect: 30,
		},
		{
			piles:  []int{30, 11, 23, 4, 20},
			H:      6,
			expect: 23,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := minEatingSpeed(tc.piles, tc.H)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
