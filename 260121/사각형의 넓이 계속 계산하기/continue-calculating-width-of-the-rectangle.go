package main

import "fmt"

func main() {
	for {
		var w, h int
		var c rune
		fmt.Scanf("%d %d %c\n", &w, &h, &c)
		area := w*h
		fmt.Println(area)
		if c == 'C' {
			break
		}
	}	
}