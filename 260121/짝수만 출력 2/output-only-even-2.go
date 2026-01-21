package main

import "fmt"

func main() {
	var b, a int
	fmt.Scanf("%d %d", &b, &a)
	i := b
	for i >= a {
		fmt.Printf("%d ", i)
		i -= 2
	}	
}