package main

import "fmt"

func main() {
	var n uint8
	fmt.Scanf("%d", &n)

	for i := range n {
		fmt.Printf("%d ", i+1)
	}
}