package leet_399

import (
	"sort"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	length := len(equations)
	points := make(map[string]struct{})
	for i := 0; i < length; i++ {
		val := values[i]
		x, y := equations[i][0], equations[i][1]
		points[x] = struct{}{}
		points[y] = struct{}{}
		if _, ok := graph[x]; !ok {
			graph[x] = map[string]float64{
				y: val,
			}
		} else {
			graph[x][y] = val
		}
		if _, ok := graph[y]; !ok {
			graph[y] = map[string]float64{
				x: 1 / val,
			}
		} else {
			graph[y][x] = 1 / val
		}
	}
	floyed(graph, points)
	queryNum := len(queries)
	result := make([]float64, queryNum)
	for i := 0; i < queryNum; i++ {
		x, y := queries[i][0], queries[i][1]
		_, ok1 := points[x]
		_, ok2 := points[y]
		if !ok1 || !ok2 {
			result[i] = -1.0
			continue
		}
		if x == y {
			result[i] = 1.0
			continue
		}
		val, ok := graph[x][y]
		if ok {
			result[i] = val
		} else {
			result[i] = -1.0
		}
	}
	return result
}

func floyed(graph map[string]map[string]float64, ponits map[string]struct{}) {
	var orderStr []string
	for k := range ponits {
		orderStr = append(orderStr, k)
	}
	sort.Strings(orderStr)
	for _, mid := range orderStr {
		for _, left := range orderStr {
			if left == mid {
				continue
			}
			for _, right := range orderStr {
				if mid == right || right == left {
					continue
				}
				if _, ok := graph[left][right]; ok {
					continue
				}
				val1, ok1 := graph[left][mid]
				val2, ok2 := graph[mid][right]
				if ok1 && ok2 {
					t := val1 * val2
					graph[left][right] = t
					graph[right][left] = 1 / t
					break
				}
			}
		}
	}
}
