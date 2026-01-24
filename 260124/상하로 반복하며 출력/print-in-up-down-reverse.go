package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			if j%2==0 {
				fmt.Print(i+1)
			} else {
				fmt.Print(n-i)
			}
		}
		fmt.Println()
	}
}