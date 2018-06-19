package leet_57

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	ns := newInterval.Start
	length := len(intervals)
	if length == 0 {
		return []Interval{newInterval}
	}

	if ns <= intervals[0].Start && newInterval.End >= intervals[length-1].End {
		return []Interval{newInterval}
	}
	idx := sort.Search(length, func(i int) bool {
		return ns <= intervals[i].Start
	})
	var ans []Interval
	if idx > 0 {
		ans = intervals[:idx-1]
		mergedInterval := overlap(intervals[idx-1], newInterval)
		if nil != mergedInterval {
			newInterval = *mergedInterval
		} else {
			ans = append(ans, intervals[idx-1])
		}
	}
	for i := idx; i < length; i++ {
		mergedInterval := overlap(newInterval, intervals[i])
		if nil == mergedInterval {
			t := make([]Interval, length-i)
			copy(t, intervals[i:])
			ans = append(ans, newInterval)
			ans = append(ans, t...)
			return ans
		}
		newInterval = *mergedInterval
	}
	return append(ans, newInterval)
}

func overlap(left, right Interval) *Interval {
	if right.Start > left.End {
		return nil
	}
	maxEnd := left.End
	if right.End > maxEnd {
		maxEnd = right.End
	}
	return &Interval{
		Start: left.Start,
		End:   maxEnd,
	}
}
