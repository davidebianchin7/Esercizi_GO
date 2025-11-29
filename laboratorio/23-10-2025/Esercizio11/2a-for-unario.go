package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
Specifiche del Programma

	Questo programma genera numeri casuali e li somma in modo cumulativo fino a quando il totale raggiunge o supera l'obiettivo impostato.

	* Obiettivo (TARGET = 20): La somma dei numeri generati deve raggiungere o superare 20.
	* Range dei Numeri (MAX = 10): Ogni numero generato è un intero casuale tra 1 e 10.
	* Funzionamento: Il ciclo continua finché la variabile di somma (t) è minore di TARGET. Ad ogni giro, viene generato un numero casuale (n), stampato, e aggiunto alla somma (t).
	* Output: Stampa la sequenza dei numeri generati, seguita dal totale finale raggiunto (che sarà >= 20) su una nuova riga.

	Suggerimento per la variabile 't':
	La variabile 't' rappresenta la somma cumulativa. Si suggerisce di rinominarla in 'sommaTotale' per una migliore leggibilità.
    
    */
    const TARGET = 20
    const MAX = 10
    //rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Print("\n",t,"\n")
}

