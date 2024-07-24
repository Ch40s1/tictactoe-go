package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// create 2d slice with empty values
var board = [][]string{
	{"", "", ""},
	{"", "", ""},
	{"", "", ""},
}

func updateBoard(data map[string]string) {
	inputCords := data["cord"]
	indices := strings.Split(inputCords, ",")
	row, err1 := strconv.Atoi(indices[0])
	col, err2 := strconv.Atoi(indices[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting indices to integers")
		return
	}
	// Update the board at the given indices
	board[row][col] = data["symbol"]
}

func GetUserMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received: %+v\n", data)
	updateBoard(data)

	response := map[string]string{"status": "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"0,0": board[0][0],
		"0,1": board[0][1],
		"0,2": board[0][2],
		"1,0": board[1][0],
		"1,1": board[1][1],
		"1,2": board[1][2],
		"2,0": board[2][0],
		"2,1": board[2][1],
		"2,2": board[2][2],
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
