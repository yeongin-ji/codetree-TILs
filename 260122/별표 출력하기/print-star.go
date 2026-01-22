package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := 1; i <= n-1; i++ {
		for j := i; j <= n-1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}