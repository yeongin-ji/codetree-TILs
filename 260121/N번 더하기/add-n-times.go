package main

import "fmt"

func main() {
	var a, n int
	fmt.Scanf("%d %d", &a, &n)
	for range n {
		a += n
		fmt.Println(a)
	}	
}