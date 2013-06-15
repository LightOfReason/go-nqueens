package nqueens

var queens = 0

//******************************************************************************
// Parallel Solver
//******************************************************************************

func SolveParallel(nQueens int) uint64 {
	queens = nQueens
	solutions := make(chan uint64)

	for startColumn := 0; startColumn < queens; startColumn++ {
		gameField := make([]int, queens)
		gameField[0] = startColumn

		go func() {
			solutions <- sequential(gameField, 1)
		}()
	}

	result := uint64(0)
	for startColumn := 0; startColumn < queens; startColumn++ {
		result += <-solutions
	}

	return result
}

//******************************************************************************
// Sequential Solver
//******************************************************************************

func SolveSequential(nQueens int) uint64 {
	queens = nQueens
	gameField := make([]int, queens)
	return sequential(gameField, 0)
}

func sequential(gameField []int, currentRow int) uint64 {
	if currentRow >= queens {
		return 1 //solution found
	}

	solutions := uint64(0)
	for column := 0; column < queens; column++ {
		gameField[currentRow] = column

		if !hasCollison(gameField, currentRow) {
			solutions += sequential(gameField, currentRow+1)
		}
	}

	return solutions
}

func hasCollison(gameField []int, currentRow int) bool {
	for row := 0; row < currentRow; row++ {
		difColumn := abs(gameField[row] - gameField[currentRow])
		difRow := abs(row - currentRow)

		if difColumn == 0 || difColumn == difRow {
			return true
		}
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
