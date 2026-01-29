package main

import "fmt"

func main() {
	arr := [5]string{
		"apple", "banana", "grape", "blueberry", "orange",
	}	
	var c string
	fmt.Scan(&c)
	cnt := 0
	for _, s := range arr {
		if c[0]==s[2] || c[0]==s[3] {
			fmt.Println(s)
			cnt++
		}
	}
	fmt.Print(cnt)
}