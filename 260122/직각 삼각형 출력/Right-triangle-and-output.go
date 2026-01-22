package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 1; j <= i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}	
}