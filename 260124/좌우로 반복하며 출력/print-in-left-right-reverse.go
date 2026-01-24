package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range n {
		if i%2==0 {
			for i := range n {
				fmt.Print(i+1)
			}
		} else {
			for i := range n {
				fmt.Print(n-i)
			}
		}
		fmt.Println()
	}	
}