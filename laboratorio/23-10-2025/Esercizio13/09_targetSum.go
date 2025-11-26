package main
import "fmt"
import "math/rand"

func main(){
    const TARGET = 50

    var somma int = 0
    var conteggio int = 0

    for somma < TARGET {
        n := rand.Intn(10) + 1

        fmt.Print(n, " ")

        somma += n
        conteggio++
    }

    fmt.Println()

    fmt.Printf("%d numeri casuali\n", conteggio)
    fmt.Printf("Somma: %d\n", somma)

    media := float64(somma) / float64(conteggio)
    fmt.Printf("Media: %f\n", media)

    
}