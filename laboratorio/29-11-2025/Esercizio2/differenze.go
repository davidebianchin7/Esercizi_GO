package main

import (
	"fmt"
)

func main() {
	var valoreCorrente, valorePrecedente float64
	var primaLettura = true
    var indice int = 0
    fmt.Print("Inserisci una serie di numeri separati da uno spazio: ")

	for {
		_, err := fmt.Scanf("%f", &valoreCorrente)

		if err != nil {
            //fmt.Println(err)
            fmt.Println("Hai inserito un valore non valido.")
			break
		}

		if valoreCorrente == 0 {
			break
		}

		if !primaLettura {
            indice++
			differenza := valoreCorrente - valorePrecedente
			fmt.Printf("la %d differenza è: %.2f\n", indice, differenza)
		}

		valorePrecedente = valoreCorrente
		primaLettura = false
	}
}
