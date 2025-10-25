package main
import "fmt"

func main(){
  var k int

  fmt.Print("Inserisci 5 numeri a piacere: ")
  for i := 0; i < 5; i++ {
    fmt.Scan(&k)
    doppio := k * 2
    fmt.Printf("il dopo di %d è %d\n", k, doppio)
  }
}
