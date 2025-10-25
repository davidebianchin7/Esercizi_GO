package main
import "fmt"
func main() {
    /*
    Scrivere qui le specifiche per questo programma e caricare il file su upload
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

