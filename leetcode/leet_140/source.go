package leet_140

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	q := make([]*list, len(s)+1)
	search(s, 0, wordDict, q)
	result := make([]string, 0)
	temp := make([]string, 0)
	fullfill(0, q, wordDict, temp, &result)
	return result
}

func fullfill(pos int, q []*list, wordDict []string, temp []string, result *[]string) {
	if q[pos] == emptyList {
		*result = append(*result, strings.Join(temp, " "))
		return
	}
	for _, nd := range q[pos].buf {
		temp = append(temp, wordDict[nd.dictIDX])
		fullfill(nd.nextPos, q, wordDict, temp, result)
		temp = temp[:len(temp)-1]
	}
}

func search(s string, startPos int, wordDict []string, q []*list) {
	lens := len(s)
	if startPos >= lens {
		q[startPos] = emptyList
		return
	}
	if q[startPos] != nil {
		return
	}
	q[startPos] = new(list)
	p := s[startPos]
	for i := 0; i < len(wordDict); i++ {
		lenwd := len(wordDict[i])
		endPos := startPos + lenwd
		if p != wordDict[i][0] || endPos > lens {
			continue
		}
		if s[startPos:endPos] == wordDict[i] {
			search(s, endPos, wordDict, q)
			if q[endPos] != nil && (q[endPos].Len() > 0 || q[endPos] == emptyList) {
				q[startPos].Push(i, endPos)
			}
		}
	}
}

var emptyList = new(list)

type list struct {
	buf  []node
	last int
}

type node struct {
	dictIDX int
	nextPos int
}

func newList() *list {
	return &list{last: -1}
}

func (l *list) Len() int {
	return len(l.buf)
}

func (l *list) Push(dictIDX, nextPos int) {
	l.last++
	nd := node{dictIDX, nextPos}
	if len(l.buf) > l.last {
		l.buf[l.last] = nd
	} else {
		l.buf = append(l.buf, nd)
	}
}

func (l *list) Pop() {
	l.last--
}
