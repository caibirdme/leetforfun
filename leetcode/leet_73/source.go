package leet_73

func setZeroes(matrix [][]int) {
	const flag = 0x80000000
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = flag
				for t := 0; t < n; t++ {
					if matrix[i][t] != 0 && matrix[i][t] != flag {
						matrix[i][t] = flag
					}
				}
				for t := 0; t < m; t++ {
					if matrix[t][j] != 0 && matrix[t][j] != flag {
						matrix[t][j] = flag
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == flag {
				matrix[i][j] = 0
			}
		}
	}
}
