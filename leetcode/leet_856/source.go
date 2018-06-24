package leet_856

const (
	left  byte = '('
	right byte = ')'
)

type queue struct {
	buf []byte
	sp  int
}

func (q *queue) Push(op byte) {
	q.sp++
	if len(q.buf) > q.sp {
		q.buf[q.sp] = op
	} else {
		q.buf = append(q.buf, op)
	}
}

func (q *queue) Pop() byte {
	t := q.buf[q.sp]
	q.sp--
	return t
}

func (q *queue) Top() byte {
	return q.buf[q.sp]
}

func newQ() *queue {
	return &queue{sp: -1}
}

type queueInt struct {
	buf []int
	sp  int
}

func (q *queueInt) Push(num int) {
	q.sp++
	if len(q.buf) > q.sp {
		q.buf[q.sp] = num
	} else {
		q.buf = append(q.buf, num)
	}
}

func (q *queueInt) Pop() int {
	t := q.buf[q.sp]
	q.sp--
	return t
}

func (q *queueInt) Top() int {
	return q.buf[q.sp]
}

func (q *queueInt) Empty() bool {
	return q.sp < 0
}

func newQInt() *queueInt {
	return &queueInt{sp: -1}
}

func scoreOfParentheses(S string) int {
	if S == "" {
		return 0
	}
	q := newQ()
	qi := newQInt()
	lens := len(S)
	for i := 1; i < lens; i++ {
		if S[i] == ')' {
			if S[i-1] == '(' {
				// ()
				qi.Push(1)
			} else {
				// ))
				t := qi.Pop()
				for q.Top() == '+' {
					q.Pop()
					t += qi.Pop()
				}
				q.Pop()
				t <<= 1
				qi.Push(t)
			}
		} else {
			if S[i-1] == ')' {
				// )(
				q.Push('+')
			} else {
				// ((
				q.Push('*')
			}
		}
	}
	t := 0
	for !qi.Empty() {
		t += qi.Pop()
	}
	return t
}
