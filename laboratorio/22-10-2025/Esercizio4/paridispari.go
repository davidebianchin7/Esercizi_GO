package main
import "fmt"

func main(){
    var n int

    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)

    if n % 2 == 0 {
        fmt.Printf("Il numero %d è pari\n", n)
    } else {
        fmt.Printf("Il numero %d è dispari\n", n)
    }
    
    
}