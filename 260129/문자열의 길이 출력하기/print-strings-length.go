package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Print(len(s1)+len(s2))
}
