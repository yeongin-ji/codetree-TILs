package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range n {
		for i := range n {
			fmt.Print(i+1)
		}
		fmt.Println()
	}	
}