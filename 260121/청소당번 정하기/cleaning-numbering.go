package main

import "fmt"

func main() {
	var c2, c3, c12, n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		if i%12==0 {
			c12++
		} else if i%3==0 {
			c3++
		} else if i%2==0 {
			c2++
		}
	}
	fmt.Print(c2, c3, c12)
}