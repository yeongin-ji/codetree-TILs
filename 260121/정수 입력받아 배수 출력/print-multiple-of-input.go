package main

import "fmt"

func main() {
	var n uint8
	fmt.Scanf("%d", &n)
	for i := range 5 {
		fmt.Printf("%d ", (i+1)*n)
	}	
}