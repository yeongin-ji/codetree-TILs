package main

import "fmt"

func main() {
	var a, b int
	prod := 1
	fmt.Scanf("%d %d", &a, &b)
	for range b {
		prod *= a
	}	
	fmt.Print(prod)
}