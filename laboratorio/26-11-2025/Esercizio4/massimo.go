package main

import "fmt"
import "math/rand"
import "strconv"

func main(){
    var massimo int = 0
    var str string = ""
    
    for i := 0; i <= 10; i++{

        var temp = rand.Intn(31 - 10) + 10

        if massimo < temp {
            massimo = temp
        }
        
        str += strconv.Itoa(temp)
        
        if i != 10 {
            str += " "
        }

        
    }

    fmt.Printf("%s\n", str)
    fmt.Printf("Massimo: %d\n", massimo)
    
}