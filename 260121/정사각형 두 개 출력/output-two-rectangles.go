package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range 2 {
		for range n {
			for range n {
				fmt.Print("*")
			}
			fmt.Println()
		}
		fmt.Println()
	}	
}