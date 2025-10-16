package main

import "fmt"

func main() {
	/*
		Specifiche del programma:

		1.  Dichiarare una variabile di tipo intero chiamata 'voto'.
		2.  Chiedere all'utente di inserire un valore.
		3.  Leggere l'intero inserito dall'utente e memorizzarlo.
		4.  Controllare se il valore di 'voto' è non valido, ovvero se è minore di 0 oppure maggiore di 30.
		5.  Se la condizione è vera, stampare a schermo il messaggio "voto non valido".
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}

