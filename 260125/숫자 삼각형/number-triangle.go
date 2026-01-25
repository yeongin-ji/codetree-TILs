package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range i+1 {
			fmt.Printf("%d ", j+1)
		}
		fmt.Println()
	}	
}