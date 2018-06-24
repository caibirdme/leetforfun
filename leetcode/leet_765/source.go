package leet_765

func minSwapsCouples(row []int) int {
	pos := make(map[int]int)
	reverse := make(map[int]int)
	for idx, num := range row {
		pos[num] = idx
		reverse[idx] = num
	}
	count := 0
	for i := 0; i < len(row); i += 2 {
		posI := pos[i]
		posJ := pos[i+1]
		if posI+1 == posJ {
			if isOdd(posI) {
				count++
				swap(posI-1, posJ, pos, reverse)
			}
		} else if posI-1 == posJ {
			if isEven(posI) {
				count++
				swap(posJ, posI+1, pos, reverse)
			}
		} else {
			count++
			if isOdd(posI) {
				swap(posI-1, posJ, pos, reverse)
			} else {
				swap(posI+1, posJ, pos, reverse)
			}
		}
	}
	return count
}

func isEven(n int) bool {
	return !isOdd(n)
}

func isOdd(n int) bool {
	return n&1 == 1
}

func swap(i, j int, pos, reverse map[int]int) {
	numI := reverse[i]
	numJ := reverse[j]
	reverse[i] = numJ
	reverse[j] = numI
	pos[numI] = j
	pos[numJ] = i
}
