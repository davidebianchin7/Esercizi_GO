package main

import "fmt"

func main() {
	maxSum := 0
	for {
		var n int
		fmt.Scan(&n)
		if n == 999 {
			break
		}
		sum := 0
		for tmp := n; tmp > 0; tmp /= 10 {
			sum += tmp % 10
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	fmt.Println(maxSum)
}
