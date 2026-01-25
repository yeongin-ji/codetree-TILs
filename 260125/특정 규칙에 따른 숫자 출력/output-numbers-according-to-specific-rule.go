package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 1
	for i := range n {
		for range i {
			fmt.Print("  ")
		}
		for range n-i {
			fmt.Printf("%d ", cnt)
			cnt = cnt%9+1
		}
		fmt.Println()
	}	
}