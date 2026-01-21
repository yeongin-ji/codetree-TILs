package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := N; i <= 100; i++ {
		if i >= 90 {
			fmt.Print("A ")
		} else if i >= 80 {
			fmt.Print("B ")
		} else if i >= 70 {
			fmt.Print("C ")
		} else if i >= 60 {
			fmt.Print("D ")
		} else {
			fmt.Print("F ")
		}
	}	
}