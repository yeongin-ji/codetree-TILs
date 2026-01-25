package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		for j := range n-i {
			fmt.Printf("%d * %d = %d ", i+1, j+1, (i+1)*(j+1))
			if j != n-i-1 {
				fmt.Print("/ ")
			}
		}
		fmt.Println()
	}	
}