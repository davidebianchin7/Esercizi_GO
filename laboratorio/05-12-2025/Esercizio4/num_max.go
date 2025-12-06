package main

import "fmt"

func main() {
	max := 0
	count := 0

	for i := 0; i < 10; i++ {
		var n int
		fmt.Print("Inserisci un intero positivo: ")
		fmt.Scan(&n)

		if n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}
	}

	fmt.Printf("Massimo: %d\nOccorrenze: %d\n", max, count)
}
