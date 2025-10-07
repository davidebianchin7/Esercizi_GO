package main
import "fmt"

func main() {
  var scelta int;
  var controllo_ciclo int = 1;
  
  for controllo_ciclo == 1 {  
    fmt.Print("\nBenvenuto nel convertitore, fai la tua scelta\n1. Conversione decimale\n2. Conversione intere\n3. Esci\nInserisci il Numero: ");
    fmt.Scan(&scelta);

    switch scelta {
      case 1:
      
      case 2:
        
      case 3:
        fmt.Println("\n per averci scelto!");
        controllo_ciclo = 0;

      default:
        fmt.Println("\nComando non trovato!");
    }
  };
}
