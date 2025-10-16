/* 
Cosa fa il programma:
	- Generare un numero randomico tra 0 e 10
	- Stampare il valore generato
	- Stampare "true" se il valore è 1 o 10, altrimeni "false"

*/

package main
import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10) + 1
	fmt.Println(n)
	fmt.Println(n == 1 || n == 10)
}
