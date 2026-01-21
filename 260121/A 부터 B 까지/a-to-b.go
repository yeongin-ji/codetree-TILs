package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	i := a
	for i < b {
		fmt.Printf("%d ", i)
		if i%2==1 {
			i *= 2
		} else {
			i += 3
		}
	}	
}