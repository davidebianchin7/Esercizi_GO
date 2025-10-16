package main

import "fmt"

func main() {
	/*
	 Specifiche del programma:
	 - Leggere un numero intero da input.
	 - Stabilire se il numero inserito è pari o dispari.
	 - Stampare un messaggio che indica il risultato.
	*/

	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	// Se il resto della divisione di n per 2 è 0, il numero è pari.
	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else { // Altrimenti, è dispari.
		fmt.Println(n, "è dispari")
	}
}
