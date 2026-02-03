package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	cnt := 0
	for i := 0; i <= len(a)-len(b); i++ {
		if a[i]==b[0] && a[i+1]==b[1] {
			cnt++
		}
	}
	fmt.Print(cnt)
}