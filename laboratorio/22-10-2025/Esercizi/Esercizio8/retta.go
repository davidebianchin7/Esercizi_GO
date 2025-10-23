package main

import (
	"fmt"
	"math"
)

const Epsilon = 1e-9

func main() {
	var m, q float64
	var px, py float64

	fmt.Print("Inserisci m e q: ")
	_, err := fmt.Scan(&m, &q)
	if err != nil {
		fmt.Println("Errore nell'inserimento dei dati della retta:", err)
		return
	}

	fmt.Print("Inserisci x e y: ")
	_, err = fmt.Scan(&px, &py)
	if err != nil {
		fmt.Println("Errore nell'inserimento dei dati del punto:", err)
		return
	}

	y_retta := m*px + q

	if math.Abs(py-y_retta) < Epsilon {
		fmt.Println("Il punto appartiene alla retta")
	} else if py > y_retta {
		fmt.Println("Il punto sta sopra la retta")
	} else {
		fmt.Println("Il punto sta sotto la retta")
	}
}