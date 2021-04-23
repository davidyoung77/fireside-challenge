package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-app.com/fireside-challenge/connect4"
)

func testRequest(t *testing.T, body string) (string, int) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/connect4", bytes.NewBufferString(body))
	if err != nil {
		t.Errorf("%d\n", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(connect4.Connect4Handler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	return rr.Body.String(), rr.Code
}

func TestGameNotComplete(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0, 1, 1, 2, 3, 2, 2, 3, 3, 4, 4]`)
	expected := `Game not over, add another token`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestDiagonalWinPlayer1(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0, 1, 1, 2, 3, 2, 2, 3, 3, 4, 3]`)
	expected := `WINNER: Player 1`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestVerticalWinPlayer1(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0, 1, 0, 2, 0, 2, 0]`)
	expected := `WINNER: Player 1`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestHorizontalWinPlayer1(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0, 4, 1, 4, 2, 5, 3]`)
	expected := `WINNER: Player 1`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestDiagonalWinPlayer2(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[6, 0, 1, 1, 2, 3, 2, 2, 3, 3, 4, 3]`)
	expected := `WINNER: Player 2`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestVerticalWinPlayer2(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[6, 0, 1, 0, 2, 0, 2, 0]`)
	expected := `WINNER: Player 2`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestHorizontalWinPlayer2(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[6, 0, 4, 1, 4, 2, 5, 3]`)
	expected := `WINNER: Player 2`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestDraw(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0,1,2,3,4,5,6,6,5,4,3,2,1,0,2,2,2,2,2,2,2,0,1,2,3,4,5,6,6,5,4,3,2,1,0,0,1,2,3,4,5,6]`)
	expected := `DRAW`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestToManyTokens(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0,1,2,3,4,5,6,6,5,4,3,2,1,0,2,2,2,2,2,2,2,0,1,2,3,4,5,6,6,5,4,3,2,1,0,0,1,2,3,4,5,6,6]`)
	expected := `Too many tokens`

	if httpStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusOK)
	}

	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestTokenNonExistantColumns(t *testing.T) {
	responseBody, httpStatus := testRequest(t, `[0,1,2,3,4,5,6,6,4,4,3,2,1,0,2,2,2,2,2,2,2,0,1,2,3,4,5,6,6,5,4,3,2,1,0,0,7,2,3,4,5,6]`)
	expected := `doesn't exist, add tokens in columns 0 - 6`

	if httpStatus != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %d want %d\n", httpStatus, http.StatusBadRequest)
	}

	if !strings.Contains(responseBody, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}
