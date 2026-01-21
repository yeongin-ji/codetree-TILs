package main

import "fmt"

func main() {
	var a, b, s, l int
	fmt.Scanf("%d %d", &a, &b)
	if a <= b {
		s, l = a, b
	} else {
		s, l = b, a
	}
	sum := 0
	for i := s; i <= l; i++ {
		if i%5==0 {
			sum += i
		}
	}
	fmt.Print(sum)
}