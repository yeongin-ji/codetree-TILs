package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	for i := a; i <= b; i++ {
		if i%2==1 {
			fmt.Printf("%d ", i)
		}
	}
}