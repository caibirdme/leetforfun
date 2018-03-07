package leet_56

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	current := 0
	ans := []Interval{intervals[0]}
	length := len(intervals)
	for i := 1; i < length; i++ {
		newInterval := overlap(ans[current], intervals[i])
		if nil == newInterval {
			ans = append(ans, intervals[i])
			current++
		} else {
			ans[current] = *newInterval
		}
	}
	return ans
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

func quickSort(intervals []Interval, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	intervals[mid], intervals[left] = intervals[left], intervals[mid]
	pivot := partition(intervals, left, right)
	quickSort(intervals, left, pivot-1)
	quickSort(intervals, pivot+1, right)
}

func partition(intervals []Interval, left, right int) int {
	i, j := left, right
	t := intervals[i]
	for i < j {
		for ; j > i && !less(intervals[j], t); j-- {
		}
		intervals[i] = intervals[j]
		for ; i < j && !less(t, intervals[i]); i++ {
		}
		intervals[j] = intervals[i]
	}
	intervals[i] = t
	return i
}

func less(i, j Interval) bool {
	if i.Start < j.Start {
		return true
	}
	if i.Start == j.Start {
		return i.End < j.End
	}
	return false
}
