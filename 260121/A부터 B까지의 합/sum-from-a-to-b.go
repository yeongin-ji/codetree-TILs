package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Print(sum)
}