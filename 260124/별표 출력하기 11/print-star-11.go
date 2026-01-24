package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n*2+1 {
		if i%2==0 {
			for range n*2+1 {
				fmt.Print("* ")
			}
		} else {
			for range n+1 {
				fmt.Print("*   ")
			}
		}
		fmt.Println()
	}
}