package main

import "fmt"

func main() {
	var int n
	fmt.Scanf("%d", &n)
	for i := n; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
}