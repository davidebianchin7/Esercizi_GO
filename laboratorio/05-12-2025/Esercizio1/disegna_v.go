package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	w := 2*n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < w; j++ {
			if j == i || j == w-1-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
