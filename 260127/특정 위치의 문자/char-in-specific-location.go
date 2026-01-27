package main

import "fmt"

func main() {
	str := [6]string{"L", "E", "B", "R", "O", "S"}
	var c string
	fmt.Scan(&c)
	idx := -1
	for i, v := range str {
		if v == c {
			idx = i
		}
	}
	if idx != -1 {
		fmt.Print(idx)
	} else {
		fmt.Print("None")
	}
}