package main

import (
	"fmt"
	"log"
	"net/http"

	"go-app.com/fireside-challenge/connect4"
)

func main() {
	fileServer := http.FileServer(http.Dir("./connect4/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/connect4", connect4.Connect4Handler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
