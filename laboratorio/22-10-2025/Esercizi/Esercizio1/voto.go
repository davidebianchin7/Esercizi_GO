package main
import "fmt"

func main(){
    var v int

    fmt.Print("Inserisci un voto da 0 a 100: ")
    fmt.Scan(&v)

    if v > 100 || v < 0{
        fmt.Printf("il voto %d inserito non è valido\n", v)
    } else {
        if v < 60 {
            fmt.Printf("il voto %d è Insufficiente\n", v)
        } else if v < 70 {
            fmt.Printf("il voto %d è Sufficiente\n", v)
        } else if v < 80 {
            fmt.Printf("il voto %d è Buono\n", v)
        } else if v < 90 {
            fmt.Printf("il voto %d è Distinto\n", v)
        } else if v <= 100 {
            fmt.Printf("il voto %d è Ottimo\n", v)    
        }
    }
}