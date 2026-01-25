package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 2; i <= 8; i += 2 {
		for j := b; j >= a; j-- {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
			if j!=a {
				fmt.Print("/ ")
			}
		}
		fmt.Println()
	}
}