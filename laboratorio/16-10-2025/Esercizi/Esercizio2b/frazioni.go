package main

import "fmt"

func main() {
  var num1, den1, num2, den2 int;

  fmt.Print("Inserisci il primo Numeratore e il primo Denominatore: ");
  fmt.Scan(&num1, &den1);

  fmt.Print("Inserisci il secondo Numeratore e il secondo Denominatore: ");
  fmt.Scan(&num2, &den2);


  if num1 * den2 == num2 * den1 {
    fmt.Println("equivalenti");
  } esle {
    fmt.Println("non equivalenti");
  }
}
