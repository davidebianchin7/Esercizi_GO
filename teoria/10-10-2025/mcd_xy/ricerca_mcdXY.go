/*

Ricerca del MCD dati x e y

*/

package main;

import "fmt";

func main (){
  var x_input, y_input, mcd int;
  
  fmt.Print("Inserisci il primo numero intero positivo: ");
  _, err1 := fmt.Scan(&x_input);

  fmt.Print("Inserisci il secondo numero intero positivo: ");
  _, err2 := fmt.Scan(&y_input);

  if err1 != nil || err2 != nil {
    fmt.Println("Inserire un numero INTERO!");
    
  } else {
    if x_input <= 0 || y_input <= 0 {
      fmt.Println("Inserire un numero maggiore di zero!");
      
    } else {
      a := x_input;
      b := y_input;

      for b != 0 {
        temp := b;
        b = a % b;
        a = temp;
      }

      mcd = a;

      fmt.Printf("L'MCD tra %d e %d è uguale a %d", x_input, y_input, mcd);
    }
  }

}
