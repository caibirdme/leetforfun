package leet_233

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigitOne(t *testing.T) {
	var testData = []int{
		1, 8, 10, 19, 20, 25, 57, 99, 199, 599, 2000, 7865, 23434543, 1410065408,
	}
	ass := assert.New(t)
	for _, page := range testData {
		ass.Equal(cmp(page), countDigitOne(page), "%d", page)
	}
}

func cmp(k int) int {
	count := 1
	for i := 10; i <= k; i++ {
		t := i
		for t > 0 {
			q := t % 10
			if q == 1 {
				count++
			}
			t /= 10
		}
	}
	return count
}
