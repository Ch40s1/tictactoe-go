package main

import "testing"

func TestCheckRow(t *testing.T) {
	tests := []struct {
		board       [][]string
		expectedRow int
		expectedOk  bool
	}{
		// Test cases covering different scenarios
		{
			board: [][]string{
				{"x", "o", "x"},
				{"o", "x", "x"},
				{"x", "x", "x"},
			},
			expectedRow: 2,
			expectedOk:  true,
		},
		{
			board: [][]string{
				{"o", "o", "x"},
				{"o", "x", "x"},
				{"x", "o", "o"},
			},
			expectedRow: -1,
			expectedOk:  false,
		},
		// Add more test cases as needed to cover edge cases
	}

	for _, tc := range tests {
		ok, row := checkRow(tc.board)
		if ok != tc.expectedOk {
			t.Errorf("Expected checkRow(%v) to return %v, but got %v", tc.board, tc.expectedOk, ok)
		}
		if row != tc.expectedRow {
			t.Errorf("Expected row %d to be winning, but got %d", tc.expectedRow, row)
		}
	}
}
