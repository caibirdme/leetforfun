package leet_84

func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	q := new(humdrumQ)
	ans := heights[0]
	q.Push(elem{
		val: heights[0],
		pos: 0,
	})
	for i := 1; i < length; i++ {
		list := q.Range()
		for _, ele := range list {
			ans = max(ans, min(heights[i], ele.val)*(i-ele.pos+1))
		}
		ans = max(ans, heights[i])
		q.Push(elem{
			val: heights[i],
			pos: i,
		})
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type humdrumQ struct {
	data   []elem
	length int
	idx    int
}

func (q *humdrumQ) Range() []elem {
	return q.data[:q.length]
}

func (q *humdrumQ) Push(data elem) {
	if len(q.data) == 0 {
		q.data = append(q.data, data)
		q.length = 1
		return
	}
	last := q.length - 1
	for {
		if last >= 0 && data.less(q.data[last]) {
			data.pos = q.data[last].pos
			last--
		} else {
			break
		}
	}
	if last+1 < len(q.data) {
		q.data[last+1] = data
		q.length = last + 2
	} else {
		q.data = append(q.data, data)
		q.length++
	}
}

type elem struct {
	val int
	pos int
}

func (e elem) less(t elem) bool {
	return e.val < t.val
}
