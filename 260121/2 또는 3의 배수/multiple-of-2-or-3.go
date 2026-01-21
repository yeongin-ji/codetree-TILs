package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := range n {
		v := i+1
		if v%2==0 || v%3==0 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
		fmt.Print(" ")
	}	
}