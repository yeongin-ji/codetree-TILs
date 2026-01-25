package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for range n {
		var a, b int
		fmt.Scan(&a, &b)
		prod := 1
		for i := a; i <= b; i++ {
			prod *= i
		}
		fmt.Println(prod)
	}	
}