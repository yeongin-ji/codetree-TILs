package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n*2 {
		if (i+1)%2!=0 { // odd
			for range i/2+1 {
				fmt.Print("* ")
			}
		} else { // even (1, 3, 5)
			for range n-i/2 {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}	
}