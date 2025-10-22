package main

import "fmt"

func main() {
	/* Specifiche del programma:
	Scrivere un programma che chiede all'utente un voto numerico intero (voto).
	1. Il programma prima controlla se il voto è valido, cioè compreso nell'intervallo [0, 100].
	2. Se il voto non è valido (voto < 0 o voto > 100), stampa "voto non valido" e termina.
	3. Se il voto è valido, il programma stampa la corrispondente lettera americana secondo la scala:
	   - [90, 100]: A
	   - [80, 89]:  B
	   - [70, 79]:  C
	   - [60, 69]:  D
	   - [0, 59]:   F
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- R1. È vera per tutti i valori di 'voto' che sono strettamente minori di 0 (es. -1, -5) OPPURE strettamente maggiori di 100 (es. 101, 200). In breve, per tutti i valori al di fuori dell'intervallo [0, 100].

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- R2. Per tutti i valori di 'voto' compresi nell'intervallo [80, 89] (cioè da 80 incluso a 89 incluso).

- D3. Giustificate la vostra risposta alla domanda D2
- R3. L'istruzione si trova in un blocco `else if voto >= 80`. A causa della struttura a catena `if-else if`, questo blocco viene raggiunto solo se la condizione precedente (`if voto >= 90`) è FALSA.
Quindi, per stampare "B", 'voto' deve essere:
1. NON >= 90 (quindi < 90)
2. E deve essere >= 80
La combinazione di "minore di 90" e "maggiore o uguale a 80" dà l'intervallo [80, 89].

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così? } else if voto >= 80 && voto < 90 {
- R4. Il programma è impostato correttamente così com'è. La versione con `&& voto < 90` è ridondante (superflua).

- D5. Perché?
- R5. Perché, come spiegato in R3, la struttura `else if` garantisce che il controllo `voto >= 80` venga eseguito SOLO SE il controllo `voto >= 90` è già risultato falso. Se il programma arriva a quel punto, sappiamo GIÀ che `voto < 90`. Aggiungere `&& voto < 90` è inutile perché la condizione è già implicitamente soddisfatta.
*/
