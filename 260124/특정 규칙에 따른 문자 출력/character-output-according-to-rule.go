package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	for i := range n {
		// n-1, n-2, ..., 1
		for range n-1-i {
			fmt.Print("  ")
		}
		// 1, 2, ..., n-1, n
		for range i+1 {
			fmt.Print("@ ")
		}
		fmt.Println()
	}
	for i := range n-1 {
		// n-1, n-2, ..., 2, 1
		for range n-1-i {
			fmt.Print("@ ")
		}
		fmt.Println()
	}
}
