package main

import (
	"fmt"
)

func main() {

	board := [][]string{
		{"x", "o", "x"},
		{"o", "x", "x"},
		{"x", "x", "x"},
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

func checkRow(board [][]string) (bool, int) {
	for row := 0; row < 3; row++ {
		if board[row][0] == board[row][1] && board[row][1] == board[row][2] {
			return true, row
		}
	}
	return false, -1
}
