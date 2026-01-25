package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for range i {
			fmt.Print("  ")
		}
		for j := range n-i {
			fmt.Printf("%d ", (n-i)-j)
		}
		fmt.Println()
	}	
}