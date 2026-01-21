package main

import "fmt"

func main() {
	var a, b int
	var s, l int
	fmt.Scanf("%d %d", &a, &b)
	if a <= b {
		s, l = a, b
	} else {
		s, l = b, a
	}
	for i := l; i >= s; i-- {
		fmt.Printf("%d ", i)
	}
}