package main

import "fmt"

func main() {
	var divisore_input, dividendo_input int;
	
	fmt.Print("Inserisci i tuoi numeri Interi (separati da uno spazio): ");
	fmt.Scan(&divisore_input, &dividendo_input);
	
	//calcolo quoziente
	quoziente := dividendo_input / divisore_input;

	//calcolo resto
	resto := dividendo_input % divisore_input;

	//stampa
	fmt.Println("Quoziente: ", quoziente);
	fmt.Println("Resto: ", resto);
}
