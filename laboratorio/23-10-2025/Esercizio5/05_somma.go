package main

import "fmt"
import "math/rand"

func main(){
  var casual, somma int

  for i := 0; i < 10; i++ {
    casual = rand.Intn(10) + 1
    fmt.Print(casual, " ")
    somma += casual
  }

  fmt.Printf("\nLa somma è: %d\n", somma)
}
