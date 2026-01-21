package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	flag := false
	for i := a; i <= b; i++ {
		if 1920%i==0 && 2880%i==0 {
			flag = true
		}
	}	
	if flag {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}