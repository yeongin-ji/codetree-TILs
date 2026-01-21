package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	for i := b; i >= a; i-- {
		fmt.Printf("%d ", i)
	}	
}