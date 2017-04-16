package leet_20

type Stack []byte

func (s *Stack) Pop() byte {
	c := s.Top()
	length := len(*s)
	if length <= 1 {
		*s = nil
	} else {
		*s = (*s)[:length-1]
	}
	return c
}

func (s *Stack) Empty() bool {
	return nil == s || *s == nil || len(*s) == 0
}

func (s *Stack) Push(x byte) {
	*s = append(*s, x)
}

func (s *Stack) Top() byte {
	length := len(*s)
	if 0 == length {
		return 0
	}
	return (*s)[length-1]
}

func (s *Stack) String() string {
	return string(*s)
}

func match(a, b byte) bool {
	if a == '[' && b == ']' {
		return true
	}
	if a == '(' && b == ')' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}

func isValid(s string) bool {
	arr := []byte(s)
	stack := new(Stack)
	for _, c := range arr {
		switch c {
		case '[', '{', '(':
			stack.Push(c)
			continue
		}
		top := stack.Top()
		if match(top, c) {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.Empty()
}
