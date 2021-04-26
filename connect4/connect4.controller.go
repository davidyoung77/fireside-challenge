package connect4

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	winner, err := addTokensToGrid(tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
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
