/*

Stabilire se un numero n è primo

*/

package main;

import "fmt";

func main (){
  var num_input int;
  var is_primo bool = true;

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
      for i := 2; i < num_input; i++{
        if num_input % i == 0 {
          is_primo = false;
          break;
        }
      }

      if is_primo {
         fmt.Printf("%d è un numero primo", num_input)
      } else {
        fmt.Printf("%d NON è un numero primo", num_input)
      }

    }
  }

}
