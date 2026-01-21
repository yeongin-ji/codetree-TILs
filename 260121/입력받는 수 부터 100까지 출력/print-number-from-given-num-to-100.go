package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := n; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
}