/*

Avendo un numero N, partendo da 0 elevare al quadrato tutti i numeri e verificare che il risultato sia <= a n, 
quando diventera >= n allora finisce il programma, restituendo la somma dei quadrati fino a n

Es: n = 7  --->  1^2 + 2^2 + ... + 7^2

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
      for i := 0; i * i <= num_input; i++{
        somma += i * i;
      }

      fmt.Printf("La somma dei quadrati minori o uguali a %d è %d\n", num_input, somma);
    }
  }

}
