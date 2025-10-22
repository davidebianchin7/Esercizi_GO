package main;

import "fmt"

func main() {
  const NUM_SEGRETO = 7;

  var num int;

  fmt.Print("Indovina il numero (tra 1 e 10: ");
  fmt.Scan(&num);

  if num < 1 || num > 10 {
    fmt.Println("Valore non valido");
  } else if nu, == NUM_SEGRETO{
    fmt.Println("Hai indovinato");
  } else {
    fmt.Println("Non hai indovinato");
  }

}
