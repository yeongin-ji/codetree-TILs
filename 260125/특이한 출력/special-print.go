package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			fmt.Printf("(%d, %d) ", i+1, j+1)
			if (i+j+2)%4==0 {
				fmt.Println()
			}
		}
	}
}