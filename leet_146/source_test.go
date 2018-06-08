package leet_146

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	const SET = 1
	const GET = 2
	type command struct {
		t      int
		param  []int
		expect int
	}
	var testCase = []struct {
		capacity int
		cmd      []command
	}{
		{
			capacity: 2,
			cmd: []command{
				{
					t:     SET,
					param: []int{1, 1},
				},
				{
					t:     SET,
					param: []int{2, 2},
				},
				{
					t:      GET,
					param:  []int{1},
					expect: 1,
				},
				{
					t:     SET,
					param: []int{3, 3},
				},
				{
					t:      GET,
					param:  []int{2},
					expect: -1,
				},
				{
					t:     SET,
					param: []int{4, 4},
				},
				{
					t:      GET,
					param:  []int{1},
					expect: -1,
				},
				{
					t:      GET,
					param:  []int{3},
					expect: 3,
				},
				{
					t:      GET,
					param:  []int{4},
					expect: 4,
				},
			},
		},
		{
			capacity: 3,
			cmd: []command{
				{
					t:     SET,
					param: []int{1, 1},
				},
				{
					t:     SET,
					param: []int{2, 2},
				},
				{
					t:     SET,
					param: []int{3, 3},
				},
				{
					t:      GET,
					param:  []int{3},
					expect: 3,
				},
				{
					t:     SET,
					param: []int{2, 4},
				},
				{
					t:      GET,
					param:  []int{2},
					expect: 4,
				},
				{
					t:     SET,
					param: []int{4, 4},
				},
				{
					t:      GET,
					param:  []int{1},
					expect: -1,
				},
				{
					t:      GET,
					param:  []int{3},
					expect: 3,
				},
				{
					t:      GET,
					param:  []int{4},
					expect: 4,
				},
				{
					t:     SET,
					param: []int{5, 6},
				},
				{
					t:     SET,
					param: []int{3, 8},
				},
				{
					t:      GET,
					param:  []int{5},
					expect: 6,
				},
				{
					t:      GET,
					param:  []int{3},
					expect: 8,
				},
				{
					t:     SET,
					param: []int{1, 1},
				},
				{
					t:      GET,
					param:  []int{4},
					expect: -1,
				},
				{
					t:      GET,
					param:  []int{1},
					expect: 1,
				},
			},
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		cache := Constructor(tc.capacity)
		for idx, c := range tc.cmd {
			if c.t == SET {
				cache.Put(c.param[0], c.param[1])
			} else {
				actual := cache.Get(c.param[0])
				ass.Equal(c.expect, actual, "idx: %d, command: %v", idx, c)
			}
		}
	}
}
