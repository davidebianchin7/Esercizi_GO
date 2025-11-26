package main

import "fmt"

func main(){
    var a, b int

    fmt.Print("Inserisci due numeri interi (a, b): ")
    _, err := fmt.Scan(&a, &b)

    if err != nil {
        fmt.Println("Inserisci due numeri Interi!!!")
    }

    fmt.Println(MCD(a, b))
    
}

func MCD(a, b int) int {
    for b != 0 {
        temp := a % b
        a = b
        b = temp
    }

    return a
}