package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := make([]int, 8)
	solutions := [][]int{}

	solveQueens(board, 0, &solutions)

	for _, solution := range solutions {
		for _, pos := range solution {
			z01.PrintRune(rune(pos + '0'))
		}
		z01.PrintRune('\n')
	}
}

func solveQueens(board []int, col int, solutions *[][]int) {
	if col == 8 {
		solution := make([]int, 8)
		for i := 0; i < 8; i++ {
			solution[i] = board[i] + 1
		}
		*solutions = append(*solutions, solution)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solveQueens(board, col+1, solutions)
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row {
			return false
		}

		if board[i]-i == row-col {
			return false
		}

		if board[i]+i == row+col {
			return false
		}
	}
	return true
}
