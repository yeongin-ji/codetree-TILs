package main

import "fmt"

func main() {
	flag := true
	for range 5 {
		var n int
		fmt.Scan(&n)
		if n%3!=0 {
			flag = false
		}
	}	
	if flag {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}