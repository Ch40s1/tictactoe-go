package main

import (
	"fmt"
)

func main() {

	board := [][]string{
		{"x", "o", "x"},
		{"o", "x", "o"},
		{"x", "o", "x"},
	}
	// create game state
	playing := true
	// create stuct with player info for now use hardcode

	for playing {
		var i int
		fmt.Println(checkRow(board))
		fmt.Println("game in progress")
		fmt.Println("press 1 to exit")
		fmt.Scan(&i)
		defer fmt.Println("exiting game")
		if i == 1 {
			return
		} else {
			fmt.Println("press 1 to exit or wait")
		}
	}
}

func checkRow(board [][]string) bool {
	// check row
	counter := 0
	suit := board[0][0]
	for _, row := range board[0] {
		if row == suit {
			counter++
		}
	}
	// if three of the same kind
	if counter == 3 {
		return true
	}
	// set type as winner and return bool
	return false
}
