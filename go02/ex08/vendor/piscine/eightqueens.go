package piscine

import "ft"

const boardLen = 8

func printQueen(queen []int) {
	for y := 0; y < boardLen; y++ {
		ft.PrintRune(rune(queen[y]+1) + '0')
	}
	ft.PrintRune('\n')
}

func setInBoard(board [][boardLen]int, y int, x int, updown int) {
	if x >= 0 && x < boardLen && y >= 0 && y < boardLen {
		board[y][x] += updown
	}
}

func updateBoard(board [][boardLen]int, y int, x int, updown int) {
	setInBoard(board, y, x, updown)
	for i := 0; i < boardLen; i++ {
		setInBoard(board, y-i, x, updown)
		setInBoard(board, y, x-i, updown)
		setInBoard(board, y+i, x, updown)
		setInBoard(board, y, x+i, updown)
		setInBoard(board, y+i, x+i, updown)
		setInBoard(board, y-i, x-i, updown)
		setInBoard(board, y+i, x-i, updown)
		setInBoard(board, y-i, x+i, updown)
	}
}

func setQueenInRow(board [][boardLen]int, queen []int, y int) {
	if y == boardLen {
		printQueen(queen)
		return
	}
	for x := 0; x < boardLen; x++ {
		if board[y][x] == 0 {
			queen[y] = x
			updateBoard(board, y, x, 1)
			setQueenInRow(board, queen, y+1)
			updateBoard(board, y, x, -1)
		}
	}
}

func EightQueens() {
	var board [boardLen][boardLen]int
	var queen [boardLen]int

	setQueenInRow(board[:], queen[:], 0)
}
