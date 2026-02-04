package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	ra := []rune(a)
	rb := []rune(b)
	i := 0
	for i <= len(ra)-len(rb) {
		found := true
		for j := range len(rb) {
			if i+j>=len(ra) || ra[i+j]!=rb[j] {
				found = false
				break
			}
		}
		if found {
			ra = append(ra[:i], ra[i+len(rb):]...)
			i = 0
			continue
		}
		i++
	}
	a = string(ra)
	fmt.Println(a)
}