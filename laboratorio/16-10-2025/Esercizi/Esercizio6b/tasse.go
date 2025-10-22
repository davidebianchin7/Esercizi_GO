package main

import "fmt"

func main() {
	const (
		ALIQUOTA_BASSA      = 0.10
		ALIQUOTA_ALTA       = 0.24
		SCAGLIONE_SINGLE    = 32000
		SCAGLIONE_CONIUGATO = 64000
	)

	var reddito int
	var coniugato bool

	var tasse float64
	var scaglione int
	var status string

	fmt.Print("reddito? ")
	fmt.Scan(&reddito)

	fmt.Print("coniugato? (true/false) ")
	fmt.Scan(&coniugato)

	if coniugato {
		scaglione = SCAGLIONE_CONIUGATO
		status = "coniugato"
	} else {
		scaglione = SCAGLIONE_SINGLE
		status = "non coniugato"
	}

	if reddito <= scaglione {
		tasse = float64(reddito) * ALIQUOTA_BASSA
	} else {
		tasse = float64(reddito) * ALIQUOTA_ALTA
	}

	fmt.Printf("tasse per %s con reddito %d: %.0f\n", status, reddito, tasse)
}

/*
Risposta alla domanda:

D1. Quante diverse aliquote da applicare ci sono?
R1. Ci sono 2 diverse aliquote da applicare: l'aliquota bassa (10%)
    e l'aliquota alta.
    (Nota: le specifiche testuali dicono "25%", ma la tabella e
    l'esempio di esecuzione [40000 * 0.24 = 9600] usano il 24%.
    Questo programma userà 10% e 24% per rispettare l'esempio.)

D2. Quanti if/else occorrono come minimo? Di che tipo? Perché?
R2. Occorrono come minimo DUE blocchi if-else (a due vie) separati
    (non annidati).

    Perché?
    Il problema ha due fasi logiche distinte e sequenziali:

    1. Determinare lo SCAGLIONE: Per prima cosa, dobbiamo decidere
       quale soglia (32000 o 64000) usare. Questo dipende solo
       dall'input 'coniugato'. Un 'if-else' a due vie è perfetto
       per assegnare la soglia corretta a una variabile 'scaglione'.

    2. Determinare l'ALIQUOTA: Una volta nota la soglia ('scaglione'),
       dobbiamo confrontare il 'reddito' con quella soglia.
       Un *secondo* 'if-else' a due vie serve per decidere se
       applicare l'ALIQUOTA_BASSA (se reddito <= scaglione) o
       l'ALIQUOTA_ALTA (se reddito > scaglione).

    Questo approccio è più efficiente (DRY - Don't Repeat Yourself)
    rispetto a un 'if' annidato o a una catena a 4 vie, perché la
    logica di calcolo (Passo 2) è identica, cambia solo il
    parametro 'scaglione' (determinato al Passo 1).
*/
