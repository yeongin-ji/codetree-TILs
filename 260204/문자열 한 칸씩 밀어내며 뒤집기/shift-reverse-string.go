package main

import (
	"fmt"
)

func main() {
	var s string
	var q int
	fmt.Scanf("%s %d\n", &s, &q)
	r := []rune(s)
	for range q {
		var n int
		fmt.Scan(&n)
		switch n {
		case 1:
			r = append(r[1:], r[0])
			s = string(r)
			fmt.Println(s)
			r = []rune(s)
		case 2:
			r = append(r[len(r)-1:], r[:len(r)-1]...)
			s = string(r)
			fmt.Println(s)
			r = []rune(s)
		case 3:
			for i := 0; i < len(r)/2; i++ {
				t := r[i]
				r[i] = r[len(r)-1-i]
				r[len(r)-1-i] = t
			}
			s = string(r)
			fmt.Println(s)
			r = []rune(s)
		}
	}
}