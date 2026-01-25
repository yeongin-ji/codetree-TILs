package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for i := range n {
		for range i {
			fmt.Print("  ")
		}
		for range n-i {
			fmt.Printf("%c ", 'A'+cnt)
			cnt = (cnt+1)%26
		}
		fmt.Println()
	}	
}