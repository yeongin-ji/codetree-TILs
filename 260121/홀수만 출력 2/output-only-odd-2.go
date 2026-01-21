package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &b, &a)

	for i := b; i >= a; i -= 2 {
		fmt.Printf("%d ", i)
	}
}