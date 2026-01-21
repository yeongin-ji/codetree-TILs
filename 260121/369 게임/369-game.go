package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 1; i <= N; i++ {
		if i%3==0 {
			fmt.Print("0 ")
			continue
		}
		s := strconv.Itoa(i)
		flag := false
		for _, v := range s {
			if v=='3' || v=='6' || v=='9' {
				flag = true
				fmt.Print("0 ")
				break
			}
		}
		if !flag {
			fmt.Printf("%d ", i)
		}
	}
}