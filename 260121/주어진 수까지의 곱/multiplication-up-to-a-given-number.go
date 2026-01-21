package main

import "fmt"

func main() {
	var a, b int
	prod := 1
	fmt.Scanf("%d %d", &a, &b)
	for i := a; i <= b; i++ {
		prod *= i
	}	
	fmt.Print(prod)
}