package main

import "fmt"

func main() {
	var c rune
	var n int
	fmt.Scanf("%c %d", &c, &n)
	switch c {
	case 'A':
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i)
		}
	case 'D':
		for i := n; i >= 1; i-- {
			fmt.Printf("%d ", i)
		}
	}	
}