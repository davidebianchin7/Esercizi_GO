package main

import "fmt"

func main() {
    var n_pos int

    fmt.Print("Inserire un numero intero NON negativo: ")
    _, err := fmt.Scan(&n_pos)

    if err != nil {
        fmt.Println("Hai inserito un input non valido!")
    }

    posizione := -1
    contatore := 1

    for n_pos > 0 {
        cifra := n_pos % 10
        if cifra % 2 == 0 {
            posizione = contatore
            break
        }
        n_pos /= 10
        contatore++
    }

    fmt.Printf("La prima cifra pari si trova in posizione: %d", posizione)
        
}
