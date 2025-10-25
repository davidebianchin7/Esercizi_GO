package main

import "fmt"
import "math/rand"

func main() {
  var casual int

  for i := 0; i < 10; i++{
    casual = rand.Intn(11)
    fmt.Print(casual, " ")
  }
  fmt.Println()
}
