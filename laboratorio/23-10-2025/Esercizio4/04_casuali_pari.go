package main

import "fmt"
import "math/rand"

func main(){
  var casual, n_pari int

  for i := 0; i < 10; i++{
    casual = rand.Intn(11)
    fmt.Print(casual, " ")

    if casual % 2 == 0 {
      n_pari++
    }
  }
  
  fmt.Println()
  fmt.Printf("Numeri pari generati: %d\n", n_pari )
}
