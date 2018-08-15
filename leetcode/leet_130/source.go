package leet_130

func solve(board [][]byte) {
	m := len(board)
	if m <= 1 {
		return
	}
	n := len(board[0])
	if n <= 1 {
		return
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'U'
			q := newQ()
			q.Push(pos{i: i, j: 0})
			fill(q, board)
		}
		if board[i][n-1] == 'O' {
			board[i][n-1] = 'U'
			q := newQ()
			q.Push(pos{i: i, j: n - 1})
			fill(q, board)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			board[0][j] = 'U'
			q := newQ()
			q.Push(pos{i: 0, j: j})
			fill(q, board)
		}
		if board[m-1][j] == 'O' {
			board[m-1][j] = 'U'
			q := newQ()
			q.Push(pos{i: m - 1, j: j})
			fill(q, board)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'U' {
				board[i][j] = 'O'
			}
		}
	}
}

func fill(q *queue, board [][]byte) {
	m, n := len(board), len(board[0])
	for !q.Empty() {
		p := q.Pop()
		i, j := p.i, p.j
		//up
		if i-1 >= 0 && board[i-1][j] == 'O' {
			board[i-1][j] = 'U'
			q.Push(pos{i: i - 1, j: j})
		}
		//down
		if i+1 < m && board[i+1][j] == 'O' {
			board[i+1][j] = 'U'
			q.Push(pos{i: i + 1, j: j})
		}
		//left
		if j-1 >= 0 && board[i][j-1] == 'O' {
			board[i][j-1] = 'U'
			q.Push(pos{i: i, j: j - 1})
		}
		//right
		if j+1 < n && board[i][j+1] == 'O' {
			board[i][j+1] = 'U'
			q.Push(pos{i: i, j: j + 1})
		}
	}
}

type pos struct {
	i, j int
}

type queue struct {
	buf  []pos
	head int
}

func (q *queue) Push(p pos) {
	q.buf = append(q.buf, p)
}

func (q *queue) Pop() pos {
	q.head++
	return q.buf[q.head]
}

func (q *queue) Empty() bool {
	return q.head+1 >= len(q.buf)
}

func newQ() *queue {
	return &queue{head: -1}
}
