package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		if (i+1)%2==1 {
			fmt.Println("*")
			continue
		}
		for range i+1 {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}