package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	prod := 1
	for i := 1; i <= 10; i++ {
		prod *= i
		if prod >= n {
			fmt.Print(i)
			break
		}
	}	
}