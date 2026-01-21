package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for range n {
		for range m {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}