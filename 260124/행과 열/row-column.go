package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := range a {
		for j := range b {
			fmt.Printf("%d ", (i+1)*(j+1))
		}
		fmt.Println()
	}	
}