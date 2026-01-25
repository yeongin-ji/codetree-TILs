package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 1
	for i := range n {
		for range i+1 {
			fmt.Printf("%d ", cnt)
			cnt++
		}
		fmt.Println()
	}	
}