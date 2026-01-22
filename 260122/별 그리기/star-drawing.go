package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for range (n-1)-i {
			fmt.Print(" ")
		}
		for range (i+1)*2-1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := range n-1 {
		for range i+1 {
			fmt.Print(" ")
		}
		for range (n-1-i)*2-1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}