package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a > 0 {
		for range b {
			fmt.Print(a)
		}
	} else {
		fmt.Print(0)
	}
}