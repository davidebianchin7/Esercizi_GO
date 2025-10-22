package main

import "fmt"

func main() {
	var g1, g2, g3 int

	fmt.Print("componente 1, guasto rilevato? (0 per ok) ")
	fmt.Scan(&g1)

	fmt.Print("componente 2, guasto rilevato? (0 per ok) ")
	fmt.Scan(&g2)

	fmt.Print("componente 3, guasto rilevato? (0 per ok) ")
	fmt.Scan(&g3)

	fmt.Println("\nguasti rilevati a:")

	if g1 != 0 {
		fmt.Println("caricamento acqua")
	}

	if g2 != 0 {
		fmt.Println("scarico acqua")
	}

	if g3 != 0 {
		fmt.Println("sistema di riscaldamento")
	}
}
