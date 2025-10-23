package main

import "fmt"

func main() {
	/* Qui sotto trovate quattro commenti (a., b., c., d.) e poi quattro cicli for.
	   Per ciascun ciclo for scegliete il commento corrispondente a quello che fa il ciclo for e spostatelo (taglia e incolla) al posto della riga "//...". */

	// a. stampa i numeri tra 0 e n (escluso) espressi in base 3
	// b. stampa le potenze di 2 tra 1 e n (escluso)
	// c. stampa gli interi quadrati perfetti tra 0 e n (escluso)
	// d. calcola la somma dei primi n valori non negativi letti

	// ciclo 1
	// d. calcola la somma dei primi n valori non negativi letti
	var value, sum float64
	n := 5
	for i := 0; i < n; i++ {
		fmt.Scan(&value)
		if value < 0 {
			i-- // Se il valore è negativo, ripete l'iterazione
		} else {
			sum += value
		}
	}
	fmt.Println(sum)

	// ciclo 2
	// b. stampa le potenze di 2 tra 1 e n (escluso)
	n = 10
	for i := 1; i < n; i *= 2 {
		fmt.Print(i, "\t")
	}
	fmt.Println()

	// ciclo 3
	// a. stampa i numeri tra 0 e n (escluso) espressi in base 3
	const BASE = 3
	n = BASE * BASE // n = 9
	for i := 0; i < n; i++ {
		// i/BASE è la cifra più significativa, i%BASE la meno significativa
		fmt.Print(i, "\t", i/BASE, i%BASE, "\n")
	}
	fmt.Println()

	// ciclo 4
	// c. stampa gli interi quadrati perfetti tra 0 e n (escluso)
	n = 20
	for i := 0; i*i < n; i++ {
		fmt.Print(i*i, "\t")
	}
	fmt.Println()
}
