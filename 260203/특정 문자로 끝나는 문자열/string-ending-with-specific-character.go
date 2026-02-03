package main

import (
	"fmt"
)

func main() {
	var arr [10]string
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	var c string
	fmt.Scan(&c)
	flag := false
	for _, s := range arr {
		if c[0]==s[len(s)-1] {
			flag = true
			fmt.Println(s)
		}
	}
	if !flag {
		fmt.Println("None")
	}
}