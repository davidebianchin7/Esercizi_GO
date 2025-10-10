/*

Dato un numero intero > 0, calcola da quante cifre è composto

*/

package main;

import "fmt";

func main (){
  var num_input, cifre_num int;

  fmt.Print("Inserisci un numero Intero: ");
  verifica, err := fmt.Scan(&num_input);
  
  //fmt.Println("verifica: " ,verifica);
  
  if err != nil {
    fmt.Println("Inserire un numero INTERO!");
    
  } else if verifica != 1 {
    fmt.Println("Inserire un solo numero!");
  
  } else {
    if num_input <= 0 {
      fmt.Println("Inserire un numero maggiore di zero!");
  
    } else {
      cifre_num = 0;
      temp := num_input;
      for temp > 0 {
        temp /= 10;
        cifre_num++;
      }
  
      fmt.Printf("Il numero %d è composto da %d cifre", num_input, cifre_num);
    }
  }
}
