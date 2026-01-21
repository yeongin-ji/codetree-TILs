package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	i := n
	for i >= 1 {
		fmt.Printf("%d ", i)
		i--
	}	
}