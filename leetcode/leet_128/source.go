package leet_128

func longestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ans := 1
	// length * 8 for avoid collision
	hash := NewHashTable(length << 3)
	for _, num := range nums {
		p, ok := hash.Get(num - 1)
		if ok {
			merge(p, hash)
			pair := *p
			pair.Grow(num)
			merge(p, hash)
			ans = max(ans, pair.Len())
			hash.Set(num, p)
			continue
		}
		p, ok = hash.Get(num + 1)
		if ok {
			merge(p, hash)
			pair := *p
			pair.Grow(num)
			merge(p, hash)
			ans = max(ans, pair.Len())
			hash.Set(num, p)

			continue
		}
		if _, ok = hash.Get(num); !ok {
			hash.Set(num, NewSegment(num))
		}
	}
	return ans
}

func merge(p **consecutiveSegment, hash *hashTable) {
	currentPair := *p
	leftBound := currentPair.Left()
	rightBound := currentPair.Right()
	leftPair, ok := hash.Get(leftBound - 1)
	if ok && currentPair.Merge(*leftPair) {
		*leftPair = currentPair
		merge(p, hash)
	}
	rightPair, ok := hash.Get(rightBound + 1)
	if ok && currentPair.Merge(*rightPair) {
		*rightPair = currentPair
		merge(p, hash)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type hashTable struct {
	bound int
	buf   [][]hashElem
}

func NewHashTable(bound int) *hashTable {
	return &hashTable{bound: bound, buf: make([][]hashElem, bound)}
}

func (ht *hashTable) getOffset(key int) int {
	if key < 0 {
		key = -key
	}
	return key % ht.bound
}

func (ht *hashTable) Get(key int) (**consecutiveSegment, bool) {
	offset := ht.getOffset(key)
	length := len(ht.buf[offset])
	if 0 == length {
		return nil, false
	}
	for i := 0; i < length; i++ {
		if ht.buf[offset][i].key == key {
			return ht.buf[offset][i].p, true
		}
	}
	return nil, false
}

func (ht *hashTable) Set(key int, data **consecutiveSegment) {
	offset := ht.getOffset(key)
	length := len(ht.buf[offset])
	elem := hashElem{key, data}
	if 0 == length {
		ht.buf[offset] = []hashElem{elem}
		return
	}
	for i := 0; i < length; i++ {
		if ht.buf[offset][i].key == key {
			ht.buf[offset][i].p = data
			return
		}
	}
	ht.buf[offset] = append(ht.buf[offset], elem)
}

type hashElem struct {
	key int
	p   **consecutiveSegment
}

type consecutiveSegment struct {
	left  int
	right int
}

func NewSegment(num int) **consecutiveSegment {
	p := new(*consecutiveSegment)
	*p = &consecutiveSegment{left: num, right: num}
	return p
}

func (c *consecutiveSegment) Merge(pair *consecutiveSegment) bool {
	if c.left-1 <= pair.right && c.left > pair.left {
		c.left = pair.left
	} else if c.right+1 >= pair.left && c.right < pair.right {
		c.right = pair.right
	} else {
		return false
	}
	return true
}

func (c *consecutiveSegment) Len() int {
	return c.right - c.left + 1
}

func (c *consecutiveSegment) Left() int {
	return c.left
}

func (c *consecutiveSegment) Right() int {
	return c.right
}

func (c *consecutiveSegment) InSegment(num int) bool {
	return num >= c.left && num <= c.right
}

func (c *consecutiveSegment) Grow(num int) {
	if num == c.left-1 {
		c.left = num
	} else if num == c.right+1 {
		c.right = num
	}
}
