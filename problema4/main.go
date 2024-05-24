package problema4

import (
	"fmt"
	"math"
)

// Cidade representa uma cidade com coordenadas X e Y
type Cidade struct {
	X, Y float64
}

// Função para calcular a distância euclidiana entre duas cidades
func distancia(c1, c2 Cidade) float64 {
	return math.Sqrt(math.Pow(c2.X-c1.X, 2) + math.Pow(c2.Y-c1.Y, 2))
}

// Função para calcular a distância total de uma rota
func distanciaTotal(cidades []Cidade, permutacao []int) float64 {
	total := 0.0
	for i := 0; i < len(permutacao)-1; i++ {
		total += distancia(cidades[permutacao[i]], cidades[permutacao[i+1]])
	}
	// Adicionar a volta para a cidade inicial
	total += distancia(cidades[permutacao[len(permutacao)-1]], cidades[permutacao[0]])
	return total
}

// Função para gerar todas as permutações de um slice de inteiros
func permutacoes(slice []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}

	helper(slice, len(slice))
	return res
}

// Função principal para resolver o problema do caixeiro viajante
func RunExample() {
	cidades := []Cidade{
		{X: 0, Y: 0},
		{X: 4, Y: 3},
		{X: 10, Y: 2},
		{X: 4, Y: 9},
		{X: 4, Y: 42},
		{X: 15, Y: 2},
		{X: 6, Y: 12},
		{X: 5, Y: 5},
		{X: 18, Y: 5},
		{X: 9, Y: 3},
	}

	n := len(cidades)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	todasPermutacoes := permutacoes(indices) // Gerar todas as permutações possíveis
	menorDistancia := math.Inf(1)            // Inicializar com infinito
	melhorRota := []int{}                    // Inicializar com slice vazio
	for _, perm := range todasPermutacoes {  // Para cada permutação
		dist := distanciaTotal(cidades, perm) // Calcular a distância total
		if dist < menorDistancia {            // Se a distância for menor que a menor distância atual
			menorDistancia = dist
			melhorRota = perm
		}
	}

	fmt.Println("Menor distância:", menorDistancia)
	fmt.Println("Melhor rota:", melhorRota)
}

func RunExample2() {
	cidades := []Cidade{
		{X: 0, Y: 0},
		{X: 4, Y: 3},
		{X: 10, Y: 2},
		{X: 4, Y: 9},
		{X: 4, Y: 42},
		{X: 15, Y: 2},
		{X: 6, Y: 12},
		{X: 5, Y: 5},
		{X: 18, Y: 5},
		{X: 9, Y: 3},
	}

	n := len(cidades)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	todasPermutacoes := permutacoes(indices) // Gerar todas as permutações possíveis
	menorDistancia := math.Inf(1)            // Inicializar com infinito
	melhorRota := []int{}                    // Inicializar com slice vazio

	for _, perm := range todasPermutacoes { // Para cada permutação
		dist := distanciaTotal(cidades, perm) // Calcular a distância total
		if dist < menorDistancia {            // Se a distância for menor que a menor distância atual
			menorDistancia = dist
			melhorRota = perm
		}
	}

	fmt.Println("Melhor rota:")
	for i := 0; i < len(melhorRota); i++ {
		cidade1 := cidades[melhorRota[i]]
		var cidade2 Cidade
		if i == len(melhorRota)-1 {
			cidade2 = cidades[melhorRota[0]] // Volta para a cidade inicial
		} else {
			cidade2 = cidades[melhorRota[i+1]]
		}
		distanciaEntreCidades := distancia(cidade1, cidade2)
		if i == len(melhorRota)-1 {
			fmt.Printf("Cidade %d (%.2f, %.2f) -> Cidade %d (%.2f, %.2f) [Distância: %.2f]\n",
				melhorRota[i], cidade1.X, cidade1.Y, melhorRota[0], cidade2.X, cidade2.Y, distanciaEntreCidades)
		} else {
			fmt.Printf("Cidade %d (%.2f, %.2f) -> Cidade %d (%.2f, %.2f) [Distância: %.2f]\n",
				melhorRota[i], cidade1.X, cidade1.Y, melhorRota[i+1]%len(melhorRota), cidade2.X, cidade2.Y, distanciaEntreCidades)
		}
	}
	fmt.Printf("Distância total: %.2f\n", menorDistancia)
}
