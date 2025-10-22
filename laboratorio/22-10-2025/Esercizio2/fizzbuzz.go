package main
import "fmt"

func main() {
    var n int

    fmt.Print("Inserisci un numero a piacere: ")
    fmt.Scan(&n)

    if n < 0{
        fmt.Printf("il numero %d è minore di 0", n)
    } else {
        if n % 3 == 0 && n % 5 == 0 {
            fmt.Printf("il numero %d è Fizz Buzz\n", n)
        } else {
            if n % 3 == 0 {
                fmt.Printf("il numero %d è Fizz\n", n)
            } else if n % 5 == 0 {
                fmt.Printf("il numero %d è Buzz\n", n)    
            }
        }
    }
}