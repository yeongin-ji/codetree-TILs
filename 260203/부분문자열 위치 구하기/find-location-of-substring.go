package main

import (
	"fmt"
)

func main() {
	var n, m string
	fmt.Scan(&n, &m)
	flag := true
	for i := 0; i < len(n)-len(m); i++ {
		found := true
		for j := range len(m) {
			if n[i+j] != m[j] {
				found = false
			}
		}
		if found {
			fmt.Print(i)
			flag = false
			break
		}
	}
	if flag {
		fmt.Print(-1)
	}
}