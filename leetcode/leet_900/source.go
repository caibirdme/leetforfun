package leet_900

type RLEIterator struct {
	A   []int
	pos int
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{A: A, pos: 0}
}

func (this *RLEIterator) Next(n int) int {
	for n > 0 {
		if this.pos >= len(this.A) {
			return -1
		}
		if t := this.A[this.pos]; t > 0 {
			if n > t {
				n -= t
				this.pos += 2
			} else {
				this.A[this.pos] -= n
				return this.A[this.pos+1]
			}
		} else {
			this.pos += 2
		}
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
