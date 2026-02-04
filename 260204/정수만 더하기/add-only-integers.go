package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var sum int
	for _, v := range s {
		if v>='0' && v<='9' {
			sum += int(v-'0')
		}
	}
	fmt.Print(sum)
}