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

func lerCSV(caminho string) map[string][]string {
	file, err := os.Open(caminho) // Custo: C1, Número de execuções: 1
	if err != nil {               // Custo: C2, Número de execuções: 1
		log.Fatal(err) // Custo: C3, Número de execuções: 1
	}
	defer file.Close() // Custo: C4, Número de execuções: 1

	reader := csv.NewReader(file)    // Custo: C5, Número de execuções: 1
	records, err := reader.ReadAll() // Custo: C6, Número de execuções: 1
	if err != nil {                  // Custo: C7, Número de execuções: 1
		log.Fatal(err) // Custo: C8, Número de execuções: 1
	}

	grafo := make(map[string][]string)   // Custo: C9, Número de execuções: 1
	for _, record := range records[1:] { // Custo: C10, Número de execuções: 1
		cast := record[2]                            // Custo: C11, Número de execuções: n
		var atores []Ator                            // Custo: C12, Número de execuções: n
		err := json.Unmarshal([]byte(cast), &atores) // Custo: C13, Número de execuções: 1
		if err != nil {                              // Custo: C14, Número de execuções: 1
			log.Fatal(err) // Custo: C15, Número de execuções: 1
		}
		for _, ator := range atores { // Custo: C16, Número de execuções: 1
			for _, outroAtor := range atores { // Custo: C17, Número de execuções: 1
				if ator.Name != outroAtor.Name { // Custo: C18, Número de execuções: 1
					grafo[ator.Name] = append(grafo[ator.Name], outroAtor.Name) // Custo: C19, Número de execuções: 1
				}
			}
		}
	}
	return grafo
}

func numeroDeBacon(grafo map[string][]string, origem, destino string) int {
	if origem == destino { // Custo: C20, Número de execuções: 1
		return 0
	}

	visitado := make(map[string]bool)      // Custo: C21, Número de execuções: 1
	fila := []string{origem}               // Custo: C22, Número de execuções: 1
	distancia := map[string]int{origem: 0} // Custo: C23, Número de execuções: 1

	for len(fila) > 0 { // Custo: C24, Número de execuções: n
		atual := fila[0] // Custo: C25, Número de execuções: n
		fila = fila[1:]  // Custo: C26, Número de execuções: n

		for _, vizinho := range grafo[atual] { // Custo: C27, Número de execuções: n
			if !visitado[vizinho] { // Custo: C28, Número de execuções: n
				distancia[vizinho] = distancia[atual] + 1 // Custo: C29, Número de execuções: n
				if vizinho == destino {                   // Custo: C30, Número de execuções: n
					return distancia[vizinho] // Custo: C31, Número de execuções: 1
				}
				visitado[vizinho] = true     // Custo: C32, Número de execuções: n
				fila = append(fila, vizinho) // Custo: C33, Número de execuções: n
			}
		}
	}

	return -1 // Retorna -1 se não houver caminho entre origem e destino
}

func Problema5() {
	caminhoCSV := "problema5/tmdb_5000_credits.csv"                      // Custo: C1, Número de execuções: 1
	grafo := lerCSV(caminhoCSV)                                          // Custo: C2, Número de execuções: 1
	fmt.Printf("Grafo de atores construído com %d atores\n", len(grafo)) // Custo: C3, Número de execuções: 1
	origem := "Kevin Bacon"                                              // Custo: C4, Número de execuções: 1
	destino := "Shane Carruth"                                           // Custo: C5, Número de execuções: 1
	numeroBacon := numeroDeBacon(grafo, origem, destino)                 // Custo: C6, Número de execuções: 1

	if numeroBacon != -1 {
		fmt.Printf("O Número de Bacon entre %s e %s é: %d\n", origem, destino, numeroBacon)
	} else {
		fmt.Printf("Não há caminho entre %s e %s\n", origem, destino)
	}
}
