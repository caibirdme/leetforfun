package leet_37

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	var testData = [][][]byte{
		[][]byte{
			[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		[][]byte{
			[]byte("..9748..."),
			[]byte("7........"),
			[]byte(".2.1.9..."),
			[]byte("..7...24."),
			[]byte(".64.1.59."),
			[]byte(".98...3.."),
			[]byte("...8.3.2."),
			[]byte("........6"),
			[]byte("...2759.."),
		},
	}
	for _, grip := range testData {
		board := copyBoard(grip, 9, 9)
		solveSudoku(board)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%c", board[i][j])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n\n")
	}
}

func copyBoard(src [][]byte, one, two int) [][]byte {
	newSlice := make([][]byte, one)
	for i := 0; i < one; i++ {
		newSlice[i] = make([]byte, two)
		for j := 0; j < two; j++ {
			newSlice[i][j] = src[i][j]
		}
	}
	return newSlice
}
