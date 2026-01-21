package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	flag := true
	for i := a; i <= b; i++ {
		if i%c==0 {
			flag = false
		}
	}	
	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}