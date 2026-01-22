package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := n-1; i >= 0; i-- {
		for j := n-i; j > 1; j-- {
			fmt.Print("  ")
		}
		for j := i*2+1; j > 0; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}