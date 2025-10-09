package main
import "fmt"

func main() {
  var decimal_float float64;
  
  fmt.Print("Ciao User! Inserisci il tuo numero decimale: ");
  fmt.Scan(&decimal_float);
  
  var int_input int = int(decimal_float);
  var dec = decimal_float - float64(int_input);
  fmt.Print("I numeri decimali sono: ",dec);
}
