package main

import "fmt"

func main() {
	var maxPari int
	for {
		var numStr string
		fmt.Scan(&numStr)
		if numStr == "000" {
			break
		}
		pari := 0
		for _, c := range numStr {
			if c >= '0' && c <= '9' && (c-'0')%2 == 0 {
				pari++
			}
		}
		if pari > maxPari {
			maxPari = pari
		}
	}
	fmt.Println(maxPari)
}
