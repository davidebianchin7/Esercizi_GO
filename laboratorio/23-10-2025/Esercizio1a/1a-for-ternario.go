package main
import "fmt"
func main() {
    /*
    - Richiedere un numero intero
    - Stampare "*" su una riga lunga quanto il valore inserito  
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

