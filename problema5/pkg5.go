package problema5

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Ator struct {
	Name string `json:"name"`
}

// Função para ler o CSV e construir o grafo de atores
func lerCSV(caminho string) map[string][]string {
	file, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	grafo := make(map[string][]string)
	for _, record := range records[1:] {
		cast := record[2]
		var atores []Ator
		err := json.Unmarshal([]byte(cast), &atores)
		if err != nil {
			log.Fatal(err)
		}
		for _, ator := range atores {
			for _, outroAtor := range atores {
				if ator.Name != outroAtor.Name {
					grafo[ator.Name] = append(grafo[ator.Name], outroAtor.Name)
				}
			}
		}
	}
	return grafo
}

// Função para calcular o Número de Bacon usando BFS
func numeroDeBacon(grafo map[string][]string, origem, destino string) int {
	if origem == destino {
		return 0
	}

	visitado := make(map[string]bool)
	fila := []string{origem}
	distancia := map[string]int{origem: 0}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		for _, vizinho := range grafo[atual] {
			if !visitado[vizinho] {
				distancia[vizinho] = distancia[atual] + 1
				if vizinho == destino {
					return distancia[vizinho]
				}
				visitado[vizinho] = true
				fila = append(fila, vizinho)
			}
		}
	}

	return -1 // Retorna -1 se não houver caminho entre origem e destino
}

func Problema5() {
	caminhoCSV := "problema5/tmdb_5000_credits.csv"
	grafo := lerCSV(caminhoCSV)
	fmt.Printf("Grafo de atores construído com %d atores\n", len(grafo))
	origem := "Kevin Bacon"
	destino := "Shane Carruth"
	numeroBacon := numeroDeBacon(grafo, origem, destino)

	if numeroBacon != -1 {
		fmt.Printf("O Número de Bacon entre %s e %s é: %d\n", origem, destino, numeroBacon)
	} else {
		fmt.Printf("Não há caminho entre %s e %s\n", origem, destino)
	}
}
