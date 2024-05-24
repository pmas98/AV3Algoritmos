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
	menorRota := 0

	fmt.Println("Rota percorrida:")
	for i, cidade := range rota {
		if i == len(rota)-1 {
			menorRota += int(distancia(cidades[rota[i]], cidades[0]))
		} else {
			menorRota += int(distancia(cidades[rota[i]], cidades[rota[i+1]]))
		}
		fmt.Printf("Cidade %d: (%.2f, %.2f)\n", cidade, cidades[cidade].X, cidades[cidade].Y)
	}
	fmt.Printf("Largura percorrida " + fmt.Sprintf("%d", menorRota))

}
