package problema2

import (
	"fmt"
)

const N = 5

var (
	xMove = []int{2, 1, -1, -2, -2, -1, 1, 2}
	yMove = []int{1, 2, 2, 1, -1, -2, -2, -1}
)

func PrintSolution(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%2d ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func IsSafe(x, y int, board [N][N]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && board[x][y] == -1
}

func SolveKTUtil(x, y, movei int, board *[N][N]int) bool {
	if movei == N*N {
		return true
	}

	for k := 0; k < 8; k++ {
		nextX := x + xMove[k]
		nextY := y + yMove[k]
		if IsSafe(nextX, nextY, *board) {
			board[nextX][nextY] = movei

			if SolveKTUtil(nextX, nextY, movei+1, board) {
				return true
			} else {
				board[nextX][nextY] = -1
			}
		}
	}

	return false
}

func SolveKT() {
	var board [N][N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = -1
		}
	}

	board[0][0] = 0

	if !SolveKTUtil(0, 0, 1, &board) { // Custo: C1, Número de execuções: 1
		fmt.Println("Solution does not exist")
	} else {
		PrintSolution(board)
	}
}
