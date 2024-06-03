package problema4

import (
	"fmt"
	"math"
)

// Encontra a cidade não visitada mais próxima
func encontrarMaisProxima(cidades []Cidade, atual int, visitado []bool) int {
	maisProxima := -1
	menorDistancia := math.MaxFloat64
	for i := range cidades {
		if !visitado[i] {
			dist := distancia(cidades[atual], cidades[i])
			if dist < menorDistancia {
				menorDistancia = dist
				maisProxima = i
			}
		}
	}
	return maisProxima
}

// Algoritmo guloso para o problema do caixeiro-viajante
func caixeiroViajante(cidades []Cidade) []int {
	n := len(cidades)
	if n == 0 {
		return nil
	}

	visitado := make([]bool, n)
	rota := make([]int, n)
	atual := 0
	visitado[atual] = true
	rota[0] = atual

	for i := 1; i < n; i++ {
		proxima := encontrarMaisProxima(cidades, atual, visitado)
		if proxima == -1 {
			break
		}
		visitado[proxima] = true
		rota[i] = proxima
		atual = proxima
	}

	return rota
}

func Greedy() {

	// Exemplo de 10 cidades com coordenadas x e y
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

	rota := caixeiroViajante(cidades)
	menorRota := 0.0

	fmt.Println("Rota percorrida:")
	for i := 0; i < len(rota); i++ {
		if i == len(rota)-1 {
			dist := distancia(cidades[rota[i]], cidades[rota[0]])
			menorRota += dist
			fmt.Printf("Cidade %d: (%.2f, %.2f) -> Cidade %d: (%.2f, %.2f), Distância: %.2f\n",
				rota[i], cidades[rota[i]].X, cidades[rota[i]].Y, rota[0], cidades[rota[0]].X, cidades[rota[0]].Y, dist)
		} else {
			dist := distancia(cidades[rota[i]], cidades[rota[i+1]])
			menorRota += dist
			fmt.Printf("Cidade %d: (%.2f, %.2f) -> Cidade %d: (%.2f, %.2f), Distância: %.2f\n",
				rota[i], cidades[rota[i]].X, cidades[rota[i]].Y, rota[i+1], cidades[rota[i+1]].X, cidades[rota[i+1]].Y, dist)
		}
	}
	fmt.Printf("Largura percorrida: %.2f\n", menorRota)
}
