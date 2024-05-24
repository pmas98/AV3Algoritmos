package problema1

import (
	"fmt"
	"strings"
)

const N = 8

func PrintBoard(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 2*N))
}

func IsSafe(board [N][N]int, row, col int) bool {

	for i := 0; i < col; i++ { // Verifique se há uma rainha na mesma linha
		if board[row][i] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 { // Verifique se há uma rainha na diagonal esquerda superior
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row, col; i < N && j >= 0; i, j = i+1, j-1 { // Verifique se há uma rainha na diagonal esquerda inferior
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func SolveNQUtil(board [N][N]int, col int) bool {
	if col >= N { // Se todas as rainhas foram colocadas, acabe o programa
		return true
	}

	for i := 0; i < N; i++ { // Itere sobre cada linha da matrix
		if IsSafe(board, i, col) { // Verifique se a rainha pode ser colocada na posição [i][col]
			board[i][col] = 1 // Se sim, coloque a rainha

			if SolveNQUtil(board, col+1) { // Duas rainhas não podem estar na mesma coluna, então chame a função recursivamente para a próxima coluna
				PrintBoard(board)
				return true
			}

			board[i][col] = 0 // Se a rainha não puder ser colocada, remova a rainha
		}
	}

	return false
}

func SolveNQ() {
	var board [N][N]int                      // Inicializando uma matrix com a constante N
	var solutionable = SolveNQUtil(board, 0) //Chamando a função para encontrar uma solução
	if !solutionable {                       //
		fmt.Println("Solution does not exist")
		return
	}
}
