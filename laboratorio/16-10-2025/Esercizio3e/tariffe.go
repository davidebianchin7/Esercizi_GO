package main

import "fmt"

func main() {
	var eta int;
	var studente bool;
	var tariffa float64;
	
	fmt.Print("età: ");
	fmt.Scan(&eta);

	fmt.Print("studente? (t/f): ");
	fmt.Scan(&studente);

	
	if eta < 9 {
		tariffa = 0
	} else if (eta >= 9 && eta < 18) || (studente && eta >= 18) {
		tariffa = 5
	} else if (eta >= 18 && eta < 26) || (eta >= 65) {
		tariffa = 7.5
	} else {
		tariffa = 10
	}

	fmt.Printf("ingresso %g euro\n", tariffa)
}
