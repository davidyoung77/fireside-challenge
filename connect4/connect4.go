package connect4

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var game = Game{
	columns: 7,
	rows:    6,
}
var grid = newGrid()

func Connect4Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/connect4" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var tokens Tokens
	grid = newGrid()

	err := json.NewDecoder(r.Body).Decode(&tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(tokens) > 42 {
		fmt.Fprintf(w, "Too many tokens")

		return
	}

	winner := 0

	for turn, column := range tokens {
		if winner != 0 {
			break
		}

		err := addTokenToGrid(column, turn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		winner = checkWin()
	}

	if winner == 0 && len(tokens) < 42 {
		fmt.Fprintf(w, "Game not over, add another token")

		return
	}

	if winner == 0 {
		fmt.Fprintf(w, "DRAW")

		return
	}

	fmt.Fprintf(w, fmt.Sprintf("WINNER: Player %d", winner))
}

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

func addTokenToGrid(column int, turn int) error {
	for _, row := range grid {
		if column >= 7 {
			return errors.New(fmt.Sprintf("Column %v doesn't exist, add tokens in columns 0 - 6", column))
		}

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

	return nil
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
