package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [1100]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	s := arr[:]
	m := n
	for {
		max := s[0]
		idx := 0	
		for i := range m {
			if s[i] > max {
				max = s[i]
				idx = i
			}
		}
		fmt.Printf("%d ", idx+1)
		if idx == 0 {
			break
		}
		s = s[:idx]
		m = idx
	}
}