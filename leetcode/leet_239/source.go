package leet_239

var pos []int

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 || 0 == len(nums) {
		return nums
	}
	length := len(nums)
	pos = make([]int, length)
	queue := make(heap, 0, k)
	for i := 0; i < k; i++ {
		queue.Push(packItem(nums[i], i))
	}
	result := make([]int, 0, length-k+1)
	result = append(result, queue.Top().data)
	for i := k; i < length; i++ {
		invalidIdx := i - k
		queue.Remove(pos[invalidIdx])
		queue.Push(packItem(nums[i], i))
		result = append(result, queue.Top().data)
	}
	return result
}

func packItem(data, idx int) heapItem {
	return heapItem{
		idx, data,
	}
}

type heapItem struct {
	idx  int
	data int
}
type heap []heapItem

func (h heap) Top() heapItem {
	return h[0]
}

func (h *heap) Push(item heapItem) {
	length := len(*h)
	*h = append(*h, item)
	pos[item.idx] = length
	h.upModify(length)
}

func (h *heap) Remove(idx int) {
	length := len(*h)
	last := length - 1
	if last != idx {
		(*h)[idx] = (*h)[last]
		pos[(*h)[idx].idx] = idx
		*h = (*h)[:last]
		h.upModify(idx)
		h.downModify(idx)
	} else {
		*h = (*h)[:last]
	}
}

func (h heap) upModify(idx int) {
	for idx > 0 {
		parent := (idx - 1) >> 1
		if h[idx].data <= h[parent].data {
			break
		}
		pos[h[idx].idx] = parent
		pos[h[parent].idx] = idx
		h[idx], h[parent] = h[parent], h[idx]
		idx = parent
	}
}

func (h heap) downModify(idx int) {
	length := len(h) - 1
	for idx < length {
		lch := idx<<1 + 1
		if lch > length {
			break
		}
		rch := lch + 1
		t := lch
		if rch <= length && h[rch].data > h[lch].data {
			t = rch
		}
		if h[idx].data >= h[t].data {
			break
		}
		pos[h[idx].idx] = t
		pos[h[t].idx] = idx
		h[idx], h[t] = h[t], h[idx]
		idx = t
	}
}
