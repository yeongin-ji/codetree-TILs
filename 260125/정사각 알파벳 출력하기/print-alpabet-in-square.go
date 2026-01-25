package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for range n {
		for range n {
			fmt.Printf("%c", 'A'+cnt)
			cnt = (cnt+1)%26
		}
		fmt.Println()
	}
		
}