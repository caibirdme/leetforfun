package leet_126

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	endIdx := inSlice(wordList, endWord)
	if endIdx == -1 {
		return nil
	}
	lenwl := len(wordList)
	q := newQ()
	q.Push(endIdx)
	record := make([][]int, lenwl)
	dist := make(map[int]int)
	shortest := -1
	dist[endIdx] = 0
	var ans [][]string
	var acc = []string{beginWord}
	for !q.Empty() {
		cur := q.Pop()
		nextStep := dist[cur] + 1
		if shortest != -1 && nextStep > shortest {
			continue
		}
		if canChange([]byte(wordList[cur]), []byte(beginWord)) {
			shortest = nextStep
			showPath(cur, endIdx, wordList, record, &ans, &acc)
			continue
		}
		for i := 0; i < lenwl; i++ {
			if i == cur || !canChange([]byte(wordList[cur]), []byte(wordList[i])) {
				continue
			}
			s, found := dist[i]
			if found && s != nextStep {
				continue
			}
			if !found {
				dist[i] = nextStep
			}
			record[i] = append(record[i], cur)
			q.Push(i)
		}
	}
	return ans
}

func showPath(beginIdx, endIdx int, wordList []string, record [][]int, ans *[][]string, acc *[]string) {
	if beginIdx == endIdx {
		*ans = append(*ans, append(append([]string(nil), (*acc)...), wordList[endIdx]))
		return
	}
	*acc = append(*acc, wordList[beginIdx])
	for _, nextIdx := range record[beginIdx] {
		showPath(nextIdx, endIdx, wordList, record, ans, acc)
	}
	*acc = (*acc)[:len(*acc)-1]
}

func inSlice(arr []string, target string) int {
	for idx, s := range arr {
		if s == target {
			return idx
		}
	}
	return -1
}

func canChange(s1, s2 []byte) bool {
	lens1 := len(s1)
	if lens1 != len(s2) {
		return false
	}
	count := 0
	for i := 0; i < lens1; i++ {
		if s1[i] != s2[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}

type queue struct {
	buf        []int
	head, tail int
	inQ        map[int]struct{}
}

func newQ() *queue {
	return &queue{tail: -1, inQ: make(map[int]struct{})}
}

func (q *queue) Empty() bool {
	return q.head > q.tail
}

func (q *queue) Push(idx int) {
	if _, ok := q.inQ[idx]; ok {
		return
	}
	q.inQ[idx] = struct{}{}
	q.tail++
	q.buf = append(q.buf, idx)
}

func (q *queue) Pop() int {
	t := q.buf[q.head]
	delete(q.inQ, t)
	q.head++
	return t
}
