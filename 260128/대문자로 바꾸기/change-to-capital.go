package main

import "fmt"

func main() {
	var arr [5][3]string
	for i := range 5 {
		for j := range 3 {
			fmt.Scan(&arr[i][j])
		}
	}
	for i := range 5 {
		for j := range 3 {
			r := []rune(arr[i][j])
			fmt.Printf("%c ", r[0]-32)
		}
		fmt.Println()
	}
}