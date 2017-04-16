package leet_15

import (
	"fmt"
	"sort"
	"strconv"
)

type set struct {
	arr  [][]int
	used map[string]struct{}
}

func (s *set) add(arr []int) {
	if !s.canAppend(arr) {
		return
	}
	key := s.genKey(arr)
	s.used[key] = struct{}{}
	s.arr = append(s.arr, arr)
}

func (s *set) genKey(arr []int) string {
	return fmt.Sprintf("%d%d%d", strconv.Itoa(arr[0]), strconv.Itoa(arr[1]), strconv.Itoa(arr[2]))
}

func (s *set) canAppend(arr []int) bool {
	if s.used == nil {
		s.used = make(map[string]struct{})
	}
	key := s.genKey(arr)
	if _, ok := s.used[key]; !ok {
		return true
	}
	return false
}

func (s *set) output() [][]int {
	return s.arr
}

func clear(arr []int) []int {
	length := len(arr)
	if length < 4 {
		return arr
	}
	for i := 0; i < len(arr)-3; i++ {
		if arr[i] != arr[i+1] || arr[i] != arr[i+2] {
			continue
		}
		var j int
		for j = i + 3; j < len(arr) && arr[i] == arr[j]; j++ {
		}
		arr = append(arr[:i+3], arr[j:]...)
		i = j - 1
	}
	return arr
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	nums = clear(nums)
	s := new(set)
	length := len(nums)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			target := -nums[i] - nums[j]
			idx := sort.SearchInts(nums[j+1:], target)
			if pos := j + 1 + idx; pos < length && nums[pos] == target {
				s.add([]int{nums[i], nums[j], target})
			}
		}
	}
	return s.output()
}
