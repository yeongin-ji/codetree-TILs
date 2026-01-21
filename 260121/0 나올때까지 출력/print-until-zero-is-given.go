package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		fmt.Println(n)
	}
}