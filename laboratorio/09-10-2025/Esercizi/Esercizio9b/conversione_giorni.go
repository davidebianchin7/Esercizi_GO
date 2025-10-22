package main

import (
	"fmt"
)

func main() {
	const G_ANNO = 365;
	const G_SETTIMANA = 7;
	var giorni, giorniFinali, settimane, anni int;
	
	fmt.Print("numero di giorni da convertire? ");
	fmt.Scan(&giorni);

	anni = giorni / G_ANNO;
	temp1 := giorni - (anni * G_ANNO);
	settimane = temp1 / G_SETTIMANA;
	temp2 := temp1 - (settimane * G_SETTIMANA);
	giorniFinali = temp2; 

	fmt.Println(giorni, "giorni =", anni, "anni +", settimane, "settimane +", giorniFinali, "giorni");
}
