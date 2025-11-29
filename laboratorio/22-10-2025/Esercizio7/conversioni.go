package main

import (
	"fmt"
	"math"
)

const (
	SecondiInMinuto = 60.0
	SecondiInOra    = 3600.0
	MinutiInOra     = 60.0
	MinutiInGiorno  = 1440.0
	MinutiInAnno    = 525600.0
)

func main() {
	var n int

	fmt.Print("Scegli la conversione: \n",
		"1) Secondi -> Ore\n",
		"2) Secondi -> Minuti\n",
		"3) Minuti -> Ore\n",
		"4) Minuti -> Secondi\n",
		"5) Ore -> Secondi\n",
		"6) Ore -> Minuti\n",
		"7) Minuti -> Giorni e Ore\n",
		"8) Minuti -> Anni e Giorni\n",
		"La tua scelta: ",
	)
	fmt.Scan(&n)

	var valore float64
	var prompt string
	sceltaValida := true

	switch n {
	case 1, 2:
		prompt = "Inserisci i secondi: "
	case 3, 4, 7, 8:
		prompt = "Inserisci i minuti: "
	case 5, 6:
		prompt = "Inserisci le ore: "
	default:
		fmt.Println("Scelta non valida")
		sceltaValida = false
	}

	if sceltaValida {
		fmt.Print(prompt)
		fmt.Scan(&valore)

		switch n {
		case 1:
			ore := valore / SecondiInOra
			fmt.Printf("%.2f secondi, corrispondono a %.3f ore\n", valore, ore)
		case 2:
			minuti := valore / SecondiInMinuto
			fmt.Printf("%.2f secondi, corrispondono a %.3f minuti\n", valore, minuti)
		case 3:
			ore := valore / MinutiInOra
			fmt.Printf("%.2f minuti, corrispondono a %.3f ore\n", valore, ore)
		case 4:
			secondi := valore * SecondiInMinuto
			fmt.Printf("%.2f minuti, corrispondono a %.3f secondi\n", valore, secondi)
		case 5:
			secondi := valore * SecondiInOra
			fmt.Printf("%.2f ore, corrispondono a %.3f secondi\n", valore, secondi)
		case 6:
			minuti := valore * MinutiInOra
			fmt.Printf("%.2f ore, corrispondono a %.3f minuti\n", valore, minuti)
		case 7:
			giorni := math.Floor(valore / MinutiInGiorno)
			minutiRimanenti := math.Mod(valore, MinutiInGiorno)
			ore := minutiRimanenti / MinutiInOra
			fmt.Printf("%.2f minuti, corrispondono a %.0f giorni e %.3f ore\n", valore, giorni, ore)
		case 8:
			anni := math.Floor(valore / MinutiInAnno)
			minutiRimanenti := math.Mod(valore, MinutiInAnno)
			giorni := minutiRimanenti / MinutiInGiorno
			fmt.Printf("%.2f minuti, corrispondono a %.0f anni e %.3f giorni\n", valore, anni, giorni)
		}
	}
}