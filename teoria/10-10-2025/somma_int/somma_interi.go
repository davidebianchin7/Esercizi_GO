/*

Sommare i primi numeri interi fino al numero inserito compreso

Es: n = 7  --->  1 + 2 + ... + 7 = 28

*/

package main;

import "fmt";

func main (){
  var num_input, somma int;

  fmt.Print("Inserisci un numero intero positivo: ");
  verifica, err := fmt.Scan(&num_input);

  if err != nil {
    fmt.Println("Inserire un numero INTERO!");
    
  } else if verifica != 1 {
    fmt.Println("Inserire un solo numero!");
  
  } else {
    if num_input <= 0 {
      fmt.Println("Inserire un numero maggiore di zero!");
    } else {
      somma = 0;
      for i := 1; i <= num_input; i++{
        somma += i;
      }

      fmt.Printf("La somma di tutti i numeri precedenti di %d è uguale a %d", num_input, somma);
    }
  }

}
