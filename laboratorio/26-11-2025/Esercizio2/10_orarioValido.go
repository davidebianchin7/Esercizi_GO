package main

import "fmt"

func main(){
    var h, min int

    for {
        fmt.Print("Ore e Minuti: ")

        _, err := fmt.Scanf("%d %d", &h, &min)

        if err != nil {
            fmt.Println("Errore nell'input. Inserire due numeri interi.")

            continue
        }

        hValida := (h >= 0 && h <= 23)
        minValidi := (min >= 0 && min <= 59)

        if hValida == true && minValidi == true {
            break
        }
    }

    fmt.Printf("%d:%d\n", h, min)

    
}