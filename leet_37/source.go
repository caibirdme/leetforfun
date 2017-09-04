package leet_37

type fillType byte

const (
	_ fillType = iota
	cFillCol
	cFillRow
	cFillSquare
)

type sudokuState struct {
	flagCol, flagRow, flagSquare    [9][10]bool
	countCol, countRow, countSquare [9]byte
	grid                            [9][9]byte
	counter                         int
}

func (state *sudokuState) fillNumber(i, j, number int) bool {
	if !state.canFill(i, j, number) {
		return false
	}
	state.update(i, j, number)
	return true
}

func (state *sudokuState) finished() bool {
	return state.counter == 81
}

func (state *sudokuState) canFill(i, j, number int) bool {
	if state.grid[i][j] != 0 {
		return false
	}
	if number == 0 {
		return true
	}
	return !(state.flagCol[i][number] || state.flagRow[j][number] || state.flagSquare[state.getSquare(i, j)][number])
}

func (state *sudokuState) update(i, j, number int) {
	state.flagCol[i][number] = true
	state.flagRow[j][number] = true
	square := state.getSquare(i, j)
	state.flagSquare[square][number] = true
	state.countCol[i]++
	state.countRow[j]++
	state.countSquare[square]++
	state.counter++
	state.grid[i][j] = byte(number)
}

func (state *sudokuState) undo(i, j, number int) {
	state.flagCol[i][number] = false
	state.flagRow[j][number] = false
	square := state.getSquare(i, j)
	state.flagSquare[square][number] = false
	state.countCol[i]--
	state.countRow[j]--
	state.countSquare[square]--
	state.counter--
	state.grid[i][j] = 0
}

func (state *sudokuState) getSquare(i, j int) int {
	return i/3*3 + j/3
}

func (state *sudokuState) fillStrategy() (fillType, int) {
	var max byte
	which := 0
	var fType fillType
	for i := 0; i < 9; i++ {
		if state.countCol[i] > max && state.countCol[i] < 9 {
			max = state.countCol[i]
			which = i
		}
	}
	for i := 0; i < 9; i++ {
		if state.countRow[i] > max && state.countRow[i] < 9 {
			max = state.countRow[i]
			which = i
			fType = cFillRow
		}
	}
	for i := 0; i < 9; i++ {
		if state.countSquare[i] > max && state.countSquare[i] < 9 {
			max = state.countSquare[i]
			which = i
			fType = cFillSquare
		}
	}
	return fType, which
}

func newSudokuState(board [][]byte) *sudokuState {
	s := &sudokuState{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			s.update(i, j, int(board[i][j]-'0'))
		}
	}
	return s
}

func solveSudoku(board [][]byte) {
	state := newSudokuState(board)
	solve(state)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = state.grid[i][j] + '0'
		}
	}
}

func solve(state *sudokuState) {
	if state.finished() {
		return
	}
	fType, which := state.fillStrategy()
	try2Fill(fType, which, state)
}

func try2Fill(fType fillType, which int, state *sudokuState) {
	fill := fillCol
	switch fType {
	case cFillRow:
		fill = fillRow
	case cFillSquare:
		fill = fillSquare
	}
	fill(which, state)
}

func fillRow(which int, state *sudokuState) {
	var row int
	for i := 0; i < 9; i++ {
		if state.canFill(i, which, 0) {
			row = i
			break
		}
	}
	for number := 1; number <= 9; number++ {
		if state.fillNumber(row, which, number) {
			solve(state)
			if !state.finished() {
				state.undo(row, which, number)
			}
		}
	}
}

func fillCol(which int, state *sudokuState) {
	var col int
	for i := 0; i < 9; i++ {
		if state.canFill(which, i, 0) {
			col = i
			break
		}
	}
	for number := 1; number <= 9; number++ {
		if state.fillNumber(which, col, number) {
			solve(state)
			if !state.finished() {
				state.undo(which, col, number)
			}
		}
	}
}

func fillSquare(which int, state *sudokuState) {
	startI, startJ := getSquareStart(which)
	endI, endJ := startI+3, startJ+3
	var row, col int
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
			if state.canFill(i, j, 0) {
				row = i
				col = j
				break
			}
		}
	}
	for number := 1; number <= 9; number++ {
		if state.fillNumber(row, col, number) {
			solve(state)
			if !state.finished() {
				state.undo(row, col, number)
			}
		}
	}
}

func getSquareStart(which int) (int, int) {
	switch which {
	case 0:
		return 0, 0
	case 1:
		return 0, 3
	case 2:
		return 0, 6
	case 3:
		return 3, 0
	case 4:
		return 3, 3
	case 5:
		return 3, 6
	case 6:
		return 6, 0
	case 7:
		return 6, 3
	default:
		return 6, 6
	}
}
