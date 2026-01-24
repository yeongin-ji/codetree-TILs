package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			if i==0 || j==0 || i==n-1 || j==n-1 {
				fmt.Print("* ")
			} else {
				if i <= j {
					fmt.Print("  ")
				} else {
					fmt.Print("* ")
				}
			}
		}
		fmt.Println()
	}	
}