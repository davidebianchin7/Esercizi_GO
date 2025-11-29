package main

import (
	"fmt"
)

func main() {
	var capitale, tassoPercentuale, obiettivo float64
	var anni int = 0

	fmt.Print("Capitale iniziale: ")
	fmt.Scanln(&capitale)

	fmt.Print("Tasso di interesse annuo (%): ")
	fmt.Scanln(&tassoPercentuale)

	fmt.Print("Obiettivo finanziario: ")
	fmt.Scanln(&obiettivo)

	fattoreCrescita := 1.0 + (tassoPercentuale / 100.0)

	for capitale < obiettivo {
		capitale *= fattoreCrescita
		anni++
	}

	fmt.Println("\n--- Risultato ---")
	fmt.Printf("Occorrono %d anni per raggiungere l'obiettivo.\n", anni)
	fmt.Printf("Il capitale sarà di %.2f.\n", capitale)
}
