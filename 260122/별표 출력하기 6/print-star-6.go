package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range i {
			fmt.Print("  ")
		}
		for j := range (n-i)*2-1 {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := range n-1 {
		for j := range n-i-2 {
			fmt.Print("  ")
		}
		for j := range (i+1)*2+1 {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}