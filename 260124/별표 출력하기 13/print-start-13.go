package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n*2 {
		if i%2==0 { // even
			for range n-(i/2) {
				fmt.Print("* ")
			}
		} else { // odd
			for range i/2+1 {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}	
}