/*
GUESS.exe:
	- riceve tre numeri interi
	- controlla il corretto inserimento dei numeri e ne stampa il risultato
	- X 
	-stampa la media dei tre valori inseriti
*/

package main

import "fmt"

func main() {
	var num1, num2, num3 int;
	const TERMINI = 3;

	//inserimento da tasteira
	fmt.Println("Inserisci tre numeri interi: ");
	_, err := fmt.Scan(&num1, &num2, &num3);
	fmt.Println("errore in lettura: ", err);
	
	//valori letti
	fmt.Println("valori letti: ", TERMINI, "/3");

	//calcolo media
	media := (num1 + num2 + num3) / TERMINI;
	fmt.Println("media interi: ", media);
}
