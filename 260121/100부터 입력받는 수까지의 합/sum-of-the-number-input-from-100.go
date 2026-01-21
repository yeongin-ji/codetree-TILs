package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for i := n; i <= 100; i++ {
		sum += i
	}
	fmt.Print(sum)
}