package main
import "fmt"

func trent_to_cent(trentesimi float64) float64 {
    return (trentesimi * 100) / 30;
}

func main() {
  var scelta int;
  var controllo_ciclo int = 1;
  
  for controllo_ciclo == 1 {  
    fmt.Print("\nBenvenuto nel convertitore, fai la tua scelta\n1. Conversione decimale\n2. Conversione intere\n3. Esci\nInserisci il Numero: ");
    fmt.Scan(&scelta);

    switch scelta {
      case 1:
        var input_dec float64;
        fmt.Print("Inserisci il numero in trentesimi: ");
        fmt.Scan(&input_dec);
        fmt.Println("Il risultato in decimale è: ", trent_to_cent(input_dec));
      
      case 2:
       var input_int float64;
        fmt.Print("Inserisci il numero in trentesimi: ");
        fmt.Scan(&input_int);
        fmt.Println("Il risultato in intero è: ", int(trent_to_cent(input_int) + 0.5));
        
      case 3:
        fmt.Println("\n per averci scelto!");
        controllo_ciclo = 0;

      default:
        fmt.Println("\nComando non trovato!");
    }
  };
  
}
