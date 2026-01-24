package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 2
	for range n {
		for range n {
			fmt.Printf("%d ", cnt)
			cnt = cnt%8+2
		}
		fmt.Println()
	}	
}