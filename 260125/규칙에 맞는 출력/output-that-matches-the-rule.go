package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			if n-i <= j+1 {
				fmt.Printf("%d ", j+1)
			}
		}
		fmt.Println()
	}	
}