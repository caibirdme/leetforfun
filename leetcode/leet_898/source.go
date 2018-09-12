package leet_898

func subarrayBitwiseORs(A []int) int {
	length := len(A)
	if length == 1 {
		return 1
	}
	result := map[int]struct{}{
		A[0]: struct{}{},
	}
	preAns := map[int]struct{}{
		A[0]: struct{}{},
	}
	for i := 1; i < length; i++ {
		newAns := make(map[int]struct{})
		for k := range preAns {
			t := k | A[i]
			if _, ok := result[t]; !ok {
				result[t] = struct{}{}
			}
			newAns[t] = struct{}{}
		}
		newAns[A[i]] = struct{}{}
		if _, ok := result[A[i]]; !ok {
			result[A[i]] = struct{}{}
		}
		preAns = newAns
	}
	return len(result)
}
