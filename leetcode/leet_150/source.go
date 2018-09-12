package leet_150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	s := new(stack)
	for _, token := range tokens {
		switch token {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "-":
			s.Push(-s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		case "/":
			a, b := s.Pop(), s.Pop()
			s.Push(b / a)
		default:
			a, _ := strconv.Atoi(token)
			s.Push(a)
		}
	}
	return s.Pop()
}

type stack struct {
	buf []int
}

func (s *stack) Push(x int) {
	s.buf = append(s.buf, x)
}

func (s *stack) Pop() int {
	last := len(s.buf) - 1
	t := s.buf[last]
	s.buf = s.buf[:last]
	return t
}
