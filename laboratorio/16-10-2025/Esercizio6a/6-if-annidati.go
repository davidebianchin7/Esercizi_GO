package main

import "fmt"

func main() {

	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	
	var figura, operazione int
	
	fmt.Println("di che figura si tratta?")
	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	
	fmt.Println("cosa vuoi calcolare?")
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)
	
	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { 
			fmt.Println("formula = base * altezza")
		}
	} else { 
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { 
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}

/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- R1.
Ci sono 4 diversi messaggi (formule) che è possibile stampare:
1. "formula = (base + altezza) * 2"
2. "formula = base * altezza"
3. "formula = lato1 + lato2 + lato3 "
4. "formula = (base * altezza)/2"

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- R2.
Sì, sarebbe possibile eliminare la struttura annidata (riducendo il numero
totale di 'if'/'else') utilizzando un'unica catena 'if-else if' e
combinando le condizioni con l'operatore logico AND (&&).

Perché?
Attualmente abbiamo una struttura if-else esterna (per 'figura') e,
all'interno di ciascun ramo, un'altra if-else (per 'operazione').
Possiamo "appiattire" questa logica controllando le 4 combinazioni
totali (2 figure * 2 operazioni = 4 casi) in modo lineare.

Esempio della struttura alternativa:
if figura == RETTANGOLO && operazione == PERIMETRO {
    fmt.Println("formula = (base + altezza) * 2")
} else if figura == RETTANGOLO && operazione == AREA {
    fmt.Println("formula = base * altezza")
} else if figura == TRIANGOLO && operazione == PERIMETRO {
    fmt.Println("formula = lato1 + lato2 + lato3 ")
} else { // figura == TRIANGOLO && operazione == AREA
    fmt.Println("formula = (base * altezza)/2")
}

Questo approccio usa 1 'if', 2 'else if' e 1 'else' (totale 4) invece di
1 'if', 1 'else', 2 'if' interni, 2 'else' interni (totale 6).
*/
