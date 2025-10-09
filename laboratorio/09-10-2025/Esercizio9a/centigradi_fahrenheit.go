package main

import (
	"fmt"
)

func main() {
	const FHT_CENT = 9./5.;
	const GRD_CENT = 32;
	var gradi float64;
	
	fmt.Print("temperatura in centigradi? ");
	fmt.Scan(&gradi);

	fahrenheit := (gradi * FHT_CENT) + GRD_CENT ;

	fmt.Println(gradi, "°C = ", fahrenheit, "°F");
}
