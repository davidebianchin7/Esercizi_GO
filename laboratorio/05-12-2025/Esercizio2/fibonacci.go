package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero intero positivo: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Il numero inserito non è valido.")
		return
	}

	a, b := 1, 1
	for i := 0; i < n; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		a, b = b, a+b
	}
}
