package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 1
	for range n {
		for range n {
			fmt.Printf("%d", cnt)
			cnt = cnt%9+1
		}
		fmt.Println()
	}	
}