package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	i := 1
	for i <= n {
		if i%3==0 {
			fmt.Printf("%d ", i)
		}
		i++
	}
}