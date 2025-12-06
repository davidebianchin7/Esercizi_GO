package main

import "fmt"

func main() {
	var n, max, count int
	fmt.Print("Inserisci una sequenza di numeri interi (termina con 0): ")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if count == 0 || n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}
	}
	fmt.Printf("Il valore massimo è %d e compare %d volte\n", max, count)
}
