package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for range i+1 {
			fmt.Printf("%d ", i+1)
		}
		fmt.Println()
	}	
}