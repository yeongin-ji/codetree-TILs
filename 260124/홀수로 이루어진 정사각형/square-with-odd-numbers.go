package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", 11+i*2+j*2)
		}
		fmt.Println()
	}	
}