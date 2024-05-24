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

// IsSafe verifica se a posição (x, y) é válida para o cavalo
func IsSafe(x, y int, board [N][N]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && board[x][y] == -1
}

func SolveKTUtil(x, y, movei int, board *[N][N]int) bool {
	if movei == N*N { // Se todas as posições foram preenchidas, retorna true
		return true
	}

	for k := 0; k < 8; k++ { // Tenta todos os movimentos possíveis
		nextX := x + xMove[k]
		nextY := y + yMove[k]
		if IsSafe(nextX, nextY, *board) { // Se o movimento é válido
			board[nextX][nextY] = movei                    // Marca a posição
			if SolveKTUtil(nextX, nextY, movei+1, board) { // Chama recursivamente
				return true
			} else {
				board[nextX][nextY] = -1
			}
		}
	}

	return false
}

// SolveKT solves the Knight's Tour problem using backtracking
func SolveKT() {
	var board [N][N]int

	// Inicializando o tabuleiro
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = -1
		}
	}

	// Posição inicial do cavalo
	board[0][0] = 0

	if !SolveKTUtil(0, 0, 1, &board) {
		fmt.Println("Solution does not exist")
	} else {
		PrintSolution(board)
	}
}
