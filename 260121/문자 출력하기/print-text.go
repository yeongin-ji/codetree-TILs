package main

import "fmt"

func main() {
	var c rune
	fmt.Scanf("%c", &c)
	for range 8 {
		fmt.Printf("%c", c)
	}	
}