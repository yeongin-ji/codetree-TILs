package main

import (
	"fmt"
)

func main() {
	cnt := 0
	var arr [200]string
	for {
		var s string
		fmt.Scan(&s)
		if s == "0" {
			break
		}
		arr[cnt] = s
		cnt++
	}	
	fmt.Println(cnt)
	for i := range cnt {
		if i%2==0 {
			fmt.Println(arr[i])
		}
	}
}