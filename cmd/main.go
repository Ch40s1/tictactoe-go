package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ch40s1/tictactoe-go/cmd/game"
)

// func handles all routes
func main() {
	// gets working directory and handles error if doesnt return nil
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", cwd)
	//  checks for index.html file
	_, err = os.Stat("./cmd/client/index.html")
	// handle error
	if os.IsNotExist(err) {
		log.Fatal("index.html does not exist in the client directory")
	}

	// creates mux route handler
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./cmd/client"))
	// home page
	mux.Handle("/", fs)
	// get request to show board
	mux.HandleFunc("/board", game.ShowBoard)
	// post request to get user data
	mux.HandleFunc("/turn", game.GetUserMove)
	fmt.Println("listening on http://localhost:3000/")
	// uses port 3000 and route handler
	log.Fatal(http.ListenAndServe(":3000", mux))
}
