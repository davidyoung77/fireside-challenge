package connect4

import (
	"errors"
	"fmt"
)

var game = Game{
	columns: 7,
	rows:    6,
}
var grid = newGrid()

func newGrid() Grid {
	var _grid = make([][]int, game.rows)
	for i := range _grid {
		_grid[i] = make([]int, game.columns)
	}
	return _grid
}

func addTokensToGrid(tokens Tokens) (int, error) {
	winner := 0

	for turn, column := range tokens {
		if winner != 0 {
			break
		}

		if column >= 7 {
			return winner, errors.New(fmt.Sprintf("Column %v doesn't exist, add tokens in columns 0 - 6", column))
		}

		rowI := getRow(column)

		if turn%2 == 0 {
			grid[rowI][column] = 1
		} else {
			grid[rowI][column] = 2
		}

		winner = checkWin(rowI, column)
	}

	return winner, nil
}

func getRow(colI int) int {
	for i, row := range grid {
		if row[colI] == 0 {
			return i
		}
	}

	return 0
}

func checkWin(rowI int, colI int) int {
	return checkRow(rowI, grid[rowI], colI, grid[rowI][colI]) |
		checkColumn(rowI, grid[rowI], colI, grid[rowI][colI]) |
		checkDiagonal(rowI, grid[rowI], colI, grid[rowI][colI])
}

func checkRow(rowI int, row []int, colI int, token int) int {
	if 0 <= colI-3 &&
		token == row[colI-1] &&
		token == row[colI-2] &&
		token == row[colI-3] {
		return token
	}

	return 0
}

func checkColumn(rowI int, row []int, colI int, token int) int {
	if 0 <= rowI-3 &&
		token == grid[rowI-1][colI] &&
		token == grid[rowI-2][colI] &&
		token == grid[rowI-3][colI] {
		return token
	}

	return 0
}

func checkDiagonal(rowI int, row []int, colI int, token int) int {
	if 0 <= rowI-3 &&
		0 <= colI-3 &&
		token == grid[rowI-1][colI-1] &&
		token == grid[rowI-2][colI-2] &&
		token == grid[rowI-3][colI-3] {
		return token
	}

	if 0 <= rowI-3 &&
		game.columns > colI+3 &&
		token == grid[rowI-1][colI+1] &&
		token == grid[rowI-2][colI+2] &&
		token == grid[rowI-3][colI+3] {
		return token
	}

	return 0
}
