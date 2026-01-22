package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			for k := i; k <= n; k++ {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}	
}