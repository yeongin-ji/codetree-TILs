package main

import (
	"fmt"
)

func main() {
	var arr [5]rune
	for i := range 5 {
		fmt.Scan(&arr[i])
	}
	for i := range 5 {
		fmt.Printf("%c ", arr[i])
	}
}