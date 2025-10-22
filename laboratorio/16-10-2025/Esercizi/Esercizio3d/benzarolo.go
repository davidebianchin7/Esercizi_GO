package main

import "fmt"

func main() {
	var litri float64;
	var tipoCarburante int;

	const P_BENZINA = 1.78;
	const P_DIESEL = 1.98;
	const P_ALCOL = 1.23;
	const P_CHEROSENE = 1.1;

	var prezzoTotale float64;
	var nomeCarburante string;

	fmt.Print("Inserire 0, 1, 2, 3 per scegliere il tipo di carburante (benzina, diesel, alcol, cherosene): ");
	fmt.Scan(&tipoCarburante);

	fmt.Print("Inserire la quantità di carburante in L: ");
	fmt.Scan(&litri);

	
	if tipoCarburante == 0 {
		prezzoTotale = litri * P_BENZINA;
		nomeCarburante = "Benzina";
		
	} else if tipoCarburante == 1 {
		prezzoTotale = litri * P_DIESEL;
		nomeCarburante = "Diesel";
		
	} else if tipoCarburante == 2 {
		prezzoTotale = litri * P_ALCOL;
		nomeCarburante = "Alcol";
		
	} else if tipoCarburante == 3 {
		prezzoTotale = litri * P_CHEROSENE;
		nomeCarburante = "Cherosene";
		
	} else {
		fmt.Println("Errore: tipo di carburante non valido. Usare 0, 1, 2 o 3.");
		return;
		
	}

	fmt.Printf("Rifornimento di %.2f litri di %s.\n", litri, nomeCarburante);
	fmt.Printf("Costo totale: %.2f €\n", prezzoTotale);
}
