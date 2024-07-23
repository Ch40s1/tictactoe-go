package game

import (
	"encoding/json"
	"net/http"
)

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "showing board from api"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func game() {

// 	board := [][]string{
// 		{"x", "o", "x"},
// 		{"o", "x", "x"},
// 		{"x", "o", "o"},
// 	}
// 	// create game state
// 	playing := true
// 	// create stuct with player info for now use hardcode

// 	for playing {
// 		var i int
// 		fmt.Println(checkRow(board))
// 		fmt.Println(checkCol(board))
// 		fmt.Println(checkDia(board))
// 		fmt.Println("game in progress")
// 		fmt.Println("press 1 to exit")
// 		fmt.Scan(&i)
// 		defer fmt.Println("exiting game")
// 		if i == 1 {
// 			return
// 		} else {
// 			fmt.Println("press 1 to exit or wait")
// 		}
// 	}
// }

// func checkRow(board [][]string) (bool, int) {
// 	for row := 0; row < 3; row++ {
// 		if board[row][0] == board[row][1] && board[row][1] == board[row][2] {
// 			return true, row
// 		}
// 	}
// 	return false, -1
// }

// func checkCol(board [][]string) (bool, int) {
// 	for col := 0; col < 3; col++ {
// 		if board[0][col] == board[1][col] && board[1][col] == board[2][col] {
// 			return true, col
// 		}
// 	}
// 	return false, -1
// }

// func checkDia(board [][]string) (bool, int) {
// 	//diagonal one
// 	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
// 		return true, 1
// 	}
// 	//diagonal two
// 	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
// 		return true, 2
// 	}
// 	return false, -1
// }
