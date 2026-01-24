package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range n {
		for i := range n {
			fmt.Printf("%d ", n-i)
		}
		fmt.Println()
	}	
}