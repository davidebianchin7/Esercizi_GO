package main

import "fmt"

func main() {
	/*
	 Specifiche del programma:
	- Leggere un numero intero da input.
	- Stabilire se il numero è positivo, nullo o negativo.
	- Stampare "positivo" se il numero è maggiore di 0.
	- Stampare "nullo" se il numero è uguale a 0.
	- Stampare "negativo" se il numero è minore di 0.
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else {
		fmt.Println("negativo")
	}
}

