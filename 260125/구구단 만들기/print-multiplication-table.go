package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := range 9 {
		for j := b; j >= a; j-=2 {
			fmt.Printf("%d * %d = %d ", j, i+1, j*(i+1))
			if j!=a {
				fmt.Print("/ ")
			}
		}
		fmt.Println()
	}	
}