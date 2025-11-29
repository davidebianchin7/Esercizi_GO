package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero: ")
	_, err := fmt.Scan(&n)
	
	if err != nil || n <= 0 {
		return 
	}

	var k int = 1
	
	var maxSquare int = 1
	
	for {
		var nextK int = k + 1
		var nextSquare int = nextK * nextK

		if nextSquare <= n {
			maxSquare = nextSquare
			k = nextK
		} else {
			break
		}
	}
	
	fmt.Println(maxSquare)
}
