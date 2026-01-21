package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := range 5 {
		fmt.Printf("%d ", (i+1)*n)
	}
}