package main

import "fmt"

func main() {
/*
 Specifiche del Programma
  Questo programma legge una sequenza di numeri interi, e per ogni coppia consecutiva (Precedente, Corrente),
  combina le loro **ultime due cifre** per formare un nuovo numero.
  Il ciclo continua finché il **risultatoCombinato** non è esattamente 100.
  Calcolo: risultatoCombinato = (ultime 2 cifre di Precedente) * 100 + (ultime 2 cifre di Corrente)
*/
	var c, n, o int
	fmt.Scan(&n)
	for c != 100 {
		o = n
		fmt.Scan(&n)
		c = (o%100)*100 + n%100
		fmt.Println(c)
	}
}

