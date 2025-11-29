package main

import "fmt"

func main(){
  var n int

  fmt.Print("Inserisci un numero per creare la tabellina: ")
  fmt.Scan(&n)

  for i := 0; i <= 10; i++{
    calc := n * i

    fmt.Printf("%d x %d = %d\n", i, n, calc)
  }
}
