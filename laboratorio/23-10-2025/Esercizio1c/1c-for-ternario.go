package main

import "fmt"

func main() {
	/*
	Scrivere qui le specifiche per questo programma:
	
	Il programma legge una sequenza di K numeri interi (dove K è una costante
	definita a 5). Per ogni numero da leggere, stampa il prompt "un numero: ".
	Dopo aver letto tutti i K numeri, il programma stampa la loro somma totale.
	*/
	
	const K = 5
	var n int
	
	// Sostituire il nome della variabile s con un nome più significativo
	somma := 0 // 's' è stato rinominato in 'somma'
	
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma += n // somma = somma + n
	}
	fmt.Println(somma)
}
