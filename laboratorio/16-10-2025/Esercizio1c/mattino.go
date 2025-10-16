package main

import "fmt"

func main() {
	var orario int

	fmt.Print("scrivi un orario (int): ")
	fmt.Scan(&orario)

	if orario >= 6 && orario < 13 {
		fmt.Println("mattino")
	}
}
