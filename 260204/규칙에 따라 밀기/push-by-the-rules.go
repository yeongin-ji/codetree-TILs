package main

import (
	"fmt"
)

func main() {
	var s, cmd string
	fmt.Scan(&s, &cmd)
	r := []rune(s)
	for _, c := range cmd {
		switch c {
		case 'L':
			r = append(r[1:], r[0])
		case 'R':
			r = append(r[len(r)-1:], r[:len(r)-1]...)
		}
	}
	s = string(r)
	fmt.Println(s)
}