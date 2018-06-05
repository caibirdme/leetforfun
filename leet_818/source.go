package leet_818

func racecar_bfs(target int) int {
	if target == 1 {
		return 1
	}
	baseStep := 0
	if target < 0 {
		baseStep = 1
		target = -target
	}
	now := 0
	next := 1
	for next <= target {
		now = next
		baseStep++
		next = (next+1)<<1 - 1
	}
	if now == target {
		return baseStep
	}
	return baseStep + bfs(target-now, now+1)
}

type elem struct {
	v    int
	pos  int
	step int
}

type hash struct {
	v, pos int
}

func bfs(target int, v int) int {
	visited := make(map[hash]struct{})
	q := new(list)
	q.tail = -1
	e := elem{
		v:    v << 1,
		pos:  v,
		step: 1,
	}
	q.Push(e)
	visited[hash{e.v, e.pos}] = struct{}{}
	e = elem{
		v:    -1,
		pos:  0,
		step: 1,
	}
	q.Push(e)
	visited[hash{e.v, e.pos}] = struct{}{}
	for !q.Empty() {
		cur := q.Pop()
		e = elem{
			v:    cur.v << 1,
			pos:  cur.pos + cur.v,
			step: cur.step + 1,
		}
		if _, ok := visited[hash{e.v, e.pos}]; !ok && e.pos < target<<1 {
			q.Push(e)
			visited[hash{e.v, e.pos}] = struct{}{}
			if e.pos == target {
				return e.step
			}
		}
		e = elem{
			v:    reverse(cur.v),
			pos:  cur.pos,
			step: cur.step + 1,
		}
		if _, ok := visited[hash{e.v, e.pos}]; !ok {
			q.Push(e)
			visited[hash{e.v, e.pos}] = struct{}{}
		}
	}
	return 0
}

func reverse(v int) int {
	if v > 0 {
		return -1
	}
	return 1
}

type list struct {
	head, tail int
	buff       []elem
}

func (l *list) Pop() elem {
	defer func() {
		l.head++
	}()
	return l.buff[l.head]
}

func (l *list) Push(e elem) {
	l.tail++
	l.buff = append(l.buff, e)
}

func (l *list) Empty() bool {
	return l.head > l.tail
}

func racecar(target int) int {
	if target <= 1 {
		return target
	}
	mem := make([]int, (target<<1)+2)
	mem[0], mem[1], mem[2] = 0, 1, 4
	return dp(target, mem)
}

func dp(target int, mem []int) int {
	if target <= 1 {
		return target
	}
	if mem[target] != 0 {
		return mem[target]
	}
	n := log2(target)
	if 1<<uint(n) == target+1 {
		mem[target] = n
		return n
	}
	case1 := dp((1<<uint(n))-1-target, mem) + n + 1
	t := (1 << uint(n-1)) - 1
	case2 := dp(t, mem) + dp(target-t, mem) + 2
	dist := min(case1, case2)
	pre := n - 1
	for m := 1; m < pre; m++ {
		// target-(1<<pre -1 )+ (1<<m -1)
		t := target - (1 << uint(pre)) + (1 << uint(m))
		case3 := dp(t, mem) + n + m + 1
		dist = min(dist, case3)
	}
	mem[target] = dist
	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func log2(n int) int {
	t := 1
	exp := 0
	for t <= n {
		exp++
		t <<= 1
	}
	return exp
}
