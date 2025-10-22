package main

import "fmt"

func main() {
	var angolo1, angolo2 float64;
	const SOMMA_ANGOLI = 180;

	fmt.Print("Inserisci i tuoi numeri (separati da uno spazio): ");
	fmt.Scan(&angolo1, &angolo2);
	
	angolo3 := SOMMA_ANGOLI - (angolo1 + angolo2);

	fmt.Println("L'ampiezza del terzo angolo è: ", angolo3);
}
