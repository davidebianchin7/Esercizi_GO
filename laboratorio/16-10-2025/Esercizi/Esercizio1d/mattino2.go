package main

import "fmt"

func main() {
	var ore, minuti int
	const MIN = 60;

	fmt.Print("Inserisci le ore e i minuti: ")
	fmt.Scan(&ore, &minuti)

	minutiTotali := ore * MIN + minuti

	if minutiTotali >= 330 && minutiTotali < 750 {
		fmt.Println("mattino")
	}
}
