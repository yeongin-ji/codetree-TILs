package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		if i%2==0 {
			for j := range n {
				fmt.Printf("%d ", i*n+(j+1))
			}
		} else {
			for j := range n {
				fmt.Printf("%d ", i*n-j)
			}
		}
		fmt.Println()
	}	
}