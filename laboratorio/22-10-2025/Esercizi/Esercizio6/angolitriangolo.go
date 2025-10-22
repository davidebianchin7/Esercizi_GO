package main
import "fmt"

func main(){
    var ang1, ang2, ris_ang int
    const AMPIEZZA = 180

    fmt.Print("Inserire le ampiezze dei due angoli: ")
    fmt.Scan(&ang1, &ang2)


    if ang1 + ang2 > 180 {
        fmt.Println("I due angoli non appartengono ad un triangolo")
    } else {
        ris_ang = AMPIEZZA - (ang1 + ang2)

        fmt.Printf("L'ampiezza del terzo angolo è %d\n", ris_ang)
    }
}