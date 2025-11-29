package main

import "fmt"
import "math/rand"

const MAX = 10

func main() {
	numeroDaIndovinare := rand.Intn(MAX) + 1

	tentativiMassimi := MAX / 2
	tentativiFatti := 0
	indovinato := false

	fmt.Printf("Indovina un numero tra 1 e %d. Hai un massimo di %d tentativi.\n", MAX, tentativiMassimi)
	fmt.Println("------------------------------------------------------------------")

	for !indovinato && tentativiFatti < tentativiMassimi {
		var inputUtente int

		fmt.Printf("indovina il numero (%d di %d): ", tentativiFatti+1, tentativiMassimi)
		_, err := fmt.Scanf("%d", &inputUtente)

		if err != nil {
			fmt.Println("Input non valido! Digita un numero intero.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if inputUtente < 1 || inputUtente > MAX {
			fmt.Println("fuori intervallo!")
			continue
		}

		tentativiFatti++

		if inputUtente == numeroDaIndovinare {
			indovinato = true
		} else if inputUtente < numeroDaIndovinare {
			fmt.Println("Troppo basso!")
		} else {
			fmt.Println("Troppo alto!")
		}
	}

	fmt.Println("------------------------------------------------------------------")
	if indovinato {
		fmt.Printf("hai indovinato! N. di tentativi: %d\n", tentativiFatti)
	} else {
		fmt.Printf("hai perso, il numero era %d\n", numeroDaIndovinare)
	}
}
