package main

import "fmt"

func main() {
	/*
		Specifiche del programma:
		Scrivere un programma che legge un numero intero (num) da tastiera.
		Il programma controlla due condizioni in modo indipendente:
		1. Se il numero è divisibile per 3, stampa "Fizz".
		2. Se il numero è divisibile per 5, stampa "Buzz".
		Se il numero è divisibile per entrambi (es. 15), stamperà sia "Fizz"
		che "Buzz" (su righe separate).
	*/
	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)
	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}

/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if),
il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- R1.
No, il programma non si comporterebbe nello stesso modo.
Di conseguenza, no, non darei le stesse specifiche.

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- R2.
Il programma attuale usa due istruzioni `if` separate e indipendenti.
Ciò significa che entrambe le condizioni (num%3 e num%5) vengono
sempre controllate, una dopo l'altra. Se entrambe sono vere (es. con 15),
vengono eseguiti entrambi i blocchi di stampa.

Se usassimo un `else if`, le condizioni diventerebbero mutuamente esclusive.
La seconda condizione (`else if num%5 == 0`) verrebbe controllata solo se
la prima (`if num%3 == 0`) fosse falsa.
Ad esempio, con 15:
 1. `15%3 == 0` è VERO.
 2. Stampa "Fizz".
 3. Il blocco `else if` viene completamente saltato.
Il risultato sarebbe solo "Fizz", e non "Fizz" seguito da "Buzz".

- D3. Cosa stampa il programma per ciascuno dei seguenti input: 8, 9 10, 15?
- R3
   8:   (non stampa nulla)
   9:   Fizz
   10:  Buzz
   15:  Fizz
        Buzz
*/
