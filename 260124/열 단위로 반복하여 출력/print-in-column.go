package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for range n {
			fmt.Print(i+1)
		}
		fmt.Println()
	}	
}