package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	i := 0
	for i < n {
		fmt.Println("*")
		i++
	}	
}