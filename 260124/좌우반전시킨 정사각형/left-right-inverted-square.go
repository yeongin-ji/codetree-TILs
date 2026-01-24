package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", (i+1)+(i+1)*j)
		}
		fmt.Println()
	}	
}