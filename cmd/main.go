package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ch40s1/tictactoe-go/cmd/game"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", cwd)

	_, err = os.Stat("./cmd/client/index.html")
	if os.IsNotExist(err) {
		log.Fatal("index.html does not exist in the client directory")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./cmd/client"))
	mux.Handle("/", fs)
	mux.HandleFunc("/game", game.ShowBoard)
	fmt.Println("listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
