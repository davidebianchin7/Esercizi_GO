package main

import "fmt"

func main() {
	var valore1, valore2 int;

	fmt.Print("Inserisci i tuoi numeri (separati da uno spazio): ");
	_, err := fmt.Scan(&valore1, &valore2);
	fmt.Println(err);

	fmt.Println("Valore 1 e 2 in ordine di inserimento: ", valore1, valore2);

	valore1, valore2 = valore2, valore1;
	
	fmt.Println("Valore 1 e 2 scambiati: ", valore1, valore2);
}
