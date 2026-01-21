package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	prod := 1
	for i := 1; i <= b; i++ {
		if i%a==0 {
			prod *= i
		}
	}	
	fmt.Print(prod)
}