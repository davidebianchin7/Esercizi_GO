package main

import "fmt"

func main() {
	var num int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	if num > 0 && num%10 == 0 {
		fmt.Println("positivo divisibile per 10")
	} else if num%10 == 0 {
		fmt.Println("divisibile per 10")
	} else if num > 0 {
		fmt.Println("positivo")
	}
}
