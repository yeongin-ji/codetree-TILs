package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	l1, l2 := len(s1), len(s2)
	if l1 < l2 {
		fmt.Print(s2, l2)
	} else if l1 > l2 {
		fmt.Print(s1, l1)
	} else {
		fmt.Print("same")
	}
}