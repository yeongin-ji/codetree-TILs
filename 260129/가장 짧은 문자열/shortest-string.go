package main

import "fmt"

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)
	l1, l2, l3 := len(s1), len(s2), len(s3)
	max, min := l1, l1
	if l2 > max {
		max = l2
	}
	if l3 > max {
		max = l3
	}
	if l2 < min {
		min = l2
	}
	if l3 < min {
		min = l3
	}
	fmt.Print(max-min)
}