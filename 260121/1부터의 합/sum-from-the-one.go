package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= n {
			fmt.Print(i)
			break
		}
	}
}