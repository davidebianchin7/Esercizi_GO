package main
import "fmt"

func main() {
    var num1, num2 int
    var div float64
    
    fmt.Print("Inserisci due numeri: ")
    fmt.Scan(&num1, &num2)

    div = float64(num1) / float64(num2);

    if num2 == 0 {
        fmt.Printf("Non si può usare %d perchè non si può dividere per 0\n", num2)
    } else {
        fmt.Printf("%d / %d = %.1f\n",num1, num2, div)
    }
}