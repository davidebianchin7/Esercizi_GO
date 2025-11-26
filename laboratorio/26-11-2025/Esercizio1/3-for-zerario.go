package main

import "fmt"
import "math/rand"

func main(){
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        
        fmt.Print(n, " ")

        if n == 0 {
            break
        }
        c++
    }

    fmt.Println()
    fmt.Println(c)
}