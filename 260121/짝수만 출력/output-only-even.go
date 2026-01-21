package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	i := a
	for i <= b {
		fmt.Printf("%d ", i)
		i += 2
	}
}