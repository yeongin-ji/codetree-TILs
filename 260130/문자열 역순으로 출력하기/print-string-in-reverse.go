package main

import "fmt"

func main() {
	var s [4]string
	for i := range 4 {
		fmt.Scan(&s[i])
	}	
	for i := range 4 {
		fmt.Println(s[3-i])
	}
}