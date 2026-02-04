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
			var a, b int
			fmt.Scan(&a, &b)
			t := r[a-1]
			r[a-1] = r[b-1]
			r[b-1] = t
			s = string(r)
			fmt.Println(s)
		case 2:
			var x, y string
			fmt.Scanf("%s %s\n", &x, &y)
			for i, v := range r {
				if v == rune(x[0]) {
					r[i] = rune(y[0])
				}
			}
			s = string(r)
			fmt.Println(s)
		}
	}

}