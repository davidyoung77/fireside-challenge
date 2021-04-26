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
	return Grid{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
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

		for _, row := range grid {

			if row[column] != 0 {
				continue
			}

			if turn%2 == 0 {
				row[column] = 1
			} else {
				row[column] = 2
			}

			break
		}

		winner = checkWin()
	}

	return winner, nil
}

func checkWin() int {
	for rowI, row := range grid {
		for colI, token := range row {
			if token == 0 {
				return 0
			}

			return checkRow(rowI, row, colI, token) |
				checkColumn(rowI, row, colI, token) |
				checkDiagonal(rowI, row, colI, token)
		}
	}

	return 0
}

func checkRow(rowI int, row []int, colI int, token int) int {
	if game.columns >= colI+3 &&
		token == row[colI+1] &&
		token == row[colI+2] &&
		token == row[colI+3] {
		return token
	}

	return 0
}

func checkColumn(rowI int, row []int, colI int, token int) int {
	if game.rows >= rowI+3 &&
		token == grid[rowI+1][colI] &&
		token == grid[rowI+2][colI] &&
		token == grid[rowI+3][colI] {
		return token
	}

	return 0
}

func checkDiagonal(rowI int, row []int, colI int, token int) int {
	if game.rows > rowI+3 &&
		game.columns > colI+3 &&
		token == grid[rowI+1][colI+1] &&
		token == grid[rowI+2][colI+2] &&
		token == grid[rowI+3][colI+3] {
		return token
	}

	if rowI-3 > 0 &&
		colI-3 > 0 &&
		token == grid[rowI+1][colI-1] &&
		token == grid[rowI+2][colI-2] &&
		token == grid[rowI+3][colI-3] {
		return token
	}

	return 0
}
